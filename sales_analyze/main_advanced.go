package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/gocarina/gocsv"
	"github.com/olekukonko/tablewriter"
)

// SalesRecord 销售记录结构体，使用CSV标签映射
type SalesRecord struct {
	Date     string  `csv:"日期"`
	Product  string  `csv:"产品"`
	Quantity int     `csv:"销量"`
	Amount   float64 `csv:"销售额"`
	Region   string  `csv:"地区"`
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

func main() {
	// 设置颜色输出
	color.Set(color.FgCyan, color.Bold)
	fmt.Println("📊 高级销售数据分析系统")
	color.Unset()
	fmt.Println(strings.Repeat("=", 50))

	// 读取CSV文件
	records, err := loadSalesData("sales_data.csv")
	if err != nil {
		color.Red("❌ 读取数据失败: %v", err)
		return
	}

	color.Green("✅ 成功读取 %d 条销售记录", len(records))
	fmt.Println()

	// 执行各种分析
	analyzeOverall(records)
	fmt.Println()
	analyzeByProduct(records)
	fmt.Println()
	analyzeByRegion(records)
	fmt.Println()
	analyzeByDate(records)
}

// loadSalesData 使用gocsv库加载销售数据
func loadSalesData(filename string) ([]SalesRecord, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close()

	var records []SalesRecord
	if err := gocsv.UnmarshalFile(file, &records); err != nil {
		return nil, fmt.Errorf("解析CSV文件失败: %w", err)
	}

	return records, nil
}

// analyzeOverall 总体分析
func analyzeOverall(records []SalesRecord) {
	color.Set(color.FgYellow, color.Bold)
	fmt.Println("📈 总体销售分析")
	color.Unset()

	var totalAmount float64
	var totalQuantity int

	for _, record := range records {
		totalAmount += record.Amount
		totalQuantity += record.Quantity
	}

	avgAmount := totalAmount / float64(len(records))

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"指标", "数值"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	table.Append([]string{"总销售额", fmt.Sprintf("¥ %.2f", totalAmount)})
	table.Append([]string{"总销量", fmt.Sprintf("%d 件", totalQuantity)})
	table.Append([]string{"平均订单金额", fmt.Sprintf("¥ %.2f", avgAmount)})
	table.Append([]string{"订单数量", fmt.Sprintf("%d 笔", len(records))})

	table.Render()
}

// analyzeByProduct 按产品分析
func analyzeByProduct(records []SalesRecord) {
	color.Set(color.FgMagenta, color.Bold)
	fmt.Println("🛍️  产品销售分析")
	color.Unset()

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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"产品", "销量", "销售额", "平均订单", "订单数"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	for _, product := range products {
		table.Append([]string{
			product.Product,
			fmt.Sprintf("%d", product.TotalQty),
			fmt.Sprintf("¥ %.2f", product.TotalAmount),
			fmt.Sprintf("¥ %.2f", product.AvgAmount),
			fmt.Sprintf("%d", product.RecordCount),
		})
	}

	table.Render()
}

// analyzeByRegion 按地区分析
func analyzeByRegion(records []SalesRecord) {
	color.Set(color.FgBlue, color.Bold)
	fmt.Println("🗺️  地区销售分析")
	color.Unset()

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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"地区", "销量", "销售额", "订单数", "市场占比"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	// 计算总销售额用于计算占比
	var totalAmount float64
	for _, region := range regions {
		totalAmount += region.TotalAmount
	}

	for _, region := range regions {
		marketShare := (region.TotalAmount / totalAmount) * 100
		table.Append([]string{
			region.Region,
			fmt.Sprintf("%d", region.TotalQty),
			fmt.Sprintf("¥ %.2f", region.TotalAmount),
			fmt.Sprintf("%d", region.RecordCount),
			fmt.Sprintf("%.1f%%", marketShare),
		})
	}

	table.Render()
}

// analyzeByDate 按日期分析
func analyzeByDate(records []SalesRecord) {
	color.Set(color.FgGreen, color.Bold)
	fmt.Println("📅 日期销售分析")
	color.Unset()

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

	sort.Slice(dates, func(i, j int) bool {
		// 简单的字符串日期比较，实际项目中建议解析为time.Time
		return dates[i] < dates[j]
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"日期", "销量", "销售额", "日增长率"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

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
				growthRate = color.GreenString("+%.1f%%", rate)
			} else {
				growthRate = color.RedString("%.1f%%", rate)
			}
			prevAmount = amount
		}

		table.Append([]string{
			date,
			fmt.Sprintf("%d", qty),
			fmt.Sprintf("¥ %.2f", amount),
			growthRate,
		})
	}

	table.Render()

	// 显示趋势分析
	fmt.Println()
	color.Set(color.FgCyan)
	fmt.Println("📊 趋势分析:")
	color.Unset()
	
	if len(dates) >= 2 {
		firstDay := dateMap[dates[0]]
		lastDay := dateMap[dates[len(dates)-1]]
		totalGrowth := ((lastDay - firstDay) / firstDay) * 100
		
		if totalGrowth > 0 {
			color.Green("📈 整体增长: +%.1f%%", totalGrowth)
		} else {
			color.Red("📉 整体下降: %.1f%%", totalGrowth)
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
	
	color.Green("🏆 最佳销售日: %s (¥ %.2f)", bestDay, maxAmount)
	color.Red("📉 最低销售日: %s (¥ %.2f)", worstDay, minAmount)
}