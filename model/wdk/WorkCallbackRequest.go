package wdk

// WorkCallbackRequest 结构体
type WorkCallbackRequest struct {
	// 子单列表
	WorkCallbackSubOrderInfoList []WorkCallbackSubOrderInfo `json:"work_callback_sub_order_info_list,omitempty" xml:"work_callback_sub_order_info_list>work_callback_sub_order_info,omitempty"`
	// 经营店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 回传通知状态 ACCEPTED = 商户接单 REJECTED = 商户取消订单 PICKED = 拣货完成 PACKAGED = 打包出库 SHIPPING = 开始配送 SIGN = 用户签收 REFUSED = 用户拒收
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 状态备注, 如商户拒单原因备注
	StatusRemark string `json:"status_remark,omitempty" xml:"status_remark,omitempty"`
	// 配送人员名称 SHIPPING/SIGN状态必填
	DelivererName string `json:"deliverer_name,omitempty" xml:"deliverer_name,omitempty"`
	// 配送人员联系方式 SHIPPING/SIGN状态必填
	DelivererPhone string `json:"deliverer_phone,omitempty" xml:"deliverer_phone,omitempty"`
	// 配送公司编码 FENGNIAO = 蜂鸟 MEITUAN = 美团 DADA = 达达 SHUNFENG = 顺丰 ELEZB = 饿了么众包 BINGEX = 闪送 SELF = 商家自送 OTHER = 其他运力
	DelivererCompany string `json:"deliverer_company,omitempty" xml:"deliverer_company,omitempty"`
	// 配送物流单号
	LogisticsNo string `json:"logistics_no,omitempty" xml:"logistics_no,omitempty"`
	// 业务订单编码
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}
