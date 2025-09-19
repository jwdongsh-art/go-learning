#!/bin/bash

# 🎨 Emoji & Unicode 演示程序运行脚本

echo "🎭 Emoji & Unicode 演示项目"
echo "=============================="
echo ""
echo "请选择要运行的演示程序:"
echo ""


echo "1️⃣  图标和颜色演示 (icon_demo.go)"
echo "2️⃣  Unicode字符输出演示 (unicode_demo.go)"  
echo "3️⃣  Unicode工具助手 (unicode_helper.go)"
echo "4️⃣  VS Code Unicode显示原理 (vscode_unicode_demo.go)"
echo "5️⃣  Emoji输入指南 (emoji_input_guide.go)"
echo "6️⃣  运行所有演示"
echo "0️⃣  退出"
echo ""

read -p "请输入选择 (0-6): " choice

case $choice in
    1)
        echo "🎨 运行图标和颜色演示..."
        go run -tags icon_demo icon_demo.go
        ;;
    2)
        echo "🔤 运行Unicode字符输出演示..."
        go run -tags unicode_demo unicode_demo.go
        ;;
    3)
        echo "🔧 运行Unicode工具助手..."
        go run -tags unicode_helper unicode_helper.go
        ;;
    4)
        echo "🖥️ 运行VS Code Unicode显示原理..."
        go run -tags vscode_unicode_demo vscode_unicode_demo.go
        ;;
    5)
        echo "⌨️ 运行Emoji输入指南..."
        go run -tags emoji_input_guide emoji_input_guide.go
        ;;
    6)
        echo "🚀 运行所有演示程序..."
        echo ""
        echo "1️⃣ 图标和颜色演示:"
        echo "===================="
        go run -tags icon_demo icon_demo.go
        echo ""
        echo "2️⃣ Unicode字符演示:"
        echo "===================="
        go run -tags unicode_demo unicode_demo.go
        echo ""
        echo "3️⃣ Unicode工具助手:"
        echo "===================="
        go run -tags unicode_helper unicode_helper.go
        echo ""
        echo "4️⃣ VS Code原理说明:"
        echo "===================="
        go run -tags vscode_unicode_demo vscode_unicode_demo.go
        echo ""
        echo "5️⃣ Emoji输入指南:"
        echo "==================="
        go run -tags emoji_input_guide emoji_input_guide.go
        echo ""
        echo "🎉 所有演示完成!"
        ;;
    0)
        echo "👋 再见!"
        exit 0
        ;;
    *)
        echo "❌ 无效选择，请输入 0-6"
        ;;
esac
