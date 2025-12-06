package main

import "fmt"

type MyStruct struct {
	SliceField []int
}

func main() {
	fmt.Println("=== 切片赋值操作 ===")
	
	// 初始化结构体
	struct1 := MyStruct{
		SliceField: []int{}, // 或 nil
	}
	fmt.Printf("初始状态: %v, len=%d, cap=%d\n", 
		struct1.SliceField, len(struct1.SliceField), cap(struct1.SliceField))
	
	// 方式 1: 使用 append 追加元素（最常用）
	struct1.SliceField = append(struct1.SliceField, 1)
	struct1.SliceField = append(struct1.SliceField, 2, 3, 4)  // 可以一次追加多个
	fmt.Printf("append 后: %v\n", struct1.SliceField)  // [1 2 3 4]
	
	// 方式 2: 通过索引赋值（需要先确保有元素）
	// 如果切片为空，需要先 make 或 append
	struct2 := MyStruct{
		SliceField: make([]int, 5),  // 创建长度为 5 的切片
	}
	struct2.SliceField[0] = 10
	struct2.SliceField[1] = 20
	struct2.SliceField[2] = 30
	fmt.Printf("索引赋值: %v\n", struct2.SliceField)  // [10 20 30 0 0]
	
	// 方式 3: 整体赋值
	struct3 := MyStruct{
		SliceField: []int{},
	}
	struct3.SliceField = []int{100, 200, 300}
	fmt.Printf("整体赋值: %v\n", struct3.SliceField)
	
	// 方式 4: 从另一个切片赋值
	otherSlice := []int{5, 6, 7, 8}
	struct3.SliceField = otherSlice
	fmt.Printf("从变量赋值: %v\n", struct3.SliceField)
	
	// 方式 5: 修改指定索引的值
	struct3.SliceField[0] = 999
	fmt.Printf("修改索引0: %v\n", struct3.SliceField)  // [999 6 7 8]
	
	// 方式 6: 使用循环赋值
	struct4 := MyStruct{
		SliceField: make([]int, 5),
	}
	for i := 0; i < 5; i++ {
		struct4.SliceField[i] = i * 10
	}
	fmt.Printf("循环赋值: %v\n", struct4.SliceField)  // [0 10 20 30 40]
	
	fmt.Println("\n=== 切片删除操作 ===")
	
	// 准备测试数据
	struct5 := MyStruct{
		SliceField: []int{1, 2, 3, 4, 5, 6, 7, 8},
	}
	fmt.Printf("删除前: %v, len=%d\n", struct5.SliceField, len(struct5.SliceField))
	
	// 方式 1: 删除指定索引的元素（保持顺序）
	// 删除索引 2 的元素（值为 3）
	index := 2
	struct5.SliceField = append(struct5.SliceField[:index], struct5.SliceField[index+1:]...)
	fmt.Printf("删除索引2后: %v, len=%d\n", struct5.SliceField, len(struct5.SliceField))  // [1 2 4 5 6 7 8]
	
	// 方式 2: 删除第一个元素
	struct6 := MyStruct{
		SliceField: []int{1, 2, 3, 4, 5},
	}
	struct6.SliceField = struct6.SliceField[1:]  // 从索引 1 开始
	fmt.Printf("删除第一个: %v\n", struct6.SliceField)  // [2 3 4 5]
	
	// 方式 3: 删除最后一个元素
	struct7 := MyStruct{
		SliceField: []int{1, 2, 3, 4, 5},
	}
	struct7.SliceField = struct7.SliceField[:len(struct7.SliceField)-1]
	fmt.Printf("删除最后一个: %v\n", struct7.SliceField)  // [1 2 3 4]
	
	// 方式 4: 删除指定值的元素（删除第一个匹配的）
	struct8 := MyStruct{
		SliceField: []int{1, 2, 3, 2, 4, 2, 5},
	}
	valueToDelete := 2
	struct8.SliceField = deleteValue(struct8.SliceField, valueToDelete)
	fmt.Printf("删除值2（第一个）: %v\n", struct8.SliceField)  // [1 3 2 4 2 5]
	
	// 方式 5: 删除所有匹配的值
	struct9 := MyStruct{
		SliceField: []int{1, 2, 3, 2, 4, 2, 5},
	}
	valueToDeleteAll := 2
	struct9.SliceField = deleteAllValues(struct9.SliceField, valueToDeleteAll)
	fmt.Printf("删除所有值2: %v\n", struct9.SliceField)  // [1 3 4 5]
	
	// 方式 6: 删除多个索引（从后往前删除，避免索引变化）
	struct10 := MyStruct{
		SliceField: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	indicesToDelete := []int{2, 5, 7}  // 要删除的索引
	struct10.SliceField = deleteIndices(struct10.SliceField, indicesToDelete)
	fmt.Printf("删除多个索引[2,5,7]: %v\n", struct10.SliceField)
	
	// 方式 7: 清空切片
	struct11 := MyStruct{
		SliceField: []int{1, 2, 3, 4, 5},
	}
	struct11.SliceField = []int{}  // 或 nil
	fmt.Printf("清空后: %v, len=%d, is nil: %v\n", 
		struct11.SliceField, len(struct11.SliceField), struct11.SliceField == nil)
	
	// 方式 8: 使用 copy 删除（更安全，避免内存泄漏）
	struct12 := MyStruct{
		SliceField: []int{1, 2, 3, 4, 5},
	}
	indexToDelete := 2
	newSlice := make([]int, len(struct12.SliceField)-1)
	copy(newSlice, struct12.SliceField[:indexToDelete])
	copy(newSlice[indexToDelete:], struct12.SliceField[indexToDelete+1:])
	struct12.SliceField = newSlice
	fmt.Printf("使用copy删除索引2: %v\n", struct12.SliceField)  // [1 2 4 5]
	
	fmt.Println("\n=== 完整示例：赋值和删除 ===")
	
	// 创建并赋值
	example := MyStruct{
		SliceField: []int{},
	}
	
	// 1. 追加元素
	example.SliceField = append(example.SliceField, 10, 20, 30, 40, 50)
	fmt.Printf("追加后: %v\n", example.SliceField)
	
	// 2. 修改元素
	example.SliceField[0] = 100
	fmt.Printf("修改索引0: %v\n", example.SliceField)
	
	// 3. 删除索引 2
	example.SliceField = append(example.SliceField[:2], example.SliceField[3:]...)
	fmt.Printf("删除索引2后: %v\n", example.SliceField)
	
	// 4. 继续追加
	example.SliceField = append(example.SliceField, 60, 70)
	fmt.Printf("再次追加: %v\n", example.SliceField)
	
	// 5. 删除最后一个
	example.SliceField = example.SliceField[:len(example.SliceField)-1]
	fmt.Printf("删除最后一个: %v\n", example.SliceField)
	
	fmt.Println("\n=== 性能提示 ===")
	fmt.Println("1. append 是最常用的追加方式，性能好")
	fmt.Println("2. 删除元素时，从后往前删除可以避免索引变化问题")
	fmt.Println("3. 如果频繁删除，考虑使用其他数据结构（如链表）")
	fmt.Println("4. 删除后，底层数组可能不会立即释放，需要重新切片")
}

// 辅助函数：删除第一个匹配的值
func deleteValue(slice []int, value int) []int {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// 辅助函数：删除所有匹配的值
func deleteAllValues(slice []int, value int) []int {
	result := []int{}
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

// 辅助函数：删除多个索引（从后往前删除）
func deleteIndices(slice []int, indices []int) []int {
	// 创建索引映射，方便查找
	indexMap := make(map[int]bool)
	for _, idx := range indices {
		indexMap[idx] = true
	}
	
	result := []int{}
	for i, v := range slice {
		if !indexMap[i] {
			result = append(result, v)
		}
	}
	return result
}
