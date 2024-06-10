use std::f64::consts;
use std::io::{self, Write};
use std::str::FromStr;
use std::time::{Duration, Instant};

#[derive(Debug)]
enum Token {
    Number(f64),
    Operator(char),
    LeftParenthesis,
    RightParenthesis,
}

impl Token {
    fn associativity(&self) -> &'static str {
        match *self {
            Token::Operator('^') => "R",
            _ => "L",
        }
    }
    fn precedence(&self) -> u8 {
        match *self {
            Token::Operator('^') => 4, // Highest precedence for exponentiation
            Token::Operator('*') | Token::Operator('/') => 3, // Multiplication and division
            Token::Operator('+') | Token::Operator('-') => 2, // Addition and subtraction
            _ => 0,                    // Numbers and other characters
        }
    }
}

struct Calculator {}

impl Calculator {
    fn get_expression(&self, expression_string: &str) -> Vec<Token> {
        let mut ex = Vec::new();
        let mut value = String::new();

        for ch in expression_string.chars() {
            match ch {
                ' ' => continue,
                '(' | ')' | '^' | '+' | '-' | '*' | '/' => {
                    if !value.is_empty() {
                        if let Ok(number) = f64::from_str(&value) {
                            ex.push(Token::Number(number));
                        } else {
                            panic!("Could not parse {} into float64", value);
                        }
                        value.clear();
                    }
                    ex.push(match ch {
                        '(' => Token::LeftParenthesis,
                        ')' => Token::RightParenthesis,
                        _ => Token::Operator(ch),
                    });
                }
                _ => value.push(ch),
            }
        }

        if !value.is_empty() {
            if let Ok(number) = f64::from_str(&value) {
                ex.push(Token::Number(number));
            } else {
                panic!("Could not parse {} into float64", value);
            }
        }

        ex
    }

    fn infix_to_postfix(&self, ex: Vec<Token>) -> Vec<Token> {
        let mut result = Vec::new();
        let mut stack = Vec::new();

        for value in ex {
            match value {
                Token::Number(_) => result.push(value),
                Token::LeftParenthesis => stack.push(value),
                Token::RightParenthesis => {
                    while let Some(op) = stack.pop() {
                        if let Token::LeftParenthesis = op {
                            break;
                        }
                        result.push(op);
                    }
                }
                Token::Operator(op) => {
                    while let Some(top_op) = stack.last() {
                        if let Token::LeftParenthesis = top_op {
                            break;
                        }
                        if (top_op.precedence() > value.precedence())
                            || (top_op.precedence() == value.precedence()
                                && value.associativity() == "L")
                        {
                            result.push(stack.pop().unwrap());
                        } else {
                            break;
                        }
                    }
                    stack.push(Token::Operator(op));
                }
            }
        }

        while let Some(op) = stack.pop() {
            result.push(op);
        }

        result
    }

    fn calculate_arithmetic_expression(&self, expression_string: &str) -> f64 {
        let start_time = Instant::now();
        let ex = self.get_expression(expression_string);
        let ex = self.infix_to_postfix(ex);

        let mut stack = Vec::new();
        for value in ex {
            match value {
                Token::Number(num) => stack.push(num),
                Token::Operator('+') => {
                    let operand2 = stack.pop().unwrap();
                    let operand1 = stack.pop().unwrap();
                    stack.push(operand1 + operand2);
                }
                Token::Operator('-') => {
                    let operand2 = stack.pop().unwrap();
                    let operand1 = stack.pop().unwrap();
                    stack.push(operand1 - operand2);
                }
                Token::Operator('*') => {
                    let operand2 = stack.pop().unwrap();
                    let operand1 = stack.pop().unwrap();
                    stack.push(operand1 * operand2);
                }
                Token::Operator('/') => {
                    let operand2 = stack.pop().unwrap();
                    let operand1 = stack.pop().unwrap();
                    stack.push(operand1 / operand2);
                }
                Token::Operator('^') => {
                    let operand2 = stack.pop().unwrap();
                    let operand1 = stack.pop().unwrap();
                    stack.push(operand1.powf(operand2));
                }
                _ => {}
            }
        }

        if stack.len() != 1 {
            panic!("Invalid expression: leftover operands");
        }

        let duration = start_time.elapsed();
        println!("Finished, took {} nanoseconds", duration.as_nanos());
        stack.pop().unwrap()
    }
}

fn main() {
    let calculator = Calculator {};

    loop {
        print!("Enter an arithmetic expression: ");
        io::stdout().flush().unwrap();

        let mut expression = String::new();
        io::stdin()
            .read_line(&mut expression)
            .expect("Failed to read line");

        let result = calculator.calculate_arithmetic_expression(&expression.trim());
        println!("= {}", result);
    }
}
