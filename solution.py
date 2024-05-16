import operator

OPERATIONS = {
    "+": operator.add,
    "-": operator.sub,
    "*": operator.mul,
    "/": operator.truediv
}


def find_solution_expressions(nums):
    stack = [(nums, "")]
    expressions = []

    while stack:
        current_nums, expr = stack.pop()
        if len(current_nums) == 1:
            if current_nums[0] == 24:
                expressions.append(expr)
            continue

        for i, num1 in enumerate(current_nums):
            for j, num2 in enumerate(current_nums):
                if i == j:
                    continue

                next_nums = [x for k, x in enumerate(current_nums) if k != i and k != j]

                for op, func in OPERATIONS.items():
                    if op == "/" and num2 == 0:
                        continue

                    result = func(num1, num2)
                    if result < 0 or result == float('inf') or result != int(result):
                        continue

                    result = int(result)
                    next_nums.append(result)

                    new_expr = f"({num1} {op} {num2}) = {result}" if not expr else f"{expr} -> ({num1} {op} {num2}) = {result}"
                    stack.append((next_nums[:], new_expr))
                    next_nums.pop()

    return expressions


def sort_expr(expr):
    expr = expr.replace("(", "").replace(")", "")
    if "+" in expr:
        op = "+"
    elif "-" in expr:
        op = "-"
    elif "*" in expr:
        op = "*"
    else:
        op = "/"
    expr_split_op = expr.split(f" {op} ")
    expr_split_eq = expr_split_op[1].split(" = ")
    nums = [expr_split_op[0], expr_split_eq[0]]
    result = expr_split_eq[1]
    if op in ["+", "*"]:
        nums.sort()
    return f"({nums[0]} {op} {nums[1]}) = {result}"


def deduplicate(expressions):
    unique_expressions = set()

    for expression in expressions:
        expr_list = []
        for expr in expression.split(" -> "):
            expr_list.append(sort_expr(expr))
        expr3 = expr_list.pop()
        unique_expressions.add(f"{expr_list[0]} -> {expr_list[1]} -> {expr3}")

    return list(unique_expressions)


def get_solution(nums):
    expressions = find_solution_expressions(nums)
    unique_expressions = deduplicate(expressions)
    unique_expressions.sort()
    return unique_expressions
