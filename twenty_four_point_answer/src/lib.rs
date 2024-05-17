use pyo3::prelude::*;
use std::collections::{HashMap, HashSet};
use std::ops::{Add, Sub, Mul, Div};
use lazy_static::lazy_static;

lazy_static! {
    static ref OPERATIONS: HashMap<&'static str, fn(f64, f64) -> f64> = {
        let mut m = HashMap::new();
        m.insert("+", Add::add as fn(f64, f64) -> f64);
        m.insert("-", Sub::sub as fn(f64, f64) -> f64);
        m.insert("*", Mul::mul as fn(f64, f64) -> f64);
        m.insert("/", Div::div as fn(f64, f64) -> f64);
        m
    };
}
fn find_solution_expressions(nums: Vec<u32>) -> Vec<String> {
    let mut stack = vec![(nums, String::new())];
    let mut expressions = vec![];

    while let Some((current_nums, expr)) = stack.pop() {
        if current_nums.len() == 1 {
            if current_nums[0] == 24 {
                expressions.push(expr);
            }
            continue;
        }

        for (i, &num1) in current_nums.iter().enumerate() {
            for (j, &num2) in current_nums.iter().enumerate() {
                if i == j {
                    continue;
                }

                let mut next_nums: Vec<u32> = current_nums.iter()
                    .enumerate()
                    .filter(|&(k, _)| k != i && k != j)
                    .map(|(_, &x)| x)
                    .collect();

                for (op, func) in OPERATIONS.iter() {
                    if *op == "/" && num2 == 0 {
                        continue;
                    }

                    let result = func(num1 as f64, num2 as f64);
                    if result < 0.0 || result.is_infinite() || result.fract().abs() > 0.0 {
                        continue;
                    }

                    let result = result as u32;
                    next_nums.push(result);

                    let new_expr = if expr.is_empty() {
                        format!("({} {} {}) = {}", num1, op, num2, result)
                    } else {
                        format!("{} -> ({} {} {}) = {}", expr, num1, op, num2, result)
                    };

                    stack.push((next_nums.clone(), new_expr));
                    next_nums.pop();
                }
            }
        }
    }

    expressions
}

fn sort_expr(expr: &str) -> String {
    let expr = expr.replace("(", "").replace(")", "");
    let op = if expr.contains("+") {
        "+"
    } else if expr.contains("-") {
        "-"
    } else if expr.contains("*") {
        "*"
    } else {
        "/"
    };
    let expr_split_op = expr.split(&format!(" {} ", op)).collect::<Vec<&str>>();
    let expr_split_eq = expr_split_op[1].split(" = ").collect::<Vec<&str>>();
    let mut nums = vec![expr_split_op[0], expr_split_eq[0]];
    let result = expr_split_eq[1];
    if op == "+" || op == "*" {
        nums.sort();
    }
    return format!("({} {} {}) = {}", nums[0], op, nums[1], result);
}

fn deduplicate(expressions: Vec<String>) -> Vec<String> {
    let mut unique_expressions: HashSet<String> = HashSet::new();

    for expression in expressions {
        let mut expr_list = vec![];
        for expr in expression.split(" -> ") {
            expr_list.push(sort_expr(expr));
        }
        let expr3 = expr_list.pop().unwrap();
        unique_expressions.insert(format!("{} -> {} -> {}", expr_list[0], expr_list[1], expr3));
    }
    return unique_expressions.into_iter().collect();
}

#[pyfunction]
fn get_solution(nums: Vec<u32>) -> PyResult<Vec<String>> {
    let expressions = find_solution_expressions(nums);
    let mut unique_expressions = deduplicate(expressions);
    unique_expressions.sort();
    Ok(unique_expressions)
}

#[pymodule]
fn twenty_four_point_answer(_py: Python, m: &Bound<'_, PyModule>) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(get_solution, m)?)?;
    Ok(())
}