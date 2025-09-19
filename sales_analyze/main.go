package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// SalesData 表示销售数据结构
type SalesData struct {
	Date     string
	Product  string
	Quantity int
	Amount   float64
	Region   string
}

func main() {
	// 打开CSV文件
	file, err := os.Open("sales_data.csv")
	if err != nil {
		fmt.Printf("无法打开文件: %v\n", err)
		return
	}
	defer file.Close()

	// 创建CSV读取器
	reader := csv.NewReader(file)

	// 读取所有记录
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("读取CSV文件失败: %v\n", err)
		return
	}

	// 跳过标题行，从第二行开始处理数据
	if len(records) < 2 {
		fmt.Println("CSV文件没有数据")
		return
	}

	var salesData []SalesData
	var totalAmount float64
	productQuantity := make(map[string]int)
	productAmount := make(map[string]float64)

	// 处理每一行数据（跳过标题行）
	for i, record := range records[1:] {
		if len(record) != 5 {
			fmt.Printf("第%d行数据格式错误\n", i+2)
			continue
		}

		// 解析销量
		quantity, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Printf("第%d行销量数据错误: %v\n", i+2, err)
			continue
		}

		// 解析销售额
		amount, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			fmt.Printf("第%d行销售额数据错误: %v\n", i+2, err)
			continue
		}

		// 创建销售数据对象
		data := SalesData{
			Date:     record[0],
			Product:  record[1],
			Quantity: quantity,
			Amount:   amount,
			Region:   record[4],
		}
		salesData = append(salesData, data)

		// 累计总销售额
		totalAmount += amount

		// 累计各产品销量和销售额
		productQuantity[data.Product] += quantity
		productAmount[data.Product] += amount
	}

	// 输出统计结果
	fmt.Println("=== 销售数据统计报告 ===")
	fmt.Printf("总销售额: %.2f 元\n", totalAmount)
	fmt.Println()

	fmt.Println("各产品销量统计:")
	fmt.Println("产品名称\t销量\t销售额")
	fmt.Println("------------------------")
	for product, quantity := range productQuantity {
		amount := productAmount[product]
		fmt.Printf("%s\t\t%d\t%.2f\n", product, quantity, amount)
	}

	fmt.Println()
	fmt.Printf("共处理 %d 条销售记录\n", len(salesData))
}
