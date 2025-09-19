package main

import (
	"fmt"
	"strconv"
	"strings"
)

// UnicodeHelper Unicode字符助手
type UnicodeHelper struct{}

// FromCodePoint 从Unicode码点创建字符
func (h *UnicodeHelper) FromCodePoint(codePoint string) (string, error) {
	// 处理 "U+1F600" 格式
	if strings.HasPrefix(codePoint, "U+") {
		codePoint = codePoint[2:]
	}

	// 解析十六进制数
	value, err := strconv.ParseInt(codePoint, 16, 32)
	if err != nil {
		return "", fmt.Errorf("无效的Unicode码点: %s", codePoint)
	}

	return string(rune(value)), nil
}

// ToCodePoint 将字符转换为Unicode码点
func (h *UnicodeHelper) ToCodePoint(char string) []string {
	var codePoints []string
	for _, r := range char {
		codePoints = append(codePoints, fmt.Sprintf("U+%04X", r))
	}
	return codePoints
}

// CommonEmojis 常用emoji映射
var CommonEmojis = map[string]string{
	"grinning":    "U+1F600", // 😀
	"joy":         "U+1F602", // 😂
	"heart_eyes":  "U+1F60D", // 😍
	"thinking":    "U+1F914", // 🤔
	"thumbs_up":   "U+1F44D", // 👍
	"thumbs_down": "U+1F44E", // 👎
	"fire":        "U+1F525", // 🔥
	"rocket":      "U+1F680", // 🚀
	"star":        "U+2B50",  // ⭐
	"check":       "U+2705",  // ✅
	"cross":       "U+274C",  // ❌
	"warning":     "U+26A0",  // ⚠️
	"info":        "U+2139",  // ℹ️
	"bulb":        "U+1F4A1", // 💡
	"computer":    "U+1F4BB", // 💻
	"phone":       "U+1F4F1", // 📱
	"chart_up":    "U+1F4C8", // 📈
	"chart_down":  "U+1F4C9", // 📉
	"bar_chart":   "U+1F4CA", // 📊
}

func main() {
	helper := &UnicodeHelper{}

	fmt.Println("🔧 Unicode工具演示")
	fmt.Println(strings.Repeat("=", 40))

	// 示例1: 从码点创建字符
	fmt.Println("\n📝 从Unicode码点创建字符:")
	testCodes := []string{"U+1F600", "1F680", "U+2705", "26A0"}

	for _, code := range testCodes {
		char, err := helper.FromCodePoint(code)
		if err != nil {
			fmt.Printf("❌ %s: %v\n", code, err)
		} else {
			fmt.Printf("✅ %-8s -> %s\n", code, char)
		}
	}

	// 示例2: 从字符获取码点
	fmt.Println("\n🔍 从字符获取Unicode码点:")
	testChars := []string{"😀", "🚀", "✅", "中文"}

	for _, char := range testChars {
		codes := helper.ToCodePoint(char)
		fmt.Printf("%-4s -> %s\n", char, strings.Join(codes, " "))
	}

	// 示例3: 使用预定义emoji
	fmt.Println("\n🎭 使用预定义emoji:")
	for name, code := range CommonEmojis {
		char, _ := helper.FromCodePoint(code)
		fmt.Printf("%-12s: %-8s -> %s\n", name, code, char)
	}

	// 示例4: 在实际项目中的使用
	fmt.Println("\n💼 项目中的实际使用示例:")

	// 状态指示器
	statusEmojis := map[string]string{
		"success": CommonEmojis["check"],
		"error":   CommonEmojis["cross"],
		"warning": CommonEmojis["warning"],
		"info":    CommonEmojis["info"],
	}

	for status, code := range statusEmojis {
		emoji, _ := helper.FromCodePoint(code)
		fmt.Printf("%s 状态: %s %s 操作完成\n",
			strings.Title(status), emoji, status)
	}

	// 进度指示器
	fmt.Println("\n📊 进度指示器:")
	rocket, _ := helper.FromCodePoint(CommonEmojis["rocket"])
	chart, _ := helper.FromCodePoint(CommonEmojis["chart_up"])
	fire, _ := helper.FromCodePoint(CommonEmojis["fire"])

	for i := 0; i <= 100; i += 25 {
		var icon string
		switch {
		case i == 0:
			icon = "⏳"
		case i < 50:
			icon = rocket
		case i < 100:
			icon = chart
		default:
			icon = fire
		}
		fmt.Printf("%s 进度: %d%%\n", icon, i)
	}

	// 示例5: 批量创建emoji
	fmt.Println("\n🌈 批量创建表情:")
	// 笑脸范围: U+1F600 - U+1F64F
	fmt.Print("笑脸系列: ")
	for i := 0x1F600; i <= 0x1F60F; i++ {
		fmt.Printf("%c", rune(i))
	}
	fmt.Println()

	// 手势范围: U+1F44D - U+1F44F
	fmt.Print("手势系列: ")
	for i := 0x1F44D; i <= 0x1F450; i++ {
		fmt.Printf("%c", rune(i))
	}
	fmt.Println()
}
