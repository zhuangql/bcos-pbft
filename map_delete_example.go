package main

import "fmt"

// 定义结构体类型
type MyStruct struct {
	Data map[string]int
}

// 为结构体添加删除方法
func (s *MyStruct) deleteKey(key string) {
	delete(s.Data, key)
}

func main() {
	fmt.Println("=== 删除 map 元素的基本方法 ===")
	
	// 创建并初始化 map
	m := map[string]int{
		"apple":  10,
		"banana": 20,
		"cherry": 30,
		"date":   40,
		"elder":  50,
	}
	fmt.Printf("删除前: %v\n", m)
	
	// 方式 1: 使用 delete 函数删除元素（最常用）
	delete(m, "banana")
	fmt.Printf("删除 'banana' 后: %v\n", m)
	
	// 删除不存在的键（不会报错，安全）
	delete(m, "nonexistent")
	fmt.Printf("删除不存在的键后: %v\n", m)
	
	// 继续删除
	delete(m, "cherry")
	fmt.Printf("删除 'cherry' 后: %v\n", m)
	
	fmt.Println("\n=== 删除前检查元素是否存在 ===")
	
	m2 := map[string]int{
		"apple":  10,
		"banana": 20,
		"cherry": 30,
	}
	
	// 检查键是否存在，存在才删除
	key := "banana"
	if _, exists := m2[key]; exists {
		delete(m2, key)
		fmt.Printf("键 '%s' 存在，已删除: %v\n", key, m2)
	} else {
		fmt.Printf("键 '%s' 不存在\n", key)
	}
	
	// 检查并删除（更简洁的写法）
	key2 := "nonexistent"
	if _, ok := m2[key2]; ok {
		delete(m2, key2)
	} else {
		fmt.Printf("键 '%s' 不存在，无需删除\n", key2)
	}
	
	fmt.Println("\n=== 删除多个元素 ===")
	
	m3 := map[string]int{
		"apple":  10,
		"banana": 20,
		"cherry": 30,
		"date":   40,
		"elder":  50,
	}
	fmt.Printf("删除前: %v\n", m3)
	
	// 方式 1: 逐个删除
	keysToDelete := []string{"banana", "date"}
	for _, key := range keysToDelete {
		delete(m3, key)
	}
	fmt.Printf("删除多个键后: %v\n", m3)
	
	// 方式 2: 删除满足条件的元素
	m4 := map[string]int{
		"apple":  10,
		"banana": 20,
		"cherry": 30,
		"date":   40,
		"elder":  50,
	}
	fmt.Printf("删除前: %v\n", m4)
	
	// 删除值小于 30 的元素
	for key, value := range m4 {
		if value < 30 {
			delete(m4, key)
		}
	}
	fmt.Printf("删除值 < 30 的元素后: %v\n", m4)
	
	fmt.Println("\n=== 清空 map ===")
	
	m5 := map[string]int{
		"apple":  10,
		"banana": 20,
		"cherry": 30,
	}
	fmt.Printf("清空前: %v, len=%d\n", m5, len(m5))
	
	// 方式 1: 重新分配（推荐）
	m5 = make(map[string]int)
	fmt.Printf("清空后: %v, len=%d\n", m5, len(m5))
	
	// 方式 2: 重新赋值为空 map 字面量
	m6 := map[string]int{
		"apple":  10,
		"banana": 20,
	}
	m6 = map[string]int{}
	fmt.Printf("清空后: %v, len=%d\n", m6, len(m6))
	
	// 方式 3: 循环删除所有元素（不推荐，效率低）
	m7 := map[string]int{
		"apple":  10,
		"banana": 20,
		"cherry": 30,
	}
	for key := range m7 {
		delete(m7, key)
	}
	fmt.Printf("循环删除后: %v, len=%d\n", m7, len(m7))
	
	fmt.Println("\n=== 在结构体中使用 map ===")
	
	struct1 := MyStruct{
		Data: map[string]int{
			"apple":  10,
			"banana": 20,
			"cherry": 30,
		},
	}
	fmt.Printf("结构体中的 map: %v\n", struct1.Data)
	
	// 删除结构体中的 map 元素
	delete(struct1.Data, "banana")
	fmt.Printf("删除后: %v\n", struct1.Data)
	
	// 通过方法删除
	struct1.deleteKey("cherry")
	fmt.Printf("通过方法删除后: %v\n", struct1.Data)
	
	fmt.Println("\n=== 删除并获取值 ===")
	
	m8 := map[string]int{
		"apple":  10,
		"banana": 20,
		"cherry": 30,
	}
	
	// 删除前先获取值
	keyToDelete := "banana"
	if value, exists := m8[keyToDelete]; exists {
		fmt.Printf("删除键 '%s'，其值为: %d\n", keyToDelete, value)
		delete(m8, keyToDelete)
	}
	fmt.Printf("删除后: %v\n", m8)
	
	fmt.Println("\n=== 删除嵌套 map 中的元素 ===")
	
	nestedMap := map[string]map[string]int{
		"fruits": {
			"apple":  10,
			"banana": 20,
		},
		"vegetables": {
			"carrot": 15,
			"potato": 25,
		},
	}
	fmt.Printf("嵌套 map 删除前: %v\n", nestedMap)
	
	// 删除嵌套 map 中的元素
	if fruits, exists := nestedMap["fruits"]; exists {
		delete(fruits, "banana")
	}
	fmt.Printf("删除嵌套元素后: %v\n", nestedMap)
	
	// 删除整个嵌套 map
	delete(nestedMap, "vegetables")
	fmt.Printf("删除整个嵌套 map 后: %v\n", nestedMap)
	
	fmt.Println("\n=== 安全删除函数封装 ===")
	
	m9 := map[string]int{
		"apple":  10,
		"banana": 20,
		"cherry": 30,
	}
	
	// 使用封装的删除函数
	safeDelete(m9, "banana")
	safeDelete(m9, "nonexistent")  // 不会报错
	fmt.Printf("安全删除后: %v\n", m9)
	
	// 删除并返回是否成功
	success := deleteAndCheck(m9, "cherry")
	fmt.Printf("删除 'cherry': %v, map: %v\n", success, m9)
	
	success2 := deleteAndCheck(m9, "nonexistent")
	fmt.Printf("删除 'nonexistent': %v, map: %v\n", success2, m9)
	
	fmt.Println("\n=== 批量删除 ===")
	
	m10 := map[string]int{
		"apple":  10,
		"banana": 20,
		"cherry": 30,
		"date":   40,
		"elder":  50,
	}
	keysToDelete2 := []string{"banana", "date", "nonexistent"}
	
	deletedCount := deleteMultiple(m10, keysToDelete2)
	fmt.Printf("删除了 %d 个元素，剩余: %v\n", deletedCount, m10)
	
	fmt.Println("\n=== 注意事项 ===")
	fmt.Println("1. delete() 函数删除不存在的键不会报错，是安全的")
	fmt.Println("2. 删除元素后，map 的长度会减少")
	fmt.Println("3. 清空 map 推荐使用 make() 重新分配")
	fmt.Println("4. 在循环中删除元素是安全的（Go 1.21+）")
	fmt.Println("5. 删除操作是 O(1) 时间复杂度")
}

// 安全删除函数（检查存在性）
func safeDelete(m map[string]int, key string) {
	if _, exists := m[key]; exists {
		delete(m, key)
	}
}

// 删除并返回是否成功
func deleteAndCheck(m map[string]int, key string) bool {
	if _, exists := m[key]; exists {
		delete(m, key)
		return true
	}
	return false
}

// 批量删除，返回删除的元素数量
func deleteMultiple(m map[string]int, keys []string) int {
	count := 0
	for _, key := range keys {
		if _, exists := m[key]; exists {
			delete(m, key)
			count++
		}
	}
	return count
}
