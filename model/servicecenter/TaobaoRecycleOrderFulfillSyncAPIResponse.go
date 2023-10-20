package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRecycleOrderFulfillSyncAPIResponse 同步回收单最终履约方式 API返回值
// taobao.recycle.order.fulfill.sync
//
// 同步回收单最终履约方式
type TaobaoRecycleOrderFulfillSyncAPIResponse struct {
	model.CommonResponse
	TaobaoRecycleOrderFulfillSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoRecycleOrderFulfillSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoRecycleOrderFulfillSyncAPIResponseModel).Reset()
}

// TaobaoRecycleOrderFulfillSyncAPIResponseModel is 同步回收单最终履约方式 成功返回结果
type TaobaoRecycleOrderFulfillSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"recycle_order_fulfill_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 同步回收单最终履约方式结果
	Data *OfnRecycleOrderSyncFulfillTypeDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoRecycleOrderFulfillSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoRecycleOrderFulfillSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoRecycleOrderFulfillSyncAPIResponse)
	},
}

// GetTaobaoRecycleOrderFulfillSyncAPIResponse 从 sync.Pool 获取 TaobaoRecycleOrderFulfillSyncAPIResponse
func GetTaobaoRecycleOrderFulfillSyncAPIResponse() *TaobaoRecycleOrderFulfillSyncAPIResponse {
	return poolTaobaoRecycleOrderFulfillSyncAPIResponse.Get().(*TaobaoRecycleOrderFulfillSyncAPIResponse)
}

// ReleaseTaobaoRecycleOrderFulfillSyncAPIResponse 将 TaobaoRecycleOrderFulfillSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoRecycleOrderFulfillSyncAPIResponse(v *TaobaoRecycleOrderFulfillSyncAPIResponse) {
	v.Reset()
	poolTaobaoRecycleOrderFulfillSyncAPIResponse.Put(v)
}
