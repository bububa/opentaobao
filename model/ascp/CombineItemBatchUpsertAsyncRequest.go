package ascp

// CombineItemBatchUpsertAsyncRequest 结构体
type CombineItemBatchUpsertAsyncRequest struct {
	// 组合货品列表
	CombineScItems []CombineScItem `json:"combine_sc_items,omitempty" xml:"combine_sc_items>combine_sc_item,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ERP调翱象接口创建货品的时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
