package bus

// OfflineRefundTicketPriceRq 
type OfflineRefundTicketPriceRq struct {

    // 代理商店铺id
    AgentId   int64 `json:"agent_id,omitempty"`

    // top平台用户唯一识别
    AppKey   string `json:"app_key,omitempty"`

    // 退票信息集合
    ListRefundInfos   []SingleRefundInfo `json:"list_refund_infos,omitempty"`

}
