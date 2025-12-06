package main

import "fmt"

// 演示数组和切片的赋值方式

type MyStruct struct {
	// 数组（固定大小）
	ArrayField [5]int
	
	// 切片（动态大小）
	SliceField []int
}

func main() {
	fmt.Println("=== 数组赋值方式 ===")
	
	// 方式 1: 声明时直接赋值（字面量）
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("字面量赋值: %v\n", arr1)
	
	// 方式 2: 简化写法
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("简化写法: %v\n", arr2)
	
	// 方式 3: 部分赋值（未指定的为 0）
	arr3 := [5]int{1, 2}  // 只赋值前两个，后面是 0
	fmt.Printf("部分赋值: %v\n", arr3)  // [1 2 0 0 0]
	
	// 方式 4: 指定索引赋值
	arr4 := [5]int{0: 10, 2: 30, 4: 50}
	fmt.Printf("指定索引: %v\n", arr4)  // [10 0 30 0 50]
	
	// 方式 5: 逐个元素赋值
	var arr5 [5]int
	arr5[0] = 100
	arr5[1] = 200
	arr5[2] = 300
	fmt.Printf("逐个赋值: %v\n", arr5)  // [100 200 300 0 0]
	
	// 方式 6: 使用循环赋值
	var arr6 [5]int
	for i := 0; i < 5; i++ {
		arr6[i] = i * 10
	}
	fmt.Printf("循环赋值: %v\n", arr6)  // [0 10 20 30 40]
	
	// 方式 7: 数组复制（整个数组赋值）
	arr7 := [5]int{1, 2, 3, 4, 5}
	arr8 := arr7  // 复制整个数组
	arr8[0] = 999
	fmt.Printf("原数组: %v\n", arr7)  // [1 2 3 4 5]
	fmt.Printf("复制后修改: %v\n", arr8)  // [999 2 3 4 5]
	
	fmt.Println("\n=== 切片赋值方式 ===")
	
	// 方式 1: 字面量赋值
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("字面量赋值: %v\n", slice1)
	
	// 方式 2: 使用 make 创建后赋值
	slice2 := make([]int, 5)  // 长度为 5，元素为 0
	slice2[0] = 10
	slice2[1] = 20
	slice2[2] = 30
	fmt.Printf("make 后赋值: %v\n", slice2)  // [10 20 30 0 0]
	
	// 方式 3: 使用 append 追加
	slice3 := []int{}  // 空切片
	slice3 = append(slice3, 1, 2, 3)
	fmt.Printf("append 追加: %v\n", slice3)
	
	// 方式 4: 从数组创建切片
	arr9 := [5]int{10, 20, 30, 40, 50}
	slice4 := arr9[:]  // 整个数组转为切片
	slice5 := arr9[1:3]  // 索引 1 到 3（不包含 3）
	fmt.Printf("从数组创建: %v\n", slice4)  // [10 20 30 40 50]
	fmt.Printf("数组切片: %v\n", slice5)  // [20 30]
	
	// 方式 5: 使用循环赋值
	slice6 := make([]int, 5)
	for i := 0; i < 5; i++ {
		slice6[i] = i * 100
	}
	fmt.Printf("循环赋值: %v\n", slice6)  // [0 100 200 300 400]
	
	// 方式 6: 切片复制（使用 copy）
	slice7 := []int{1, 2, 3, 4, 5}
	slice8 := make([]int, len(slice7))
	copy(slice8, slice7)  // 复制切片
	slice8[0] = 999
	fmt.Printf("原切片: %v\n", slice7)  // [1 2 3 4 5]
	fmt.Printf("复制后修改: %v\n", slice8)  // [999 2 3 4 5]
	
	// 方式 7: 直接赋值整个切片
	slice9 := []int{1, 2, 3}
	slice10 := slice9  // 注意：这是引用，不是复制！
	slice10[0] = 999
	fmt.Printf("原切片: %v\n", slice9)  // [999 2 3] - 也被修改了！
	fmt.Printf("赋值后修改: %v\n", slice10)  // [999 2 3]
	
	fmt.Println("\n=== 在结构体中赋值 ===")
	
	// 数组在结构体中
	struct1 := MyStruct{
		ArrayField: [5]int{1, 2, 3, 4, 5},  // 字面量
	}
	fmt.Printf("结构体数组（字面量）: %v\n", struct1.ArrayField)
	
	// 或者先创建结构体，再赋值
	struct2 := MyStruct{}
	struct2.ArrayField[0] = 100
	struct2.ArrayField[1] = 200
	fmt.Printf("结构体数组（逐个）: %v\n", struct2.ArrayField)
	
	// 或者整个数组赋值
	arr10 := [5]int{10, 20, 30, 40, 50}
	struct2.ArrayField = arr10
	fmt.Printf("结构体数组（整体）: %v\n", struct2.ArrayField)
	
	// 切片在结构体中
	struct3 := MyStruct{
		SliceField: []int{1, 2, 3, 4, 5},  // 字面量
	}
	fmt.Printf("结构体切片（字面量）: %v\n", struct3.SliceField)
	
	// 或者先创建，再赋值
	struct4 := MyStruct{
		SliceField: make([]int, 5),
	}
	struct4.SliceField[0] = 100
	struct4.SliceField[1] = 200
	fmt.Printf("结构体切片（make 后赋值）: %v\n", struct4.SliceField)
	
	// 或者使用 append
	struct5 := MyStruct{
		SliceField: []int{},  // 或 nil
	}
	struct5.SliceField = append(struct5.SliceField, 10, 20, 30)
	fmt.Printf("结构体切片（append）: %v\n", struct5.SliceField)
	
	// 或者整个切片赋值
	slice11 := []int{100, 200, 300}
	struct5.SliceField = slice11
	fmt.Printf("结构体切片（整体赋值）: %v\n", struct5.SliceField)
	
	fmt.Println("\n=== 常见场景 ===")
	
	// 场景 1: 从函数返回值赋值
	slice12 := getNumbers()
	fmt.Printf("函数返回: %v\n", slice12)
	
	// 场景 2: 从另一个变量赋值
	value := []int{1, 2, 3, 4, 5}  // 这就是你问的 value []int
	struct6 := MyStruct{
		SliceField: value,  // 直接赋值
	}
	fmt.Printf("从变量赋值: %v\n", struct6.SliceField)
	
	// 场景 3: 修改现有切片
	struct6.SliceField[0] = 999
	fmt.Printf("修改后: %v\n", struct6.SliceField)
	
	// 场景 4: 追加元素
	struct6.SliceField = append(struct6.SliceField, 6, 7, 8)
	fmt.Printf("追加后: %v\n", struct6.SliceField)
}

// 辅助函数：返回切片
func getNumbers() []int {
	return []int{10, 20, 30}
}
