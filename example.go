package main

import "fmt"

// 示例：在结构体中声明数组和 map

// 1. 声明固定大小的数组
type StructWithArray struct {
	// 固定大小的数组
	Numbers [5]int           // 5个整数的数组
	Names   [10]string       // 10个字符串的数组
	Matrix  [3][3]int        // 二维数组（3x3）
}

// 2. 声明切片（slice，动态数组）
type StructWithSlice struct {
	// 切片（动态数组）
	Numbers []int            // 整数切片
	Names   []string         // 字符串切片
	Items   []interface{}    // 任意类型切片
}

// 3. 声明 map
type StructWithMap struct {
	// map 的基本声明
	Scores    map[string]int        // key 是 string，value 是 int
	Config    map[string]string     // key 和 value 都是 string
	Data      map[int]string        // key 是 int，value 是 string
	Complex   map[string][]int      // key 是 string，value 是 int 切片
	Nested    map[string]map[string]int  // 嵌套 map
}

// 4. 组合使用：同时包含数组、切片和 map
type ComplexStruct struct {
	// 数组
	FixedArray [5]int
	
	// 切片
	DynamicSlice []string
	
	// map
	KeyValueMap map[string]int
	
	// 其他字段
	ID   int
	Name string
}

func main() {
	// 示例 1: 使用数组
	arrStruct := StructWithArray{
		Numbers: [5]int{1, 2, 3, 4, 5},
		Names:   [10]string{"Alice", "Bob"},
		Matrix:  [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
	}
	fmt.Println("数组结构体:", arrStruct)

	// 示例 2: 使用切片
	sliceStruct := StructWithSlice{
		Numbers: []int{1, 2, 3, 4, 5},
		Names:   []string{"Alice", "Bob", "Charlie"},
	}
	fmt.Println("切片结构体:", sliceStruct)

	// 示例 3: 使用 map（需要先初始化）
	mapStruct := StructWithMap{
		Scores: make(map[string]int),
		Config: make(map[string]string),
	}
	mapStruct.Scores["Alice"] = 95
	mapStruct.Scores["Bob"] = 87
	mapStruct.Config["host"] = "localhost"
	mapStruct.Config["port"] = "8080"
	fmt.Println("Map 结构体:", mapStruct)

	// 示例 4: 使用字面量初始化 map
	mapStruct2 := StructWithMap{
		Scores: map[string]int{
			"Alice": 95,
			"Bob":   87,
		},
		Config: map[string]string{
			"host": "localhost",
			"port": "8080",
		},
	}
	fmt.Println("Map 结构体（字面量）:", mapStruct2)

	// 示例 5: 复杂结构体
	complex := ComplexStruct{
		FixedArray:   [5]int{1, 2, 3, 4, 5},
		DynamicSlice: []string{"a", "b", "c"},
		KeyValueMap: map[string]int{
			"key1": 100,
			"key2": 200,
		},
		ID:   1,
		Name: "Test",
	}
	fmt.Println("复杂结构体:", complex)
}
