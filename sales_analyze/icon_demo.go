package main

import (
	"fmt"
	"strings"
)

// ANSI 颜色代码常量
const (
	ColorReset  = "\033[0m"    // 重置颜色
	ColorRed    = "\033[31m"   // 红色
	ColorGreen  = "\033[32m"   // 绿色
	ColorYellow = "\033[33m"   // 黄色
	ColorBlue   = "\033[34m"   // 蓝色
	ColorPurple = "\033[35m"   // 紫色
	ColorCyan   = "\033[36m"   // 青色
	ColorWhite  = "\033[37m"   // 白色
	ColorBold   = "\033[1m"    // 粗体
)

func main() {
	fmt.Println("🎨 图标和颜色演示程序")
	fmt.Println(strings.Repeat("=", 30))
	
	// 1. Unicode Emoji 图标演示
	fmt.Println("\n📱 Unicode Emoji 图标:")
	fmt.Println("📊 图表  📈 上升  📉 下降")
	fmt.Println("🛍️ 购物  🗺️ 地图  📅 日历")
	fmt.Println("🏆 奖杯  ✅ 成功  ❌ 错误  ⚠️ 警告")
	fmt.Println("🎯 目标  🚀 火箭  💡 灯泡  🔥 火焰")
	
	// 2. ANSI 颜色代码演示
	fmt.Println("\n🌈 ANSI 颜色代码演示:")
	fmt.Printf("%s红色文本%s\n", ColorRed, ColorReset)
	fmt.Printf("%s绿色文本%s\n", ColorGreen, ColorReset)
	fmt.Printf("%s黄色文本%s\n", ColorYellow, ColorReset)
	fmt.Printf("%s蓝色文本%s\n", ColorBlue, ColorReset)
	fmt.Printf("%s紫色文本%s\n", ColorPurple, ColorReset)
	fmt.Printf("%s青色文本%s\n", ColorCyan, ColorReset)
	fmt.Printf("%s%s粗体文本%s\n", ColorBold, ColorYellow, ColorReset)
	
	// 3. 组合使用
	fmt.Println("\n🎭 组合使用演示:")
	fmt.Printf("%s%s🎉 彩色图标组合 🎊%s\n", ColorBold, ColorGreen, ColorReset)
	fmt.Printf("%s⚡ %s高性能%s%s ⚡%s\n", ColorYellow, ColorBold, ColorReset, ColorYellow, ColorReset)
	
	// 4. 实际应用场景
	fmt.Println("\n💼 实际应用场景:")
	printStatus("success", "✅ 数据加载成功")
	printStatus("error", "❌ 连接失败")
	printStatus("warning", "⚠️ 内存使用率较高")
	printStatus("info", "📊 正在分析数据...")
	
	// 5. 进度条模拟
	fmt.Println("\n📈 进度条演示:")
	for i := 0; i <= 100; i += 20 {
		fmt.Printf("\r%s处理进度: ", ColorCyan)
		for j := 0; j < i/5; j++ {
			fmt.Print("█")
		}
		for j := i/5; j < 20; j++ {
			fmt.Print("░")
		}
		fmt.Printf(" %d%%%s", i, ColorReset)
		if i == 100 {
			fmt.Printf(" %s🎉 完成!%s\n", ColorGreen, ColorReset)
		}
	}
}

// 状态打印函数
func printStatus(level, message string) {
	switch level {
	case "success":
		fmt.Printf("%s%s%s\n", ColorGreen, message, ColorReset)
	case "error":
		fmt.Printf("%s%s%s\n", ColorRed, message, ColorReset)
	case "warning":
		fmt.Printf("%s%s%s\n", ColorYellow, message, ColorReset)
	case "info":
		fmt.Printf("%s%s%s\n", ColorBlue, message, ColorReset)
	default:
		fmt.Println(message)
	}
}

// 可用的常见 Unicode Emoji (按分类整理):
/*
📊 数据分析类:
� �📈 📉 📋 📌 📍 📎 📏 📐 📑 📒 📓 📔 📕 📖 📗 📘 📙 📚 📛 📜 📝

� 商业办公类:
� � � � � �  📨 📩 � � � �📬 📭 📮 �    💹

🎯 状态指示类:
✅ ❌ ⚠️ ℹ️ � � � � � � ⚪ ⚫ 🔘  🔷 🔸 🔹 🔺 🔻 🔼 🔽

� 技术开发类:
� ⚡ � � � � ⚙️ � � � � � �️ ⌨️ �️ �️ � � �

🏆 成就奖励类:
🏆 🥇 🥈 🥉 🎖️ 🏅 🎗️ 🎯 🎪 🎭 🎨 🎬 🎵 🎶 🎸 🎹 🎺 🎻

🌟 情感表达类:
😀 😃 😄 � 😆 � � 🤣 😊 � � � � � � 🥰 �   😚

� 强调突出类:
� ⭐ 🌟 ✨ � � 🎉 🎊 🎈 🎁 🎀 🏃 � ⚡ 🌈 🦄

�️ 地理位置类:
�️ 🌍 🌎 🌏 🌐 🏔️ ⛰️ 🏕️ 🏖️ 🏜️ 🏝️ 🏞️ 🏟️ 🏛️ 🏗️ 🏘️ 🏚️ 🏠

� 时间日期类:
� � �️ ⏰ ⏲️ ⏱️ ⏳ ⌛ � � � � � � � � � � � �

� 交通工具类:
� � �   🏎️ � 🚑 🚒 � �  🚜 🏍️ � � � � ✈️ �

🍎 食物饮品类:
🍎 🍊 🍋 🍌 🍉 🍇 🍓 � 🍈 🍒 🍑 🥭 🍍 🥥 🥝 🍅 🍆 🥑 🥦 🥬

官方参考资源:
- Unicode官方: https://unicode.org/emoji/
- Emojipedia: https://emojipedia.org/
- 完整列表: https://unicode.org/emoji/charts/full-emoji-list.html
*/
