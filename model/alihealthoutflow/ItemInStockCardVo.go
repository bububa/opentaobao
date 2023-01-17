package alihealthoutflow

// ItemInStockCardVo 结构体
type ItemInStockCardVo struct {
	// 疫苗卡片数据
	ItemInStockList []ItemInStockVo `json:"item_in_stock_list,omitempty" xml:"item_in_stock_list>item_in_stock_vo,omitempty"`
	// 标签名
	TabName string `json:"tab_name,omitempty" xml:"tab_name,omitempty"`
}
