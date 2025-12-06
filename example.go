package main

import "fmt"

// 示例：在Go结构体中声明数组和map

// 1. 声明固定长度的数组
type Person struct {
	Name    string
	Age     int
	Scores  [5]int        // 固定长度为5的整数数组
	Tags    [3]string     // 固定长度为3的字符串数组
}

// 2. 声明切片（slice）- 动态数组
type Student struct {
	Name     string
	Grades   []int        // 整数切片
	Courses  []string     // 字符串切片
}

// 3. 声明map
type Employee struct {
	Name     string
	Salary   int
	Skills   map[string]int        // map[string]int - 键是字符串，值是整数
	Metadata map[string]string     // map[string]string - 键值都是字符串
}

// 4. 组合使用：结构体中同时包含数组、切片和map
type Company struct {
	Name      string
	Employees []Employee           // 切片：Employee结构体的切片
	Departments map[string][]string // map：键是部门名，值是员工名字符串切片
	Locations  [10]string          // 数组：最多10个位置
}

func main() {
	// 示例1：使用固定长度数组
	person := Person{
		Name:   "张三",
		Age:    25,
		Scores: [5]int{90, 85, 92, 88, 95},
		Tags:   [3]string{"开发", "Go", "后端"},
	}
	fmt.Println("Person:", person)

	// 示例2：使用切片
	student := Student{
		Name:    "李四",
		Grades:  []int{95, 87, 92},
		Courses: []string{"数学", "英语", "编程"},
	}
	fmt.Println("Student:", student)

	// 示例3：使用map（需要先初始化）
	employee := Employee{
		Name:   "王五",
		Salary: 10000,
		Skills: map[string]int{
			"Go":     8,
			"Python": 7,
			"Java":   6,
		},
		Metadata: map[string]string{
			"部门": "技术部",
			"级别": "高级",
		},
	}
	fmt.Println("Employee:", employee)

	// 示例4：组合使用
	company := Company{
		Name: "示例公司",
		Employees: []Employee{
			{Name: "员工1", Salary: 8000},
			{Name: "员工2", Salary: 9000},
		},
		Departments: map[string][]string{
			"技术部": {"员工1", "员工2"},
			"销售部": {"员工3"},
		},
		Locations: [10]string{"北京", "上海", "深圳"},
	}
	fmt.Println("Company:", company)

	// 重要提示：
	// 1. 数组：固定长度，声明时指定大小，如 [5]int
	// 2. 切片：动态长度，使用 []int，更常用
	// 3. map：需要初始化后才能使用，否则会panic
	//    可以使用 make() 初始化：make(map[string]int)
	//    或者直接使用字面量初始化
}
