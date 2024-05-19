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
fn find_solution_expressions(nums: Vec<u32>, expressions: &mut HashSet<String>) {
    let mut stack = vec![(nums, String::new())];

    while !stack.is_empty() {
        let (current_nums, expr) = stack.pop().unwrap();
        if current_nums.len() == 1 {
            if current_nums[0] == 24 {
                expressions.insert(expr);
            }
            continue;
        }

        for (i, &num1) in current_nums.iter().enumerate() {
            for (j, &num2) in current_nums.iter().enumerate() {
                if i == j {
                    continue;
                }

                let mut next_nums: Vec<u32> = Vec::with_capacity(current_nums.len() - 1);
                for (k, &v) in current_nums.iter().enumerate() {
                    if k != i && k != j {
                        next_nums.push(v);
                    }
                }

                for (op, func) in OPERATIONS.iter() {
                    if *op == "/" && num2 == 0 {
                        continue;
                    }

                    if current_nums.len() == 2 && num1 > 24 && (*op == "*" || *op == "+") {
                        continue;
                    }

                    let result = func(num1 as f64, num2 as f64);
                    if result < 0.0 || result.is_infinite() || result.fract() != 0.0 {
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
}

fn deduplicate(expressions: HashSet<String>, unique_expressions: &mut HashSet<String>) {
    for expression in expressions {
        let mut expr_list = vec![];
        let mut buffer = String::with_capacity(2);
        for expr in expression.split(" -> ") {
            let mut op: char = '+';
            let mut num1 = String::with_capacity(2);
            let mut num2 = String::with_capacity(2);
            for sc in expr.chars() {
                if sc == '(' || sc == ')' || sc == ' ' {
                    continue;
                }
                if sc == '+' || sc == '-' || sc == '*' || sc == '/' {
                    op = sc;
                    num1 = buffer.clone();
                    buffer.clear();
                    continue;
                }
                if sc == '=' {
                    num2 = buffer.clone();
                    buffer.clear();
                    continue;
                }
                buffer.push(sc);
            }
            let result = buffer.clone();
            buffer.clear();
            let mut nums = vec![num1, num2];
            if op == '+' || op == '*' {
                nums.sort();
            }
            expr_list.push(format!("({} {} {}) = {}", nums[0], op, nums[1], result));
        }
        let expr3 = expr_list.pop().unwrap();
        unique_expressions.insert(format!("{} -> {} -> {}", expr_list[0], expr_list[1], expr3));
    }
}

#[pyfunction]
fn get_solution(nums: Vec<u32>) -> PyResult<Vec<String>> {
    let mut expressions = HashSet::new();
    let mut unique_expressions: HashSet<String> = HashSet::new();

    find_solution_expressions(nums, &mut expressions);
    deduplicate(expressions, &mut unique_expressions);

    let mut result: Vec<String> = unique_expressions.into_iter().collect();
    result.sort();
    Ok(result)
}

#[pymodule]
fn twenty_four_point_answer(_py: Python, m: &Bound<'_, PyModule>) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(get_solution, m)?)?;
    Ok(())
}