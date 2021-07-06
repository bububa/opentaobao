package bus

// OfflineRefundTicketPriceRq 结构体
type OfflineRefundTicketPriceRq struct {
	// 退票信息集合
	ListRefundInfos []SingleRefundInfo `json:"list_refund_infos,omitempty" xml:"list_refund_infos>single_refund_info,omitempty"`
	// top平台用户唯一识别
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 代理商店铺id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
}
