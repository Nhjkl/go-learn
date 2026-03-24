package main

import (
	"fmt"
	"strings"
	"testing"
)

func multiply2(a, b, c int) int {
	return a * b * c
}

func curryMultiply(a int) func(int) func(int) int {
	return func(b int) func(int) int {
		return func(c int) int {
			return a * b * c
		}
	}
}

func createLogger(prefix string) func(string) string {
	return func(text string) string {
		return fmt.Sprintf("%s : %s", prefix, text)
	}
}

func repeatString(n int) func(string) string {
	return func(s string) string {
		return strings.Repeat(s, n)
	}
}

// 创建一个比较函数
func greaterThan(threshold int) func(int) bool {
	return func(value int) bool {
		return value > threshold
	}
}

// 创建一个范围检查函数
func inRange(min, max int) func(int) bool {
	return func(value int) bool {
		return value >= min && value <= max
	}
}

// 部分应用示例
func partialAdd(a, b int) func(int) int {
	sum := a + b // 先计算前两个参数的和
	return func(c int) int {
		return sum + c
	}
}

func TestCurrying(t *testing.T) {
	t.Log("--------- currying ----------")

	addFive := curryAdd(5)

	t.Logf("addFive(3) = %d", addFive(3))
	t.Logf("addFive(10) = %d", addFive(10))
	t.Logf("curryAdd(1)(2) = %d", curryAdd(1)(2))
	t.Logf("multiply2(2,2,2) = %d", multiply2(2, 2, 2))

	cm1 := curryMultiply(2) // 2 * b * c
	cm2 := cm1(2)           // 2 * 2 * c
	cm3 := cm2(2)           // 2 * 2 * 2
	t.Logf("cm3 = %d", cm3)

	t.Logf("curryMultiply(2)(2)(2) = %d", curryMultiply(2)(2)(2))

	// 创建不同前缀的日志函数
	infoLog := createLogger("INFO")
	warnLog := createLogger("WARN")
	errorLog := createLogger("ERROR")
	t.Log(infoLog("应用程序启动"))
	t.Log(warnLog("内存使用率较高"))
	t.Log(errorLog("数据库连接失败"))

	repeat3Times := repeatString(8)
	t.Logf("repeat3Times(\"Go\") = %s\n", repeat3Times("Go"))

	t.Log("--- 示例5: 条件判断柯里化 ---")

	isGreaterThan10 := greaterThan(10)
	t.Logf("isGreaterThan10(5) = %v\n", isGreaterThan10(5))
	t.Logf("isGreaterThan10(15) = %v\n", isGreaterThan10(15))

	isInValidRange := inRange(1, 100)
	t.Logf("isInValidRange(50) = %v\n", isInValidRange(50))
	t.Logf("isInValidRange(150) = %v\n", isInValidRange(150))

	// ==================== 示例6: 部分应用 ====================
	t.Log("--- 示例6: 部分应用 vs 柯里化 ---")
	partialAdd10 := partialAdd(5, 5) // 先计算 5+5=10
	t.Logf("partialAdd(5,5)(3) = %d\n", partialAdd10(3))
	t.Log()

	// ==================== 柯里化的优势 ====================
	t.Log("========== 柯里化的优势 ==========")
	t.Log("1. 函数复用：可以创建预配置的函数")
	t.Log("2. 代码组织：将复杂函数拆分为多个步骤")
	t.Log("3. 灵活性：可以按需组合不同的函数")
	t.Log("4. 可读性：代码更清晰，意图更明确")
}
