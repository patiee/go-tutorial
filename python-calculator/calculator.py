import math
import time

class Calculator:
    def __init__(self):
        pass

    def get_expression(self, expression_string):
        ex = []
        value = ""

        for ch in expression_string:
            if ch == ' ':
                continue
            if ch in '()^+-*/':
                if value != "":
                    try:
                        number = float(value)
                    except ValueError:
                        raise ValueError(f"Could not parse {value} into float")
                    ex.append(number)
                    value = ""
                ex.append(ch)
            else:
                value += ch

        if value != "":
            try:
                number = float(value)
            except ValueError:
                raise ValueError(f"Could not parse {value} into float")
            ex.append(number)

        return ex

    def get_precedence(self, operation):
        if operation == '^':
            return 3
        elif operation in '*/':
            return 2
        elif operation in '+-':
            return 1
        else:
            return 0

    def get_associativity(self, operation):
        if operation == '^':
            return 'R'
        return 'L'

    def infix_to_postfix(self, ex):
        result = []
        stack = []

        for value in ex:
            if isinstance(value, float):
                result.append(value)
            elif value == '(':
                stack.append(value)
            elif value == ')':
                while stack and stack[-1] != '(':
                    result.append(stack.pop())
                if not stack:
                    raise ValueError("Invalid expression, mismatched brackets")
                stack.pop()
            else:
                while stack and stack[-1] != '(' and (
                    self.get_precedence(stack[-1]) > self.get_precedence(value) or 
                    (self.get_precedence(stack[-1]) == self.get_precedence(value) and self.get_associativity(value) == 'L')
                ):
                    result.append(stack.pop())
                stack.append(value)

        while stack:
            if stack[-1] == '(':
                raise ValueError("Invalid expression, mismatched brackets")
            result.append(stack.pop())

        return result

    def calculate_arithmetic_expression(self, expression_string):
        start = time.time_ns()
        ex = self.get_expression(expression_string)
        ex = self.infix_to_postfix(ex)

        stack = []
        for value in ex:
            if value == '+':
                if len(stack) < 2:
                    raise ValueError("Invalid expression: not enough operands for +")
                operand2 = stack.pop()
                operand1 = stack.pop()
                stack.append(operand1 + operand2)
            elif value == '-':
                if len(stack) < 2:
                    raise ValueError("Invalid expression: not enough operands for -")
                operand2 = stack.pop()
                operand1 = stack.pop()
                stack.append(operand1 - operand2)
            elif value == '*':
                if len(stack) < 2:
                    raise ValueError("Invalid expression: not enough operands for *")
                operand2 = stack.pop()
                operand1 = stack.pop()
                stack.append(operand1 * operand2)
            elif value == '/':
                if len(stack) < 2:
                    raise ValueError("Invalid expression: not enough operands for /")
                operand2 = stack.pop()
                operand1 = stack.pop()
                stack.append(operand1 / operand2)
            elif value == '^':
                if len(stack) < 2:
                    raise ValueError("Invalid expression: not enough operands for ^")
                operand2 = stack.pop()
                operand1 = stack.pop()
                stack.append(math.pow(operand1, operand2))
            else:
                stack.append(value)

        if len(stack) != 1:
            raise ValueError("Invalid expression: leftover operands")

        print("Finished, took", time.time_ns() - start)
        return stack[0]

def main():
    calc = Calculator()

    while True:
        try:
            expression = input("arithmetic expression: ")
            result = calc.calculate_arithmetic_expression(expression)
            print(f"= {result}")
        except KeyboardInterrupt:
            print("\nExiting.")
            break
        except Exception as e:
            print("Error:", e)

if __name__ == "__main__":
    main()

