package main

import (
	"errors"
	"testing"
)

func add(a int, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

// 返回值
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

// 命名返回值
func calculate(a, b int) (sum, product int) {
	sum = a + b
	product = a * b
	return // 自动返回命名的返回值 return  sum, product
}

// 可变参数
func sumFn(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 闭包
func makeCounter() func() int {
	conunt := 0

	return func() int {
		conunt++
		return conunt
	}
}

// 闭包柯里化
func curryAdd(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

// 高阶函数：函数作为参数
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func TestFunc(t *testing.T) {
	t.Log("------------ func test--------------")

	num := add(1, 1)

	t.Logf("num: %d", num)

	num = multiply(2, 2)

	t.Logf("num: %d", num)

	num1, err := divide(2, 0)
	if err != nil {
		t.Logf("num1: %d -- err: %v", num1, err)
	}

	num1, err = divide(4, 2)
	if err == nil {
		t.Logf("num1: %d -- err: %v", num1, err)
	}

	sum, product := calculate(2, 3)
	t.Logf("sum %d --- product %d", sum, product)

	total := sumFn(1, 2, 3)
	t.Logf("total: %d", total)

	c1 := makeCounter()
	c2 := makeCounter()

	t.Logf("c1: %d", c1())
	t.Logf("c1: %d", c1())
	t.Logf("c2: %d", c2())

	cu := curryAdd(0)

	t.Logf("cu : %d", cu(1))
	t.Logf("cu : %d", cu(2))

	// 高阶函数
	t.Log("=== 高阶函数 ===")
	t.Log("applyOperation(10, 20, add):", applyOperation(10, 20, add))
	t.Log("applyOperation(10, 20, multiply):", applyOperation(10, 20, multiply))
	t.Log("applyOperation(100, 5, func(a, b int) int { return a / b }):",
		applyOperation(100, 5, func(a, b int) int { return a / b }))
}
