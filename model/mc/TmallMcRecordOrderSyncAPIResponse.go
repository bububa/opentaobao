package mc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMcRecordOrderSyncAPIResponse 订单信息同步 API返回值
// tmall.mc.record.order.sync
//
// 订单信息同步(零售云接口)
type TmallMcRecordOrderSyncAPIResponse struct {
	model.CommonResponse
	TmallMcRecordOrderSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TmallMcRecordOrderSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallMcRecordOrderSyncAPIResponseModel).Reset()
}

// TmallMcRecordOrderSyncAPIResponseModel is 订单信息同步 成功返回结果
type TmallMcRecordOrderSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_mc_record_order_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 同步成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TmallMcRecordOrderSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = false
}

var poolTmallMcRecordOrderSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallMcRecordOrderSyncAPIResponse)
	},
}

// GetTmallMcRecordOrderSyncAPIResponse 从 sync.Pool 获取 TmallMcRecordOrderSyncAPIResponse
func GetTmallMcRecordOrderSyncAPIResponse() *TmallMcRecordOrderSyncAPIResponse {
	return poolTmallMcRecordOrderSyncAPIResponse.Get().(*TmallMcRecordOrderSyncAPIResponse)
}

// ReleaseTmallMcRecordOrderSyncAPIResponse 将 TmallMcRecordOrderSyncAPIResponse 保存到 sync.Pool
func ReleaseTmallMcRecordOrderSyncAPIResponse(v *TmallMcRecordOrderSyncAPIResponse) {
	v.Reset()
	poolTmallMcRecordOrderSyncAPIResponse.Put(v)
}
