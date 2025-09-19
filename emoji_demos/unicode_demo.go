//go:build unicode_demo
// +build unicode_demo

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("🔤 Unicode字符输出演示")
	fmt.Println(strings.Repeat("=", 30))

	// 方法1: 直接使用Unicode转义序列
	fmt.Println("\n方法1: Unicode转义序列")
	fmt.Printf("U+1F600: \\u%04X -> 不支持4字节Unicode\n", 0x1F600)
	fmt.Printf("需要使用: \\U%08X\n", 0x1F600)
	fmt.Printf("输出: \\U0001F600 -> \U0001F600\n")

	// 方法2: 使用rune类型
	fmt.Println("\n方法2: 使用rune类型")
	var emoji rune = 0x1F600
	fmt.Printf("rune值: %U -> %c\n", emoji, emoji)

	// 方法3: 使用string函数转换
	fmt.Println("\n方法3: 使用string()转换")
	emojiStr := string(rune(0x1F600))
	fmt.Printf("string(0x1F600): %s\n", emojiStr)

	// 方法4: 使用utf8.RuneCountInString验证
	fmt.Println("\n方法4: UTF-8编码信息")
	fmt.Printf("字符串长度(字节): %d\n", len(emojiStr))
	fmt.Printf("字符串长度(字符): %d\n", utf8.RuneCountInString(emojiStr))

	// 方法5: 常见Unicode字符演示
	fmt.Println("\n方法5: 常见Unicode字符")
	unicodeChars := map[string]rune{
		"U+1F600": 0x1F600, // 😀 grinning face
		"U+1F601": 0x1F601, // 😁 beaming face with smiling eyes
		"U+1F602": 0x1F602, // 😂 face with tears of joy
		"U+1F603": 0x1F603, // 😃 grinning face with big eyes
		"U+1F604": 0x1F604, // 😄 grinning face with smiling eyes
		"U+1F605": 0x1F605, // 😅 grinning face with sweat
		"U+1F606": 0x1F606, // 😆 grinning squinting face
		"U+1F607": 0x1F607, // 😇 smiling face with halo
		"U+1F608": 0x1F608, // 😈 smiling face with horns
		"U+1F609": 0x1F609, // 😉 winking face
	}

	for unicode, char := range unicodeChars {
		fmt.Printf("%-8s -> %c\n", unicode, char)
	}

	// 方法6: 动态构造Unicode字符
	fmt.Println("\n方法6: 动态构造Unicode")
	baseEmoji := 0x1F600
	for i := 0; i < 10; i++ {
		fmt.Printf("U+%04X: %c ", baseEmoji+i, rune(baseEmoji+i))
	}
	fmt.Println()

	// 方法7: 技术相关Unicode字符
	fmt.Println("\n方法7: 技术相关Unicode字符")
	techChars := map[string]rune{
		"U+1F4BB": 0x1F4BB, // 💻 laptop computer
		"U+1F4F1": 0x1F4F1, // 📱 mobile phone
		"U+1F4CA": 0x1F4CA, // 📊 bar chart
		"U+1F4C8": 0x1F4C8, // 📈 chart increasing
		"U+1F4C9": 0x1F4C9, // 📉 chart decreasing
		"U+1F680": 0x1F680, // 🚀 rocket
		"U+26A1":  0x26A1,  // ⚡ high voltage
		"U+1F4A1": 0x1F4A1, // 💡 light bulb
		"U+1F527": 0x1F527, // 🔧 wrench
		"U+1F528": 0x1F528, // 🔨 hammer
	}

	for unicode, char := range techChars {
		fmt.Printf("%-8s -> %c\n", unicode, char)
	}

	// 方法8: 在代码中使用的实际例子
	fmt.Println("\n方法8: 代码中的实际使用")
	fmt.Printf("直接在字符串中: \"开始处理 %c\"\n", rune(0x1F680)) // 🚀
	fmt.Printf("格式化输出: \"进度: %c%c%c\"\n",
		rune(0x1F4CA), rune(0x1F4C8), rune(0x2705)) // 📊📈✅

	// 方法9: 检查字符是否为有效的Unicode
	fmt.Println("\n方法9: Unicode有效性检查")
	testRunes := []rune{0x1F600, 0x41, 0x4E2D, 0x110000} // 超出范围的Unicode
	for _, r := range testRunes {
		if utf8.ValidRune(r) {
			fmt.Printf("U+%04X: %c (有效)\n", r, r)
		} else {
			fmt.Printf("U+%04X: 无效的Unicode字符\n", r)
		}
	}
}
