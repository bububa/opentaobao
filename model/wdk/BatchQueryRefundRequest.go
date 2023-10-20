package wdk

import (
	"sync"
)

// BatchQueryRefundRequest 结构体
type BatchQueryRefundRequest struct {
	// 店铺ID列表，order_from=4时必填，其他非必填
	StoreIds []string `json:"store_ids,omitempty" xml:"store_ids>string,omitempty"`
	// 结束时间，必填
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 下单终端：APP/POS，非必填
	OrderClient string `json:"order_client,omitempty" xml:"order_client,omitempty"`
	// 开始时间，必填
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 废弃字段
	SyncStatus string `json:"sync_status,omitempty" xml:"sync_status,omitempty"`
	// 渠道店id，order_from=4时非必填，其他必填
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 经营店id，非必填
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// aelophy 翱象商家必填
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 页码，从0开始，必填
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 页大小，必填
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 渠道来源 3：饿了么 4：盒马&amp;淘鲜达 20：商家自有渠道，必填
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
}

var poolBatchQueryRefundRequest = sync.Pool{
	New: func() any {
		return new(BatchQueryRefundRequest)
	},
}

// GetBatchQueryRefundRequest() 从对象池中获取BatchQueryRefundRequest
func GetBatchQueryRefundRequest() *BatchQueryRefundRequest {
	return poolBatchQueryRefundRequest.Get().(*BatchQueryRefundRequest)
}

// ReleaseBatchQueryRefundRequest 释放BatchQueryRefundRequest
func ReleaseBatchQueryRefundRequest(v *BatchQueryRefundRequest) {
	v.StoreIds = v.StoreIds[:0]
	v.EndTime = ""
	v.OrderClient = ""
	v.StartTime = ""
	v.SyncStatus = ""
	v.ShopId = ""
	v.StoreId = ""
	v.BizType = ""
	v.PageIndex = 0
	v.PageSize = 0
	v.OrderFrom = 0
	poolBatchQueryRefundRequest.Put(v)
}
