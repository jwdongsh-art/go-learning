//go:build !demo
// +build !demo

package main

import "fmt"

// 默认主程序 - 用于VS Code IntelliSense支持
// 实际的演示程序请通过 ./run_demos.sh 运行
func main() {
	fmt.Println("🎭 Emoji & Unicode 演示项目")
	fmt.Println("==============================")
	fmt.Println("")
	fmt.Println("请使用以下方式运行演示:")
	fmt.Println("1. 交互式菜单: ./run_demos.sh")
	fmt.Println("2. 直接运行: go run -tags demo_name demo_file.go")
	fmt.Println("")
	fmt.Println("可用的演示:")
	fmt.Println("- go run -tags icon_demo icon_demo.go")
	fmt.Println("- go run -tags unicode_demo unicode_demo.go")
	fmt.Println("- go run -tags unicode_helper unicode_helper.go")
	fmt.Println("- go run -tags vscode_unicode_demo vscode_unicode_demo.go")
	fmt.Println("- go run -tags emoji_input_guide emoji_input_guide.go")
}
