import hashlib
from dataclasses import dataclass
import operator

ops = {
    "+": operator.add,
    "-": operator.sub,
    "*": operator.mul,
    "/": operator.truediv
}

@dataclass
class Expression:
    num1: int
    num2: int
    ops: str
    result: int

    def __str__(self):
        return f"{self.num1} {self.ops} {self.num2} = {self.result}"

    def __hash__(self):
        return int(hashlib.sha256(self.__str__().encode()).hexdigest(), 16)

    def __lt__(self, other):
        return self.result < other.result


def find_solution_expressions(nums, operations):
    """Modified version to return the actual expressions that make up the 24."""
    if len(nums) == 1:
        if abs(nums[0] - 24) < 1e-6:
            return [str(nums[0])]
        else:
            return []

    expressions = []
    for i, num1 in enumerate(nums):
        for j, num2 in enumerate(nums):
            if i != j:
                # Calculate the possible next numbers after applying an operation
                next_nums = [x for k, x in enumerate(nums) if k != i and k != j]
                for op, func in operations.items():
                    if (op != "/" or num2 != 0):  # Prevent division by zero
                        try:
                            result = func(num1, num2)
                            if float(int(result)) != result:
                                continue
                            if result < 0:
                                continue
                            result = int(result)
                            next_nums.append(result)
                            for expr in find_solution_expressions(next_nums, operations):
                                expressions.append(f"({num1} {op} {num2}) = {result} -> {expr}")
                            next_nums.pop()
                        except ZeroDivisionError:
                            continue
    return expressions


def standardize_expression(expr):
    """Standardize the expression by sorting its components."""
    # Split by '->' and take everything except the last part (which is always '24')
    components = expr.split(' -> ')[:-1]
    return sorted(components)


def generate_expression(expr):
    expr = expr.replace("(", "").replace(")", "")
    if expr.find("*") >= 0:
        op = "*"
    elif expr.find("+") >= 0:
        op = "+"
    elif expr.find("-") >= 0:
        op = "-"
    elif expr.find("/") >= 0:
        op = "/"
    else:
        raise ValueError
    expr = expr.split(f' {op} ')
    expr_tmp = expr[1].split(' = ')
    result = expr_tmp[1]
    expr[1] = expr_tmp[0]
    if op == "+" or op == "*":
        expr.sort()
    expr = Expression(int(float(expr[0])), int(float(expr[1])), op, int(float(result)))
    return expr

def deduplicate(expressions):
    """Deduplicate the list of expressions."""
    data = {}
    for expression in expressions:
        expr_full = []
        for expr in standardize_expression(expression):
            expr = generate_expression(expr)
            # if expr.ops == "-" and expr.result < 0:
            #     expr.num1, expr.num2 = expr.num2, expr.num1
            #     expr.result = -expr.result
            # if expr.ops == "*" and expr.result > 0 and expr.num1 < 0 and expr.num2 < 0:
            #     expr.num1, expr.num2 = -expr.num1, -expr.num2
            # if expr.ops == "/" and expr.result > 0 and expr.num1 < 0 and expr.num2 < 0:
            #     expr.num1, expr.num2 = -expr.num1, -expr.num2
            expr_full.append(expr)
        expr_full.sort()
        data[tuple(expr_full)] = expression
    return list(data.values())


def get_solution(nums):
    expressions = find_solution_expressions(nums, ops)
    unique_expressions = []
    for i in deduplicate(expressions):
        unique_expressions.append(i.replace(' -> 24', ''))
    return unique_expressions
