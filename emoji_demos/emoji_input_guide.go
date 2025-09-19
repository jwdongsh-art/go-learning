//go:build emoji_input_guide
// +build emoji_input_guide

// 🔤 程序员快速输入Emoji指南
// 本文件演示各种在编程时快速输入emoji的方法

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("⌨️  程序员Emoji输入指南")
	fmt.Println(strings.Repeat("=", 50))

	// 方法概览
	fmt.Println("\n📋 输入方法概览:")
	fmt.Println("1️⃣  系统快捷键 (最常用)")
	fmt.Println("2️⃣  VS Code扩展 (编程专用)")
	fmt.Println("3️⃣  代码片段 (Snippets)")
	fmt.Println("4️⃣  在线工具复制")
	fmt.Println("5️⃣  命令行工具")
	fmt.Println("6️⃣  自定义快捷方式")

	showSystemShortcuts()
	showVSCodeMethods()
	showSnippetsExamples()
	showOnlineTools()
	showCommandLineTools()
	showCustomSolutions()
	showBestPractices()
}

// 1. 系统级快捷键
func showSystemShortcuts() {
	fmt.Println("\n1️⃣  系统快捷键方法:")
	fmt.Println("┌─────────────┬─────────────────────────────────┐")
	fmt.Println("│    系统     │           快捷键                │")
	fmt.Println("├─────────────┼─────────────────────────────────┤")
	fmt.Println("│   macOS     │   Control + Command + Space     │")
	fmt.Println("│   Windows   │   Windows + . 或 Windows + ;    │")
	fmt.Println("│   Linux     │   Ctrl + ; (需要配置)           │")
	fmt.Println("└─────────────┴─────────────────────────────────┘")

	fmt.Println("\n💡 macOS详细使用:")
	fmt.Println("✅  Control + Command + Space 打开emoji面板")
	fmt.Println("✅ 输入关键词搜索: 'fire' → 🔥")
	fmt.Println("✅ 最近使用的emoji会显示在顶部")
	fmt.Println("✅ 支持肤色选择和变体")
}

// 2. VS Code专用方法
func showVSCodeMethods() {
	fmt.Println("\n2️⃣  VS Code专用方法:")

	fmt.Println("\n🔌 推荐扩展:")
	fmt.Println("• :emojisense: - emoji智能提示")
	fmt.Println("• Emoji - emoji面板和搜索")
	fmt.Println("• GitMoji - Git提交专用emoji")
	fmt.Println("• Unicode code point of current character")

	fmt.Println("\n⚡ 使用方法:")
	fmt.Println("1. 安装 :emojisense: 扩展")
	fmt.Println("2. 在代码中输入 :fire: 自动转换为 🔥")
	fmt.Println("3. 或者 Cmd+Shift+P → 'Insert Emoji'")

	fmt.Println("\n📝 常用代码标记:")
	fmt.Println(":fire:        🔥  - 性能优化/热点代码")
	fmt.Println(":rocket:      🚀  - 新功能/发布")
	fmt.Println(":bug:         🐛  - 修复bug")
	fmt.Println(":sparkles:    ✨  - 新特性")
	fmt.Println(":recycle:     ♻️  - 重构代码")
	fmt.Println(":zap:         ⚡  - 性能提升")
	fmt.Println(":memo:        📝  - 文档更新")
	fmt.Println(":lock:        🔒  - 安全相关")
	fmt.Println(":wrench:      🔧  - 配置修改")
	fmt.Println(":package:     📦  - 依赖更新")
}

// 3. 代码片段示例
func showSnippetsExamples() {
	fmt.Println("\n3️⃣  VS Code Snippets配置:")

	fmt.Println(`
// 在 User Snippets 中添加 (Cmd+Shift+P → "Configure User Snippets")
{
  "Success Log": {
    "prefix": "logok",
    "body": "fmt.Printf(\"✅ $1\\n\")",
    "description": "Success log with checkmark"
  },
  "Error Log": {
    "prefix": "logerr", 
    "body": "fmt.Printf(\"❌ $1\\n\")",
    "description": "Error log with cross"
  },
  "Warning Log": {
    "prefix": "logwarn",
    "body": "fmt.Printf(\"⚠️ $1\\n\")", 
    "description": "Warning log"
  },
  "Info Log": {
    "prefix": "loginfo",
    "body": "fmt.Printf(\"ℹ️ $1\\n\")",
    "description": "Info log"
  },
  "Progress": {
    "prefix": "progress",
    "body": "fmt.Printf(\"🚀 Processing: %d%%\\n\", $1)",
    "description": "Progress indicator"
  }
}`)

	fmt.Println("\n使用方法:")
	fmt.Println("• 输入 'logok' + Tab → fmt.Printf(\"✅ %s\\n\")")
	fmt.Println("• 输入 'logerr' + Tab → fmt.Printf(\"❌ %s\\n\")")
}

// 4. 在线工具
func showOnlineTools() {
	fmt.Println("\n4️⃣  在线工具推荐:")

	fmt.Println("\n🌐 常用网站:")
	fmt.Println("• Emojipedia.org - 最全的emoji百科")
	fmt.Println("• Unicode-table.com - Unicode字符表")
	fmt.Println("• Getemoji.com - 简单的emoji复制")
	fmt.Println("• Emojicopy.com - 分类emoji复制")
	fmt.Println("• Gitmoji.dev - Git提交专用emoji")

	fmt.Println("\n🔍 搜索技巧:")
	fmt.Println("• 搜索 'programming emoji' 找编程相关")
	fmt.Println("• 搜索 'status emoji' 找状态指示")
	fmt.Println("• 搜索 'arrow emoji' 找箭头符号")
}

// 5. 命令行工具
func showCommandLineTools() {
	fmt.Println("\n5️⃣  命令行工具:")

	fmt.Println("\n📦 macOS工具:")
	fmt.Println("# 安装 emoji-cli")
	fmt.Println("brew install emoji-cli")
	fmt.Println("# 使用方法")
	fmt.Println("emoji fire    # 输出: 🔥")
	fmt.Println("emoji rocket  # 输出: 🚀")

	fmt.Println("\n🔧 自定义别名 (添加到 ~/.zshrc):")
	fmt.Println(`alias fire="echo -n 🔥"`)
	fmt.Println(`alias rocket="echo -n 🚀"`)
	fmt.Println(`alias check="echo -n ✅"`)
	fmt.Println(`alias cross="echo -n ❌"`)

	fmt.Println("\n使用方法:")
	fmt.Println("fire | pbcopy   # 复制到剪贴板")
	fmt.Println("echo \"$(check) 完成\" # 在命令中使用")
}

// 6. 自定义解决方案
func showCustomSolutions() {
	fmt.Println("\n6️⃣  自定义解决方案:")

	fmt.Println("\n💻 Alfred Workflow (macOS):")
	fmt.Println("• 创建关键词触发器")
	fmt.Println("• emoji fire → 🔥")
	fmt.Println("• emoji code → 💻")

	fmt.Println("\n⌨️  文本替换 (macOS系统设置):")
	fmt.Println("系统设置 → 键盘 → 文本替换")
	fmt.Println("• ::fire:: → 🔥")
	fmt.Println("• ::check:: → ✅")
	fmt.Println("• ::cross:: → ❌")

	fmt.Println("\n🔧 Espanso (跨平台文本扩展):")
	fmt.Println("# 安装")
	fmt.Println("brew install espanso")
	fmt.Println("# 配置文件 ~/.config/espanso/match/emoji.yml")
	fmt.Println(`matches:
  - trigger: ":fire:"
    replace: "🔥"
  - trigger: ":rocket:"
    replace: "🚀"`)
}

// 7. 最佳实践
func showBestPractices() {
	fmt.Println("\n7️⃣  编程中的最佳实践:")

	fmt.Println("\n✅ 推荐使用场景:")
	fmt.Println("• 日志输出: 🔴 ERROR, 🟡 WARN, 🟢 INFO")
	fmt.Println("• 代码注释: 🔧 修复, 🚀 新功能, ⚠️ 注意")
	fmt.Println("• Git提交: 📝 docs, 🐛 fix, ✨ feat")
	fmt.Println("• 进度指示: ⏳ 等待, 🔄 处理中, ✅ 完成")
	fmt.Println("• 状态标识: 🟢 在线, 🔴 离线, 🟡 维护中")

	fmt.Println("\n❌ 避免的场景:")
	fmt.Println("• 变量名、函数名")
	fmt.Println("• API响应数据")
	fmt.Println("• 数据库字段")
	fmt.Println("• 配置文件key")

	fmt.Println("\n🎯 实用技巧:")
	fmt.Println("• 建立个人emoji词汇表")
	fmt.Println("• 团队统一emoji使用规范")
	fmt.Println("• 在README中解释emoji含义")
	fmt.Println("• 适度使用，保持代码专业性")

	fmt.Println("\n📚 个人emoji词汇表建议:")
	emojiDict := map[string]string{
		"🔥":  "性能优化/热点",
		"🚀":  "新功能/发布",
		"🐛":  "Bug修复",
		"✨":  "新特性",
		"♻️": "重构",
		"⚡":  "性能提升",
		"📝":  "文档",
		"🔒":  "安全",
		"🔧":  "配置/工具",
		"📦":  "依赖/打包",
		"🎨":  "UI/样式",
		"💡":  "想法/改进",
		"🧪":  "实验性",
		"🚨":  "关键/紧急",
		"📊":  "数据/分析",
	}

	for emoji, desc := range emojiDict {
		fmt.Printf("%-3s : %s\n", emoji, desc)
	}
}

/*
⚡ 快速输入总结:

1. 🍎 macOS用户 (推荐):
   - Control + Command + Space (系统emoji面板)
   - VS Code + :emojisense: 扩展
   - 文本替换 (::fire:: → 🔥)

2. 🪟 Windows用户:
   - Windows + . (emoji面板)
   - VS Code扩展
   - PowerToys 文本扩展

3. 🐧 Linux用户:
   - 配置emoji输入法
   - VS Code扩展
   - 命令行工具

4. 💼 团队开发:
   - 统一emoji规范
   - 共享snippets配置
   - 文档说明emoji含义

记住: 适度使用emoji可以提高代码可读性，但不要过度使用! 🎯
*/
