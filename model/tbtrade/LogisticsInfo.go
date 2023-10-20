package tbtrade

import (
	"sync"
)

// LogisticsInfo 结构体
type LogisticsInfo struct {
	// 货品仓储id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 货品仓储code
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 如是菜鸟仓，则将菜鸟仓的区域仓code进行填充，如是商家仓发货则填充SC
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 子订单类型:标示该子交易单来源交易，还是BMS增加的，枚举值(00=交易，10=BMS绑定)
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 商品的最小库存单位Sku的id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 发货类型 CN=菜鸟发货,SC的商家仓发货
	ConsignType string `json:"consign_type,omitempty" xml:"consign_type,omitempty"`
	// 组合货品ID
	CombineItemId string `json:"combine_item_id,omitempty" xml:"combine_item_id,omitempty"`
	// 组合货品Code
	CombineItemCode string `json:"combine_item_code,omitempty" xml:"combine_item_code,omitempty"`
	// 货品BarCode
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 择配信息
	DeliveryCps string `json:"delivery_cps,omitempty" xml:"delivery_cps,omitempty"`
	// 推荐配送erp编码
	BizDeliveryCode string `json:"biz_delivery_code,omitempty" xml:"biz_delivery_code,omitempty"`
	// 仓商家编码
	BizStoreCode string `json:"biz_store_code,omitempty" xml:"biz_store_code,omitempty"`
	// 仓配建议类型
	BizSdType string `json:"biz_sd_type,omitempty" xml:"biz_sd_type,omitempty"`
	// 服务决策的发货地，国家
	SendCountry string `json:"send_country,omitempty" xml:"send_country,omitempty"`
	// 服务决策的发货地，省份
	SendState string `json:"send_state,omitempty" xml:"send_state,omitempty"`
	// 服务决策的发货地，城市
	SendCity string `json:"send_city,omitempty" xml:"send_city,omitempty"`
	// 服务决策的发货地，地区
	SendDistrict string `json:"send_district,omitempty" xml:"send_district,omitempty"`
	// 服务决策的发货地，街道地址
	SendTown string `json:"send_town,omitempty" xml:"send_town,omitempty"`
	// 服务决策的发货地最小地址编码
	SendDivisionCode string `json:"send_division_code,omitempty" xml:"send_division_code,omitempty"`
	// 服务决策的快递黑名单列表
	BlackDeliveryCps string `json:"black_delivery_cps,omitempty" xml:"black_delivery_cps,omitempty"`
	// 服务决策的快递白名单列表
	WhiteDeliveryCps string `json:"white_delivery_cps,omitempty" xml:"white_delivery_cps,omitempty"`
	// 未使用仓建议报错
	UnusedWarehouseErrorMsg string `json:"unused_warehouse_error_msg,omitempty" xml:"unused_warehouse_error_msg,omitempty"`
	// 未使用配建议报错
	UnusedDeliveryErrorMsg string `json:"unused_delivery_error_msg,omitempty" xml:"unused_delivery_error_msg,omitempty"`
	// 使用禁止配报错
	UsedBlackDeliveryErrorMsg string `json:"used_black_delivery_error_msg,omitempty" xml:"used_black_delivery_error_msg,omitempty"`
	// 承诺/最晚出库时间
	PromiseOutboundTime string `json:"promise_outbound_time,omitempty" xml:"promise_outbound_time,omitempty"`
	// 承诺/最晚揽收时间
	PromiseCollectTime string `json:"promise_collect_time,omitempty" xml:"promise_collect_time,omitempty"`
	// 主交易号
	TradeId int64 `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 子交易号
	SubTradeId int64 `json:"sub_trade_id,omitempty" xml:"sub_trade_id,omitempty"`
	// 应发数量
	NeedConsignNum int64 `json:"need_consign_num,omitempty" xml:"need_consign_num,omitempty"`
	// 商品数字编号
	NumIid int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 组合货品比例
	ItemRatio int64 `json:"item_ratio,omitempty" xml:"item_ratio,omitempty"`
	// 订单推荐配送类型      * 0：子单无配建议；ERP按照自己的逻辑进行择配。      * 1：子单有推荐配list，erp可按需参考。      * 2：子单有推荐配list，erp必须在推荐配list中选择配品牌。      * 3：子单有禁用配list，erp需要过滤配品牌。
	BizDeliveryType int64 `json:"biz_delivery_type,omitempty" xml:"biz_delivery_type,omitempty"`
}

var poolLogisticsInfo = sync.Pool{
	New: func() any {
		return new(LogisticsInfo)
	},
}

// GetLogisticsInfo() 从对象池中获取LogisticsInfo
func GetLogisticsInfo() *LogisticsInfo {
	return poolLogisticsInfo.Get().(*LogisticsInfo)
}

// ReleaseLogisticsInfo 释放LogisticsInfo
func ReleaseLogisticsInfo(v *LogisticsInfo) {
	v.ItemId = ""
	v.ItemCode = ""
	v.StoreCode = ""
	v.Type = ""
	v.SkuId = ""
	v.ConsignType = ""
	v.CombineItemId = ""
	v.CombineItemCode = ""
	v.BarCode = ""
	v.DeliveryCps = ""
	v.BizDeliveryCode = ""
	v.BizStoreCode = ""
	v.BizSdType = ""
	v.SendCountry = ""
	v.SendState = ""
	v.SendCity = ""
	v.SendDistrict = ""
	v.SendTown = ""
	v.SendDivisionCode = ""
	v.BlackDeliveryCps = ""
	v.WhiteDeliveryCps = ""
	v.UnusedWarehouseErrorMsg = ""
	v.UnusedDeliveryErrorMsg = ""
	v.UsedBlackDeliveryErrorMsg = ""
	v.PromiseOutboundTime = ""
	v.PromiseCollectTime = ""
	v.TradeId = 0
	v.SubTradeId = 0
	v.NeedConsignNum = 0
	v.NumIid = 0
	v.ItemRatio = 0
	v.BizDeliveryType = 0
	poolLogisticsInfo.Put(v)
}
