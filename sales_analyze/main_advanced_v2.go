package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// SalesRecord 销售记录结构体
type SalesRecord struct {
	Date     string
	Product  string
	Quantity int
	Amount   float64
	Region   string
}

// ProductSummary 产品汇总结构体
type ProductSummary struct {
	Product      string
	TotalQty     int
	TotalAmount  float64
	AvgAmount    float64
	RecordCount  int
}

// RegionSummary 地区汇总结构体
type RegionSummary struct {
	Region      string
	TotalQty    int
	TotalAmount float64
	RecordCount int
}

// ANSI颜色代码
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
)

func main() {
	// 打印标题
	printHeader("📊 高级销售数据分析系统", ColorCyan)
	
	// 读取CSV文件
	records, err := loadSalesData("sales_data.csv")
	if err != nil {
		printError("❌ 读取数据失败: %v", err)
		return
	}

	printSuccess("✅ 成功读取 %d 条销售记录\n", len(records))

	// 执行各种分析
	analyzeOverall(records)
	fmt.Println()
	analyzeByProduct(records)
	fmt.Println()
	analyzeByRegion(records)
	fmt.Println()
	analyzeByDate(records)
}

// 颜色打印函数
func printHeader(text, color string) {
	fmt.Printf("%s%s%s%s\n", ColorBold, color, text, ColorReset)
	fmt.Println(strings.Repeat("=", 50))
}

func printSuccess(format string, args ...interface{}) {
	fmt.Printf("%s%s%s", ColorGreen, fmt.Sprintf(format, args...), ColorReset)
}

func printError(format string, args ...interface{}) {
	fmt.Printf("%s%s%s", ColorRed, fmt.Sprintf(format, args...), ColorReset)
}

func printInfo(format string, args ...interface{}) {
	fmt.Printf("%s%s%s", ColorBlue, fmt.Sprintf(format, args...), ColorReset)
}

func printWarning(format string, args ...interface{}) {
	fmt.Printf("%s%s%s", ColorYellow, fmt.Sprintf(format, args...), ColorReset)
}

// loadSalesData 加载销售数据
func loadSalesData(filename string) ([]SalesRecord, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("读取CSV文件失败: %w", err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("CSV文件没有数据行")
	}

	var records []SalesRecord
	// 跳过标题行
	for i, row := range rows[1:] {
		if len(row) != 5 {
			fmt.Printf("⚠️  第%d行数据格式错误，跳过\n", i+2)
			continue
		}

		quantity, err := strconv.Atoi(row[2])
		if err != nil {
			fmt.Printf("⚠️  第%d行销量数据错误: %v，跳过\n", i+2, err)
			continue
		}

		amount, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Printf("⚠️  第%d行销售额数据错误: %v，跳过\n", i+2, err)
			continue
		}

		record := SalesRecord{
			Date:     row[0],
			Product:  row[1],
			Quantity: quantity,
			Amount:   amount,
			Region:   row[4],
		}
		records = append(records, record)
	}

	return records, nil
}

// printTable 打印表格
func printTable(headers []string, rows [][]string) {
	// 计算每列的最大宽度
	colWidths := make([]int, len(headers))
	for i, header := range headers {
		colWidths[i] = len(header)
	}

	for _, row := range rows {
		for i, cell := range row {
			if i < len(colWidths) && len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	// 打印分隔线
	printSeparator(colWidths)

	// 打印表头
	fmt.Print("|")
	for i, header := range headers {
		fmt.Printf(" %-*s |", colWidths[i], header)
	}
	fmt.Println()

	// 打印分隔线
	printSeparator(colWidths)

	// 打印数据行
	for _, row := range rows {
		fmt.Print("|")
		for i, cell := range row {
			if i < len(colWidths) {
				fmt.Printf(" %-*s |", colWidths[i], cell)
			}
		}
		fmt.Println()
	}

	// 打印底部分隔线
	printSeparator(colWidths)
}

func printSeparator(colWidths []int) {
	fmt.Print("+")
	for _, width := range colWidths {
		fmt.Print(strings.Repeat("-", width+2) + "+")
	}
	fmt.Println()
}

// analyzeOverall 总体分析
func analyzeOverall(records []SalesRecord) {
	printHeader("📈 总体销售分析", ColorYellow)

	var totalAmount float64
	var totalQuantity int

	for _, record := range records {
		totalAmount += record.Amount
		totalQuantity += record.Quantity
	}

	avgAmount := totalAmount / float64(len(records))

	headers := []string{"指标", "数值"}
	rows := [][]string{
		{"总销售额", fmt.Sprintf("¥ %.2f", totalAmount)},
		{"总销量", fmt.Sprintf("%d 件", totalQuantity)},
		{"平均订单金额", fmt.Sprintf("¥ %.2f", avgAmount)},
		{"订单数量", fmt.Sprintf("%d 笔", len(records))},
	}

	printTable(headers, rows)
}

// analyzeByProduct 按产品分析
func analyzeByProduct(records []SalesRecord) {
	printHeader("🛍️  产品销售分析", ColorPurple)

	productMap := make(map[string]*ProductSummary)

	for _, record := range records {
		if summary, exists := productMap[record.Product]; exists {
			summary.TotalQty += record.Quantity
			summary.TotalAmount += record.Amount
			summary.RecordCount++
		} else {
			productMap[record.Product] = &ProductSummary{
				Product:     record.Product,
				TotalQty:    record.Quantity,
				TotalAmount: record.Amount,
				RecordCount: 1,
			}
		}
	}

	// 计算平均金额
	for _, summary := range productMap {
		summary.AvgAmount = summary.TotalAmount / float64(summary.RecordCount)
	}

	// 转换为切片并按销售额排序
	var products []*ProductSummary
	for _, summary := range productMap {
		products = append(products, summary)
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].TotalAmount > products[j].TotalAmount
	})

	headers := []string{"产品", "销量", "销售额", "平均订单", "订单数"}
	var rows [][]string

	for _, product := range products {
		rows = append(rows, []string{
			product.Product,
			fmt.Sprintf("%d", product.TotalQty),
			fmt.Sprintf("¥ %.2f", product.TotalAmount),
			fmt.Sprintf("¥ %.2f", product.AvgAmount),
			fmt.Sprintf("%d", product.RecordCount),
		})
	}

	printTable(headers, rows)
	
	// 显示最佳产品
	if len(products) > 0 {
		fmt.Println()
		printSuccess("🏆 最佳销售产品: %s (¥ %.2f)\n", products[0].Product, products[0].TotalAmount)
	}
}

// analyzeByRegion 按地区分析
func analyzeByRegion(records []SalesRecord) {
	printHeader("🗺️  地区销售分析", ColorBlue)

	regionMap := make(map[string]*RegionSummary)

	for _, record := range records {
		if summary, exists := regionMap[record.Region]; exists {
			summary.TotalQty += record.Quantity
			summary.TotalAmount += record.Amount
			summary.RecordCount++
		} else {
			regionMap[record.Region] = &RegionSummary{
				Region:      record.Region,
				TotalQty:    record.Quantity,
				TotalAmount: record.Amount,
				RecordCount: 1,
			}
		}
	}

	// 转换为切片并按销售额排序
	var regions []*RegionSummary
	for _, summary := range regionMap {
		regions = append(regions, summary)
	}

	sort.Slice(regions, func(i, j int) bool {
		return regions[i].TotalAmount > regions[j].TotalAmount
	})

	// 计算总销售额用于计算占比
	var totalAmount float64
	for _, region := range regions {
		totalAmount += region.TotalAmount
	}

	headers := []string{"地区", "销量", "销售额", "订单数", "市场占比"}
	var rows [][]string

	for _, region := range regions {
		marketShare := (region.TotalAmount / totalAmount) * 100
		rows = append(rows, []string{
			region.Region,
			fmt.Sprintf("%d", region.TotalQty),
			fmt.Sprintf("¥ %.2f", region.TotalAmount),
			fmt.Sprintf("%d", region.RecordCount),
			fmt.Sprintf("%.1f%%", marketShare),
		})
	}

	printTable(headers, rows)
	
	// 显示最佳地区
	if len(regions) > 0 {
		fmt.Println()
		printSuccess("🏆 最佳销售地区: %s (¥ %.2f)\n", regions[0].Region, regions[0].TotalAmount)
	}
}

// analyzeByDate 按日期分析
func analyzeByDate(records []SalesRecord) {
	printHeader("📅 日期销售分析", ColorGreen)

	dateMap := make(map[string]float64)
	dateQtyMap := make(map[string]int)

	for _, record := range records {
		dateMap[record.Date] += record.Amount
		dateQtyMap[record.Date] += record.Quantity
	}

	// 获取所有日期并排序
	var dates []string
	for date := range dateMap {
		dates = append(dates, date)
	}

	sort.Strings(dates)

	headers := []string{"日期", "销量", "销售额", "日增长率"}
	var rows [][]string

	var prevAmount float64
	for i, date := range dates {
		amount := dateMap[date]
		qty := dateQtyMap[date]
		
		var growthRate string
		if i == 0 {
			growthRate = "-"
			prevAmount = amount
		} else {
			rate := ((amount - prevAmount) / prevAmount) * 100
			if rate > 0 {
				growthRate = fmt.Sprintf("+%.1f%%", rate)
			} else {
				growthRate = fmt.Sprintf("%.1f%%", rate)
			}
			prevAmount = amount
		}

		rows = append(rows, []string{
			date,
			fmt.Sprintf("%d", qty),
			fmt.Sprintf("¥ %.2f", amount),
			growthRate,
		})
	}

	printTable(headers, rows)

	// 显示趋势分析
	fmt.Println()
	printInfo("📊 趋势分析:\n")
	
	if len(dates) >= 2 {
		firstDay := dateMap[dates[0]]
		lastDay := dateMap[dates[len(dates)-1]]
		totalGrowth := ((lastDay - firstDay) / firstDay) * 100
		
		if totalGrowth > 0 {
			printSuccess("📈 整体增长: +%.1f%%\n", totalGrowth)
		} else {
			printWarning("📉 整体变化: %.1f%%\n", totalGrowth)
		}
	}

	// 找出最佳和最差销售日
	var maxAmount, minAmount float64
	var bestDay, worstDay string
	
	for i, date := range dates {
		amount := dateMap[date]
		if i == 0 {
			maxAmount = amount
			minAmount = amount
			bestDay = date
			worstDay = date
		} else {
			if amount > maxAmount {
				maxAmount = amount
				bestDay = date
			}
			if amount < minAmount {
				minAmount = amount
				worstDay = date
			}
		}
	}
	
	printSuccess("🏆 最佳销售日: %s (¥ %.2f)\n", bestDay, maxAmount)
	printWarning("📉 最低销售日: %s (¥ %.2f)\n", worstDay, minAmount)
}
