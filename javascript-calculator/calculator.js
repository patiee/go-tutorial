class Calculator {
    constructor() {}

    getExpression(expressionString) {
        let ex = [];
        let value = "";

        for (let ch of expressionString) {
            if (ch === ' ') {
                continue;
            }
            if ('()^+-*/'.includes(ch)) {
                if (value !== "") {
                    let number = parseFloat(value);
                    if (isNaN(number)) {
                        throw new Error(`Could not parse ${value} into float`);
                    }
                    ex.push(number);
                    value = "";
                }
                ex.push(ch);
            } else {
                value += ch;
            }
        }

        if (value !== "") {
            let number = parseFloat(value);
            if (isNaN(number)) {
                throw new Error(`Could not parse ${value} into float`);
            }
            ex.push(number);
        }

        return ex;
    }

    getPrecedence(operation) {
        if (operation === '^') {
            return 3;
        } else if ('*/'.includes(operation)) {
            return 2;
        } else if ('+-'.includes(operation)) {
            return 1;
        } else {
            return 0;
        }
    }

    getAssociativity(operation) {
        if (operation === '^') {
            return 'R';
        }
        return 'L';
    }

    infixToPostfix(ex) {
        let result = [];
        let stack = [];

        for (let value of ex) {
            if (typeof value === 'number') {
                result.push(value);
            } else if (value === '(') {
                stack.push(value);
            } else if (value === ')') {
                while (stack.length > 0 && stack[stack.length - 1] !== '(') {
                    result.push(stack.pop());
                }
                if (stack.length === 0) {
                    throw new Error("Invalid expression, mismatched brackets");
                }
                stack.pop();
            } else {
                while (stack.length > 0 && stack[stack.length - 1] !== '(' &&
                    (this.getPrecedence(stack[stack.length - 1]) > this.getPrecedence(value) ||
                        (this.getPrecedence(stack[stack.length - 1]) === this.getPrecedence(value) &&
                            this.getAssociativity(value) === 'L'))) {
                    result.push(stack.pop());
                }
                stack.push(value);
            }
        }

        while (stack.length > 0) {
            if (stack[stack.length - 1] === '(') {
                throw new Error("Invalid expression, mismatched brackets");
            }
            result.push(stack.pop());
        }

        return result;
    }

    calculateArithmeticExpression(expressionString) {
        let start = process.hrtime.bigint();
        let ex = this.getExpression(expressionString);
        ex = this.infixToPostfix(ex);

        let stack = [];
        for (let value of ex) {
            if (value === '+') {
                if (stack.length < 2) {
                    throw new Error("Invalid expression: not enough operands for +");
                }
                let operand2 = stack.pop();
                let operand1 = stack.pop();
                stack.push(operand1 + operand2);
            } else if (value === '-') {
                if (stack.length < 2) {
                    throw new Error("Invalid expression: not enough operands for -");
                }
                let operand2 = stack.pop();
                let operand1 = stack.pop();
                stack.push(operand1 - operand2);
            } else if (value === '*') {
                if (stack.length < 2) {
                    throw new Error("Invalid expression: not enough operands for *");
                }
                let operand2 = stack.pop();
                let operand1 = stack.pop();
                stack.push(operand1 * operand2);
            } else if (value === '/') {
                if (stack.length < 2) {
                    throw new Error("Invalid expression: not enough operands for /");
                }
                let operand2 = stack.pop();
                let operand1 = stack.pop();
                stack.push(operand1 / operand2);
            } else if (value === '^') {
                if (stack.length < 2) {
                    throw new Error("Invalid expression: not enough operands for ^");
                }
                let operand2 = stack.pop();
                let operand1 = stack.pop();
                stack.push(Math.pow(operand1, operand2));
            } else {
                stack.push(value);
            }
        }

        if (stack.length !== 1) {
            throw new Error("Invalid expression: leftover operands");
        }

        let end = process.hrtime.bigint();
        console.log("Finished, took", (end - start).toString(), "nanoseconds");
        return stack[0];
    }
}

function main() {
    const calc = new Calculator();
    
    const readline = require('readline').createInterface({
        input: process.stdin,
        output: process.stdout
    });

    readline.setPrompt('arithmetic expression: ');

    readline.on('line', (expression) => {
        try {
            const result = calc.calculateArithmeticExpression(expression);
            console.log(`= ${result}`);
        } catch (error) {
            console.error("Error:", error.message);
        } finally {
            readline.prompt();
        }
    });

    readline.prompt();
}

if (require.main === module) {
    main();
}
