from solution24.solution24 import GetSolution


def get_solution(nums):
    solutions: list[str] = list(
        map(
            lambda solution: solution.replace("a", str(nums[0]))
            .replace("b", str(nums[1]))
            .replace("c", str(nums[2]))
            .replace("d", str(nums[3]))
            + "=24",
            GetSolution(*nums),
        )
    )

    return solutions
