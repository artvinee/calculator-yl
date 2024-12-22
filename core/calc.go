package core

import (
	"fmt"
	"strconv"
	"unicode"
)

type TokenType int

const (
	TokenNumber TokenType = iota
	TokenOperator
	TokenLeftParen
	TokenRightParen
)

type Token struct {
	Type    TokenType
	Value   string
	IsUnary bool
}

func CalculateExpression(expression string) (float64, error) {
	tokens, err := tokenize(expression)
	if err != nil {
		return 0, err
	}

	rpnTokens, err := toRPN(tokens)
	if err != nil {
		return 0, err
	}

	result, err := evaluateRPN(rpnTokens)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func tokenize(expression string) ([]Token, error) {
	var tokens []Token
	var i int
	var prevToken Token

	for i < len(expression) {
		ch := expression[i]

		if unicode.IsSpace(rune(ch)) {
			i++
			continue
		}

		if ch == '(' {
			tokens = append(tokens, Token{Type: TokenLeftParen, Value: string(ch)})
			prevToken = tokens[len(tokens)-1]
			i++
			continue
		}

		if ch == ')' {
			tokens = append(tokens, Token{Type: TokenRightParen, Value: string(ch)})
			prevToken = tokens[len(tokens)-1]
			i++
			continue
		}

		if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			isUnary := false
			if ch == '+' || ch == '-' {
				if len(tokens) == 0 || prevToken.Type == TokenOperator || prevToken.Type == TokenLeftParen {
					isUnary = true
				}
			}

			if isUnary && i+1 < len(expression) && expression[i+1] == ' ' {
				return nil, fmt.Errorf("unary operator '%c' must be directly before the number or '(', without spaces", ch)
			}

			if !isUnary && prevToken.Type == TokenOperator && !prevToken.IsUnary {
				return nil, fmt.Errorf("two operators '%s' and '%s' cannot be next to each other at position %d", prevToken.Value, string(ch), i)
			}

			tokens = append(tokens, Token{Type: TokenOperator, Value: string(ch), IsUnary: isUnary})
			prevToken = tokens[len(tokens)-1]
			i++
			continue
		}

		if unicode.IsDigit(rune(ch)) || ch == '.' {
			start := i
			dotCount := 0
			for i < len(expression) && (unicode.IsDigit(rune(expression[i])) || expression[i] == '.') {
				if expression[i] == '.' {
					dotCount++
					if dotCount > 1 {
						return nil, fmt.Errorf("invalid number format with multiple dots at position %d", i)
					}
				}
				i++
			}
			numberStr := expression[start:i]
			tokens = append(tokens, Token{Type: TokenNumber, Value: numberStr})
			prevToken = tokens[len(tokens)-1]
			continue
		}

		return nil, fmt.Errorf("unknown character '%c' at position %d", ch, i)
	}
	return tokens, nil
}

func toRPN(tokens []Token) ([]Token, error) {
	var outputQueue []Token
	var operatorStack []Token

	for _, token := range tokens {
		switch token.Type {
		case TokenNumber:
			outputQueue = append(outputQueue, token)
		case TokenOperator:
			if token.IsUnary {
				operatorStack = append(operatorStack, token)
				continue
			}
			for len(operatorStack) > 0 {
				top := operatorStack[len(operatorStack)-1]
				if top.Type == TokenOperator && operatorPrecedence(top.Value) >= operatorPrecedence(token.Value) {
					outputQueue = append(outputQueue, top)
					operatorStack = operatorStack[:len(operatorStack)-1]
				} else {
					break
				}
			}
			operatorStack = append(operatorStack, token)
		case TokenLeftParen:
			operatorStack = append(operatorStack, token)
		case TokenRightParen:
			foundLeftParen := false
			for len(operatorStack) > 0 {
				top := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]
				if top.Type == TokenLeftParen {
					foundLeftParen = true
					break
				} else {
					outputQueue = append(outputQueue, top)
				}
			}
			if !foundLeftParen {
				return nil, fmt.Errorf("mismatched parentheses")
			}
		}
	}

	for len(operatorStack) > 0 {
		top := operatorStack[len(operatorStack)-1]
		if top.Type == TokenLeftParen || top.Type == TokenRightParen {
			return nil, fmt.Errorf("mismatched parentheses")
		}
		outputQueue = append(outputQueue, top)
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	return outputQueue, nil
}

func operatorPrecedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

func evaluateRPN(tokens []Token) (float64, error) {
	var stack []float64

	for _, token := range tokens {
		switch token.Type {
		case TokenNumber:
			num, err := strconv.ParseFloat(token.Value, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number '%s'", token.Value)
			}
			stack = append(stack, num)
		case TokenOperator:
			if token.IsUnary {
				if len(stack) < 1 {
					return 0, fmt.Errorf("missing value for unary operator '%s'", token.Value)
				}
				num := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				switch token.Value {
				case "+":
					stack = append(stack, +num)
				case "-":
					stack = append(stack, -num)
				}
			} else {
				if len(stack) < 2 {
					return 0, fmt.Errorf("missing values for operator '%s'", token.Value)
				}
				right := stack[len(stack)-1]
				left := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				var result float64
				switch token.Value {
				case "+":
					result = left + right
				case "-":
					result = left - right
				case "*":
					result = left * right
				case "/":
					if right == 0 {
						return 0, ErrDivisionByZero
					}
					result = left / right
				}
				stack = append(stack, result)
			}
		}
	}

	if len(stack) != 1 {
		return 0, ErrInvalidExpression
	}

	return stack[0], nil
}
