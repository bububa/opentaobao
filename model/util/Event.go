package util

// Event 结构体
type Event struct {
	// 淘宝订单号
	Tid string `json:"tid,omitempty" xml:"tid,omitempty"`
	// 扩展属性
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 事件状态，如QIMEN_ERP_TRANSFER，QIMEN_ERP_CHECK
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 商家平台编码.MAIN:官方渠道,JD:京东,DD:当当,PP:拍拍,YX:易讯,EBAY:ebay,AMAZON:亚马逊,SN:苏宁,GM:国美,WPH:唯品会,JM:聚美,MGJ:蘑菇街,YT:银泰,YHD:1号店,1688:1688,POS:POS门店,OTHER:其他
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
	// 外部商家名称。必须同时填写platform
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 主单号对应的erp单号，转单、审单、通知配货、出库 需要填。拆单、合单场景下不用填
	ErpOrderId string `json:"erp_order_id,omitempty" xml:"erp_order_id,omitempty"`
	// 淘宝子订单id（拆单、合单场景下不用填，其他场景需要回传,用英文逗号隔开）
	TaobaoSubOrderIds string `json:"taobao_sub_order_ids,omitempty" xml:"taobao_sub_order_ids,omitempty"`
	// 触发事件的时间
	EventTime string `json:"event_time,omitempty" xml:"event_time,omitempty"`
	// 订单创建时间,数字
	Create int64 `json:"create,omitempty" xml:"create,omitempty"`
}
