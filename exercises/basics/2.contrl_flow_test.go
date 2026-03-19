package main

import (
	"fmt"
	"testing"
)

func TestIf(t *testing.T) {
	t.Logf("--------- if -----------")

	age := 17

	if age >= 18 {
		t.Log("已成年")
	} else {
		t.Log("未成年")
	}

	// 支持初始化语句
	if num := 10; num > 5 {
		t.Logf("num(%d)大于5", num)
	}

	// 多条件判断
	scores := []int{95, 85, 70, 50}

	for index, score := range scores {
		t.Logf("index: %d, score: %d", index, score)
		if score >= 90 {
			t.Logf("分数：%d，优秀", score)
		} else if score >= 80 {
			t.Logf("分数：%d，良好", score)
		} else if score >= 60 {
			t.Logf("分数：%d，及格", score)
		} else {
			t.Logf("分数：%d，不及格", score)
		}
	}
}

func TestSwitch(t *testing.T) {
	t.Log("----------- switch -----------")

	days := []int{1, 3, 5, 7, 10}

	for index, day := range days {
		switch day {
		case 1:
			t.Log("Monday")
		case 2:
			t.Log("Tuesday")
		case 3:
			t.Log("Wednesday")
		case 4:
			t.Log("Thursday")
		case 5:
			t.Log("Friday")
		case 6, 7:
			t.Log("周末")
		default:
			t.Logf("Day %d index %d 未知", day, index)
		}
	}

	// 条件表达式
	score := 85
	switch {
	case score >= 90:
		t.Log("等级: A")
	case score >= 80:
		t.Log("等级: B")
	case score >= 70:
		t.Log("等级: C")
	case score >= 60:
		t.Log("等级: D")
	default:
		t.Log("等级: E")
	}
}

func TestSwitchFallthrough(t *testing.T) {
	t.Log("---------- switch fallthrough --------------")

	t.Log("1. 默认行为： 不穿透")

	num := 1

	switch num {
	case 1:
		t.Log("matched 1")
	case 2:
		t.Log("matched 2")
	case 3:
		t.Log("matched 3")
	default:
		t.Log("default")
	}

	switch num {
	case 1:
		t.Log("matched 1")
		break // 显示中断，其实不需要
	case 2:
		t.Log("matched 2")
	case 3:
		t.Log("matched 3")
	default:
		t.Log("default")
	}

	switch num {
	case 1:
		t.Log("  匹配到 case 1")
		fallthrough // 继续执行下一个 case（穿透）
	case 2:
		t.Log("  执行 case 2（因为 fallthrough）")
		fallthrough
	case 3:
		t.Log("  匹配到 case 3")
	default:
		t.Log("  默认分支")
	}

	t.Log("\n5. 实际应用场景：分级判断（显示所有满足的等级）")
	scores := []int{95, 85, 75, 65}
	for _, score := range scores {
		fmt.Printf("  分数 %d: ", score)
		first := true
		switch {
		case score >= 90:
			t.Log("优秀")
			first = false
			fallthrough // 优秀也属于良好和及格
		case score >= 100:
			if !first {
				t.Log("、良好")
			} else {
				t.Log("良好")
			}
			first = false
			fallthrough // 良好也属于及格
		case score >= 60:
			if !first {
				t.Log("、及格")
			} else {
				t.Log("及格")
			}
			// 这里不 fallthrough，所以不会输出"不及格"
		default:
			t.Log("不及格")
		}
		t.Log()
	}

	t.Log("\n6. 对比：不使用 fallthrough 的正常分级判断")
	scores2 := []int{95, 85, 75, 55}
	for _, score := range scores2 {
		t.Logf("  分数 %d: ", score)
		switch {
		case score >= 90:
			t.Log("优秀")
		case score >= 80:
			t.Log("良好")
		case score >= 60:
			t.Log("及格")
		default:
			t.Log("不及格")
		}
	}
}

func TestForloop(t *testing.T) {
	t.Log("------------ for loop ----------------")

	t.Log("1. 基本for循环")
	for i := 0; i < 5; i++ {
		t.Logf("i : %d", i)
	}

	t.Log("2. 类似while循环")

	i := 0

	for i < 5 {
		t.Logf("i : %d", i)
		i++
	}

	t.Log("死循环需要满足条件跳出")

	i = 0

	for {
		if i >= 5 {
			break
		}
		t.Logf("i : %d", i)
		i++
	}

	// // range遍历数组/切片
	t.Log("range遍历数组:")
	arr := [...]int{10, 20, 30, 40, 50}

	for index, value := range arr {
		t.Logf("  索引[%d] = %d\n", index, value)
	}

	for _, value := range arr {
		t.Logf("  value = %d\n", value)
	}

	// 遍历字符串
	t.Log("遍历字符串:")
	str := "Hello 世界"
	for i, char := range str {
		t.Logf("  [%d] %c\n", i, char)
	}
}

func TestDefer(t *testing.T) {
	t.Log("=== defer示例 ===")

	// 基本defer
	// t.Log("1. 基本defer（延迟到函数返回前）:")
	// defer t.Log("World")
	// t.Log("Hello")

	// 多个defer执行顺序（LIFO）
	// t.Log("\n2. 多个defer执行顺序（LIFO，后进先出）:")
	// t.Log("函数开始")
	// defer t.Log("1号defer - 最先注册，最后执行")
	// defer t.Log("2号defer - 中间注册")
	// defer t.Log("3号defer - 最后注册，最先执行")
	// t.Log("函数结束")

	returnWithDefer := func() int {
		defer t.Log("  defer: 在return之后执行")
		t.Log("  return: 先准备返回值")
		return 42
	}

	// defer在return之后执行
	t.Log("3. defer在return之后执行:")
	t.Log(returnWithDefer())
}

func TestDeferValue(t *testing.T) {
	i := 0
	defer t.Log("  defer1: i =", i) // 立即捕获i=0
	i++
	defer t.Log("  defer2: i =", i) // 立即捕获i=1
	i++
	t.Log("  函数内: i =", i) // 输出最终值2
}

func TestDeferClosure(t *testing.T) {
	i := 0
	i++
	i++
	defer func() {
		t.Log("  defer闭包: i =", i) // 捕获变量引用，输出最终值2
	}()
	t.Log("  函数内: i =", i) // 输出: 2
	// defer执行时，闭包中捕获的是i的引用，所以输出最终值2
}

func TestDeferPanic(t *testing.T) {
	defer func() {
		t.Log("  清理工作: defer在panic之后执行")
	}()
	t.Log("  开始执行")
	t.Log("  即将panic...")
	// 注意：这里不真正panic，否则会终止程序
	// 实际应用中可以在defer中使用recover捕获panic
	t.Log("  (演示结束，实际panic会执行defer)")
}

// panic和recover示例
func TestPanicRecover(t *testing.T) {
	safeOperation := func() {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("  捕获panic: %v\n", r)
				t.Log("  程序继续执行")
			}
		}()

		t.Log("  即将触发panic")
		panic("发生错误")
		t.Log("  这行不会执行")
	}

	riskyOperation := func() {
		t.Log("这会导致程序崩溃")
		panic("致命错误")
		t.Log("这行不会执行")
	}

	t.Log("=== panic和recover示例 ===")

	t.Log("触发panic但使用recover捕获:")
	safeOperation()

	t.Log("正常执行panic:")
	riskyOperation()
}

// func readFile(filename string) error {
// 	defer func() {
// 		fmt.Printf("  清理资源: %s\n", filename)
// 	}()
//
// 	fmt.Printf("  准备打开: %s\n", filename)
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
// 	defer fmt.Printf("  关闭文件: %s\n", filename)
//
// 	fmt.Printf("  成功打开: %s\n", filename)
// 	return nil
// }
