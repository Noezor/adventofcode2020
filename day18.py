from pathlib import Path
from typing import List

input_file = Path.cwd()/"inputs"/"day18.input"

ADD = "+"
MULTIPLY = "*"

def part_a(expressions :List[str]) -> int:
    results = [resolve("(" + expr + ")", evaluate_expression_a) for expr in expressions]
    return sum(results)

def part_b(expressions :List[str]) -> int:
    results = [resolve("(" + expr + ")", evaluate_expression_b) for expr in expressions]
    return sum(results)

def resolve(expression_string : str, evaluate_expression) -> int:
    beg_parenthesis = [i for i in range(len(expression_string)) if expression_string[i] == "("]
    while beg_parenthesis != []:
        id_last_open = beg_parenthesis[-1]
        id_corresponding_closed = [i for i in range(len(expression_string)) if expression_string[i] == ")" and i > id_last_open][0]

        expression_parenthesis = expression_string[(id_last_open+1):id_corresponding_closed].split(" ")
        expression_string = expression_string[:id_last_open] + str(evaluate_expression(expression_parenthesis)) + expression_string[(id_corresponding_closed+1):]
        beg_parenthesis = [i for i in range(len(expression_string)) if expression_string[i] == "("]
    return int(expression_string)

def evaluate_expression_a(exp : List) -> int:
    numbers = [int(x) for x in exp[0::2]]
    operations = [x for x in exp[1:-1:2]] 
    assert len(numbers) == len(operations)+1, (numbers, operations)
    current_value = numbers[0]

    for operation,next_number in zip(operations, numbers[1:]):
        if operation == ADD:
            current_value = current_value + next_number
        if operation == MULTIPLY:
            current_value = current_value * next_number

    return current_value

def evaluate_expression_b(exp : List) -> int:
    return evaluate_expression_a(resolve_additions(exp))

def resolve_additions(exp : List) -> List:
    new_exp = []
    i = 0
    while i < len(exp):
        if exp[i] != ADD:
            new_exp.append(exp[i])
        else :
            last_number = int(new_exp.pop())
            next_number = int(exp[i+1])
            new_exp.append(last_number + next_number)
            i += 1
        i += 1
    return new_exp

def parse_file(filepath) -> List:
    with open(filepath, 'r') as f:
        strings = f.read().split("\n")
    return strings

assert evaluate_expression_a([1, "+" ,2, "*" ,3, "+" ,4, "*", 5, "+", 6]) == 71, evaluate_expression_a([1, "+" ,2, "*" ,3, "+" ,4, "*", 5, "+", 6])
assert resolve("(1 + (2 * 3) + (4 * (5 + 6)))", evaluate_expression_a) == 51
assert evaluate_expression_b([1, "+" ,2, "*" ,3, "+" ,4, "*", 5, "+", 6]) == 231, evaluate_expression_b([1, "+" ,2, "*" ,3, "+" ,4, "*", 5, "+", 6])
assert resolve("(5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)))", evaluate_expression_b) == 669060


print("Passed tests !")

print(part_a(parse_file(input_file)))
print(part_b(parse_file(input_file)))