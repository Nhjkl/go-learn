package main

import (
	"math"
	"testing"
	"unsafe"
)

func TestBasicType(t *testing.T) {
	t.Logf("-------- 基础数据类型 --------")

	// 布尔类型

	isTrue := true

	t.Logf("isTrue : %t", isTrue)

	// 整数类型： 有符号

	var num int = 12

	t.Logf("num is : %d", num)

	// var minInt8 int8 = -int8(math.Pow(2, (8 - 1)))
	var minInt8 int8 = -1 << 7
	var maxInt8 int8 = int8(math.Pow(2, (8-1))) - 1
	// var maxInt8 int8 = 1<<7 - 1

	t.Logf("minInt8 is : %d", minInt8)
	t.Logf("maxInt8 is : %d", maxInt8)

	var minInt16 int16 = -1 << 15
	var maxInt16 int16 = 1<<15 - 1

	t.Logf("minInt16 is : %d", minInt16)
	t.Logf("maxInt16 is : %d", maxInt16)

	var minInt32 int32 = -1 << 31
	var maxInt32 int32 = 1<<31 - 1

	t.Logf("minInt32 is : %d", minInt32)
	t.Logf("maxInt32 is : %d", maxInt32)

	var minInt64 int64 = -1 << 63
	var maxInt64 int64 = 1<<63 - 1

	t.Logf("minInt64 is : %d", minInt64)
	t.Logf("maxInt64 is : %d", maxInt64)

	// 无符号
	var maxUint8 uint8 = 1<<8 - 1
	t.Logf("maxUint8 is : %d", maxUint8)
	var maxUint16 uint16 = 1<<16 - 1
	t.Logf("maxUint16 is : %d", maxUint16)
	var maxUint32 uint32 = 1<<32 - 1
	t.Logf("maxUint32 is : %d", maxUint32)
	var maxUint64 uint64 = 1<<64 - 1
	t.Logf("maxUint64 is : %d", maxUint64)

	// 类型别名
	var b byte = 65  // byte 是 uint8 的别名
	var r rune = '中' // rune 是 int32 的别名
	t.Logf("byte: %d (%c), rune: %d (%c)\n", b, b, r, r)

	// 显示类型占用的内存大小
	t.Logf("int8 size: %d bytes\n", unsafe.Sizeof(maxInt8))
	t.Logf("int16 size: %d bytes\n", unsafe.Sizeof(maxInt16))
	t.Logf("int32 size: %d bytes\n", unsafe.Sizeof(maxInt32))
	t.Logf("int64 size: %d bytes\n", unsafe.Sizeof(maxInt64))
	t.Logf("bool size: %d bytes\n", unsafe.Sizeof(isTrue))

	// 浮点型
	var price float32 = 99.99
	var pi float64 = 3.14159265359
	t.Logf("\nfloat32: %.2f, float64: %.10f\n", price, pi)
	t.Logf("float32 size: %d bytes, float64 size: %d bytes\n", unsafe.Sizeof(price), unsafe.Sizeof(pi))

	// 字符串
	var name string = "Golang"
	greeting := "Hello World" // 类型推断
	t.Logf("string: %s\n", name)
	t.Logf("类型推断: %s\n", greeting)

	// 字符串操作
	str := "Hello 世界"
	t.Logf("字符串长度: %d (字节数)\n", len(str)) // 注意：len返回字节数，不是字符数
	firstByte := str[0]
	t.Logf("第一个字节: %d ('%c')\n", firstByte, firstByte)

	// 原始字符串字面量
	raw := `这是
一个多行
字符串`
	t.Log("\n原始字符串:")
	t.Log(raw)

	// 字符（使用rune）
	var char rune = '中'
	t.Logf("rune: %c - %d", char, char)

	// 复数类型
	var c1 complex64 = 1 + 2i
	var c2 complex128 = 3.14 + 6.28i
	c3 := complex(5.0, 10.0) // 使用complex函数创建

	t.Logf("complex64: %v, 实部: %.1f, 虚部: %.1f\n", c1, real(c1), imag(c1))
	t.Logf("complex128: %.2f, 实部: %.2f, 虚部: %.2f\n", c2, real(c2), imag(c2))
	t.Logf("使用complex函数: %v\n", c3)
	t.Logf("complex64 size: %d bytes, complex128 size: %d bytes\n", unsafe.Sizeof(c1), unsafe.Sizeof(c2))

	// 这两个值的二进制表示完全相同！
	var signed int8 = -128
	var unsigned uint8 = 128

	t.Logf("有符号 int8 的 -128：\n")
	t.Logf("  十进制: %d\n", signed)
	t.Logf("  二进制: %08b\n", uint8(signed))

	t.Logf("\n无符号 uint8 的 128：\n")
	t.Logf("  十进制: %d\n", unsigned)
	t.Logf("  二进制: %08b\n", unsigned)

	t.Log("\n关键理解：")
	t.Log("  • 二进制表示相同：都是 10000000")
	t.Log("  • 类型决定了如何解读这个二进制")
	t.Log("  • int8 的最高位被看作符号位（1表示负数）")
	t.Log("  • uint8 的所有位都被看作数值（没有符号位）")
	t.Log("  • 因此编译器通过类型声明知道应该用哪种方式解读")
}

func TestArray(t *testing.T) {
	t.Log("------------- 数组 ----------------")

	// 声明数组
	var arr [5]int
	t.Logf("arr: %v", arr)

	// 初始化数组
	arr = [5]int{1, 2, 3, 4, 5}

	t.Logf("arr: %v", arr)
	t.Log(arr[0])

	// 自动推断长度
	arr2 := [...]int{1, 2, 3, 4, 5, 6}

	t.Logf("arr2: %v", arr2)
}

func TestSlice(t *testing.T) {
	t.Log("---------------- 切片 ----------------")

	var slice []int

	t.Logf("slice : %v", slice)
	t.Logf("slice : %v", slice == nil)

	// 使用make创建切片
	slice = make([]int, 3, 5) // 长度3，容量5 什么叫容量5，就是初始化的slice，也就是这个动态数组的当前最大容量是5个，如果超过会自动扩容
	t.Logf("make创建: %v, 长度: %d, 容量: %d\n", slice, len(slice), cap(slice))

	// // 直接初始化
	slice = []int{1, 2, 3, 4, 5}
	t.Logf("初始化: %v\n", slice)

	// // 追加元素
	slice = append(slice, 6)
	t.Logf("追加后: %v, 容量: %d\n", slice, cap(slice))

	// // 切片截取
	subSlice := slice[1:3]
	t.Logf("切片[1:3]: %v\n", subSlice)

	// // 切片共享底层数组， 切片后的subSlice，和slice共享同一个底层数组，所以修改subSlice会影响slice
	subSlice[0] = 999
	t.Logf("修改subSlice后: %v\n", slice)
}

func TestMap(t *testing.T) {
	t.Log("----------------- Map -----------------")
	var m map[string]int
	t.Logf("m : %v", m)
	t.Logf("m == nil -> %v", m == nil)
	// m["test"] = 1

	// 使用make初始化
	m = make(map[string]int)
	m["apple"] = 5
	m["banana"] = 3
	t.Logf("添加后: %v\n", m)

	// 直接初始化
	m2 := map[string]int{
		"apple":  10,
		"banana": 5,
		"orange": 8,
	}
	t.Logf("直接初始化: %v\n", m2)

	// 读取值
	value := m2["apple"]
	t.Logf("apple的值: %d\n", value)
	value2 := m2["orange"]
	t.Logf("orange的值: %d\n", value2)
	m2["orange"] = 100
	t.Logf("orange的值: %d\n", m2["orange"])

	// 检查key是否存在
	value, ok := m2["grape"]
	t.Logf("grape存在: %t, 值: %d\n", ok, value)

	valueA, okA := m2["apple"]
	t.Logf("apple存在: %t, 值: %d\n", okA, valueA)

	delete(m2, "banana")
	t.Logf("删除后: %v\n", m2)

	t.Log("遍历映射:")
	for key, value := range m2 {
		t.Logf("  %s: %d\n", key, value)
	}
}

func TestPointer(t *testing.T) {
	t.Log("------------ Pointer ----------")

	x := 10

	t.Logf("x : %d", x)

	p := &x

	t.Logf("x的地址: %p\n", p)
	t.Logf("指针的值: %d\n", *p)

	m := p
	t.Logf("m: %p\n", m)
	t.Logf("m指向的值: %d\n", *m)

	// 通过指针修改值
	*p = 20
	t.Logf("修改后x的值为: %d\n", x)

	// // 指针作为函数参数
	increment := func(p *int) {
		*p++
	}
	increment(&x)
	t.Logf("函数修改后x的值为: %d\n", x)

	// // nil指针
	var ptr *int
	t.Logf("nil指针: %v\n", ptr)
	// 解引用nil指针会导致panic
	// fmt.Println(*ptr) // 这行会panic
	//

	// // ===== 值传递 vs 指针传递 =====
	t.Log("=== 值传递 vs 指针传递 ===")

	// 值传递示例
	a := 10
	t.Logf("调用前 a = %d\n", a)
	valuePass := func(num int) {
		t.Logf("  函数内接收到的值: %d\n", num)
		num = 100 // 修改副本，不影响原值
		t.Logf("  函数内修改后: %d\n", num)
	}
	valuePass(a)
	t.Logf("调用后 a = %d (值未改变)\n", a)

	// 指针传递示例
	b := 10
	t.Logf("\n调用前 b = %d\n", b)
	pointerPass := func(num *int) {
		t.Logf("  函数内接收到的地址: %p\n", num)
		t.Logf("  函数内接收到的值: %d\n", *num)
		*num = 100 // 通过指针修改原值
		t.Logf("  函数内修改后: %d\n", *num)
	}
	pointerPass(&b)
	t.Logf("调用后 b = %d (值已改变)\n", b)

	// // 详细说明
	t.Log("关键区别：")
	t.Log("1. 值传递：函数接收的是值的副本，修改副本不影响原值")
	t.Log("2. 指针传递：函数接收的是地址，通过地址可以直接修改原值")
}

func TestDemonstrateSliceGrowth(t *testing.T) {
	var s []int
	t.Log("开始扩容演示：")

	for i := 0; i < 100; i++ {
		oldCap := cap(s)
		// t.Log(oldCap)
		s = append(s, i)
		newCap := cap(s)

		if newCap != oldCap {
			t.Logf("添加元素 %d: 长度=%d, 容量 %d → %d (扩容！)\n",
				i, len(s), oldCap, newCap)
		} else {
			t.Logf("添加元素 %d: 长度=%d, 容量=%d (未扩容)\n",
				i, len(s), cap(s))
		}
	}

	t.Log("发生扩容是，再原来的cap上乘以2")
}
