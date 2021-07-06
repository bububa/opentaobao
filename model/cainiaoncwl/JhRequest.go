package cainiaoncwl

// JhRequest 结构体
type JhRequest struct {
	// 集单完成时间，查询起点
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 通常指定已集货状态；指定状态查询，不指定返回所有；CONSOLIDATED 已集货;CONSIGNED 已发货;INBOUND 县仓入库;LACK 缺货LACK
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 淘宝item_id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 可以指定集货号查询，此时其他条件忽略；最大指定5条，逗号分隔；
	OrderCodeList string `json:"order_code_list,omitempty" xml:"order_code_list,omitempty"`
	// 集单完成时间，查询尾点；尾点有起点不能超过7天；起点不能是一年以前；
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 淘宝skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 1
	AreaInfo *JhAreaInfo `json:"area_info,omitempty" xml:"area_info,omitempty"`
	// 查询第几页，起点1
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 一页查询多少数据，最大100
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
