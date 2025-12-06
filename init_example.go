package main

import "fmt"

// 演示数组和切片的初始化需求

type TestStruct struct {
	// 数组
	ArrayField [5]int
	
	// 切片
	SliceField []int
}

func main() {
	fmt.Println("=== 数组初始化情况 ===")
	
	// 1. 数组：不需要显式初始化，零值已可用
	var arr1 [5]int
	fmt.Printf("未初始化的数组: %v\n", arr1)  // 输出: [0 0 0 0 0]
	fmt.Printf("可以直接访问: arr1[0] = %d\n", arr1[0])  // 输出: 0
	
	// 数组在结构体中也是如此
	var struct1 TestStruct
	fmt.Printf("结构体中的数组（未初始化）: %v\n", struct1.ArrayField)  // [0 0 0 0 0]
	fmt.Printf("可以直接访问: struct1.ArrayField[0] = %d\n", struct1.ArrayField[0])  // 0
	
	// 可以赋值
	struct1.ArrayField[0] = 100
	struct1.ArrayField[1] = 200
	fmt.Printf("赋值后: %v\n", struct1.ArrayField)  // [100 200 0 0 0]
	
	fmt.Println("\n=== 切片初始化情况 ===")
	
	// 2. 切片：零值是 nil，但可以直接使用 append
	var slice1 []int
	fmt.Printf("未初始化的切片: %v, 是否为 nil: %v\n", slice1, slice1 == nil)  // [] true
	
	// 切片可以直接使用 append，即使为 nil
	slice1 = append(slice1, 1, 2, 3)
	fmt.Printf("append 后: %v, 是否为 nil: %v\n", slice1, slice1 == nil)  // [1 2 3] false
	
	// 切片在结构体中也是如此
	var struct2 TestStruct
	fmt.Printf("结构体中的切片（未初始化）: %v, 是否为 nil: %v\n", 
		struct2.SliceField, struct2.SliceField == nil)  // [] true
	
	// 可以直接 append
	struct2.SliceField = append(struct2.SliceField, 10, 20, 30)
	fmt.Printf("append 后: %v\n", struct2.SliceField)  // [10 20 30]
	
	// 但是，如果要通过索引访问，需要先初始化或确保有元素
	// struct2.SliceField[0] = 100  // 如果切片为空，这会 panic！
	
	// 方式 1: 使用 make 初始化
	struct3 := TestStruct{
		SliceField: make([]int, 0, 10),  // 长度为 0，容量为 10
	}
	fmt.Printf("make([]int, 0, 10): %v, len=%d, cap=%d\n", 
		struct3.SliceField, len(struct3.SliceField), cap(struct3.SliceField))
	
	// 方式 2: 使用 make 初始化并指定长度
	struct4 := TestStruct{
		SliceField: make([]int, 5),  // 长度为 5，容量也为 5，元素为 0
	}
	fmt.Printf("make([]int, 5): %v, len=%d, cap=%d\n", 
		struct4.SliceField, len(struct4.SliceField), cap(struct4.SliceField))
	struct4.SliceField[0] = 100  // 现在可以安全访问
	fmt.Printf("赋值后: %v\n", struct4.SliceField)
	
	// 方式 3: 使用字面量初始化
	struct5 := TestStruct{
		SliceField: []int{1, 2, 3},
	}
	fmt.Printf("字面量初始化: %v\n", struct5.SliceField)
	
	fmt.Println("\n=== 总结 ===")
	fmt.Println("数组：")
	fmt.Println("  - 不需要显式初始化，零值已可用")
	fmt.Println("  - 可以直接通过索引访问和赋值")
	fmt.Println("  - 大小固定，在声明时确定")
	
	fmt.Println("\n切片：")
	fmt.Println("  - 零值是 nil，但可以直接使用 append")
	fmt.Println("  - 如果要通过索引访问，需要先确保有元素（make 或字面量）")
	fmt.Println("  - 推荐：如果知道大概大小，用 make 预分配容量")
	fmt.Println("  - 推荐：如果只是追加元素，可以直接 append（nil 切片也可以）")
	
	fmt.Println("\n=== 实际使用建议 ===")
	
	// 推荐做法 1: 数组直接使用
	goodArray := TestStruct{
		ArrayField: [5]int{},  // 可以显式初始化，也可以不写（零值）
	}
	goodArray.ArrayField[0] = 1
	fmt.Printf("数组推荐用法: %v\n", goodArray.ArrayField)
	
	// 推荐做法 2: 切片如果知道容量，用 make
	goodSlice1 := TestStruct{
		SliceField: make([]int, 0, 100),  // 预分配容量，避免多次扩容
	}
	for i := 0; i < 5; i++ {
		goodSlice1.SliceField = append(goodSlice1.SliceField, i)
	}
	fmt.Printf("切片推荐用法（预分配）: %v\n", goodSlice1.SliceField)
	
	// 推荐做法 3: 切片如果不知道大小，直接用 nil 然后 append
	goodSlice2 := TestStruct{
		SliceField: nil,  // 或者不写，默认就是 nil
	}
	for i := 0; i < 5; i++ {
		goodSlice2.SliceField = append(goodSlice2.SliceField, i*10)
	}
	fmt.Printf("切片推荐用法（动态增长）: %v\n", goodSlice2.SliceField)
}
