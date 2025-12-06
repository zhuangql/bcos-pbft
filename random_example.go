package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	mathrand "math/rand"
	"time"
)

func main() {
	fmt.Println("=== math/rand 包：伪随机数生成器 ===")
	
	// 1. 基本用法：需要先设置种子
	mathrand.Seed(time.Now().UnixNano())
	
	// 生成随机整数（0 到 MaxInt32）
	randomInt := mathrand.Int()
	fmt.Printf("随机整数: %d\n", randomInt)
	
	// 生成 0 到 n-1 之间的随机整数
	randomIntN := mathrand.Intn(100)  // 0-99
	fmt.Printf("0-99 随机数: %d\n", randomIntN)
	
	// 生成随机浮点数 [0.0, 1.0)
	randomFloat := mathrand.Float64()
	fmt.Printf("随机浮点数 [0.0, 1.0): %f\n", randomFloat)
	
	// 生成随机浮点数 [0.0, n)
	randomFloatN := mathrand.Float64() * 100
	fmt.Printf("0-100 随机浮点数: %f\n", randomFloatN)
	
	// 生成指定范围的随机整数 [min, max]
	min, max := 10, 50
	randomRange := mathrand.Intn(max-min+1) + min
	fmt.Printf("10-50 随机数: %d\n", randomRange)
	
	// 生成随机布尔值
	randomBool := mathrand.Intn(2) == 0
	fmt.Printf("随机布尔值: %v\n", randomBool)
	
	fmt.Println("\n=== Go 1.20+ 推荐方式：使用 rand.New ===")
	
	// Go 1.20+ 推荐：使用 rand.New 创建独立的随机数生成器
	source := mathrand.NewSource(time.Now().UnixNano())
	rng := mathrand.New(source)
	
	// 使用独立的生成器
	fmt.Printf("独立生成器 - 整数: %d\n", rng.Int())
	fmt.Printf("独立生成器 - 0-99: %d\n", rng.Intn(100))
	fmt.Printf("独立生成器 - 浮点数: %f\n", rng.Float64())
	
	fmt.Println("\n=== crypto/rand 包：加密安全的随机数 ===")
	
	// 生成加密安全的随机整数（0 到 max-1）
	maxBig := big.NewInt(100)
	secureRandom, _ := rand.Int(rand.Reader, maxBig)
	fmt.Printf("安全随机数 0-99: %d\n", secureRandom)
	
	// 生成指定范围的加密安全随机数 [min, max]
	minBig := big.NewInt(10)
	maxBig2 := big.NewInt(50)
	diff := new(big.Int).Sub(maxBig2, minBig)
	diff.Add(diff, big.NewInt(1))  // max - min + 1
	secureRandomRange, _ := rand.Int(rand.Reader, diff)
	secureRandomRange.Add(secureRandomRange, minBig)
	fmt.Printf("安全随机数 10-50: %d\n", secureRandomRange)
	
	// 生成随机字节数组
	randomBytes := make([]byte, 10)
	rand.Read(randomBytes)
	fmt.Printf("随机字节: %v\n", randomBytes)
	fmt.Printf("随机字节(hex): %x\n", randomBytes)
	
	fmt.Println("\n=== 常见应用场景 ===")
	
	// 场景 1: 生成随机字符串
	randomString := generateRandomString(10)
	fmt.Printf("随机字符串(10位): %s\n", randomString)
	
	// 场景 2: 从切片中随机选择元素
	items := []string{"apple", "banana", "cherry", "date", "elderberry"}
	randomItem := items[mathrand.Intn(len(items))]
	fmt.Printf("随机选择: %s\n", randomItem)
	
	// 场景 3: 打乱切片顺序
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	shuffled := make([]int, len(numbers))
	copy(shuffled, numbers)
	mathrand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	fmt.Printf("原数组: %v\n", numbers)
	fmt.Printf("打乱后: %v\n", shuffled)
	
	// 场景 4: 生成随机颜色（RGB）
	red := mathrand.Intn(256)
	green := mathrand.Intn(256)
	blue := mathrand.Intn(256)
	fmt.Printf("随机颜色 RGB(%d, %d, %d)\n", red, green, blue)
	
	// 场景 5: 生成随机 UUID（简化版）
	uuid := generateUUID()
	fmt.Printf("随机 UUID: %s\n", uuid)
	
	// 场景 6: 生成随机密码
	password := generatePassword(12)
	fmt.Printf("随机密码(12位): %s\n", password)
	
	fmt.Println("\n=== 性能对比 ===")
	
	// math/rand 性能测试
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		mathrand.Intn(100)
	}
	mathrandTime := time.Since(start)
	fmt.Printf("math/rand 100万次: %v\n", mathrandTime)
	
	// crypto/rand 性能测试（较慢）
	start = time.Now()
	for i := 0; i < 10000; i++ {
		maxBig := big.NewInt(100)
		rand.Int(rand.Reader, maxBig)
	}
	cryptorandTime := time.Since(start)
	fmt.Printf("crypto/rand 1万次: %v\n", cryptorandTime)
	fmt.Println("注意: crypto/rand 更安全但更慢，适合安全场景")
}

// 生成随机字符串
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[mathrand.Intn(len(charset))]
	}
	return string(b)
}

// 生成随机 UUID（简化版）
func generateUUID() string {
	mathrand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		mathrand.Int31(),
		mathrand.Int31()>>16,
		(mathrand.Int31()>>16)|0x4000,
		(mathrand.Int31()>>16)|0x8000,
		mathrand.Int63(),
	)
}

// 生成随机密码
func generatePassword(length int) string {
	const (
		lowercase = "abcdefghijklmnopqrstuvwxyz"
		uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		digits    = "0123456789"
		special   = "!@#$%^&*"
	)
	allChars := lowercase + uppercase + digits + special
	
	b := make([]byte, length)
	for i := range b {
		b[i] = allChars[mathrand.Intn(len(allChars))]
	}
	return string(b)
}
