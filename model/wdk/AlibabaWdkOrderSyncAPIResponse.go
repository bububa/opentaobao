package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOrderSyncAPIResponse 五道口外部订单同步 API返回值
// alibaba.wdk.order.sync
//
// 外部商户使用自助POS下单订单同步到五道口
type AlibabaWdkOrderSyncAPIResponse struct {
	model.CommonResponse
	AlibabaWdkOrderSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkOrderSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkOrderSyncAPIResponseModel).Reset()
}

// AlibabaWdkOrderSyncAPIResponseModel is 五道口外部订单同步 成功返回结果
type AlibabaWdkOrderSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_order_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 订单号
	Target string `json:"target,omitempty" xml:"target,omitempty"`
	// 返回码
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkOrderSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Target = ""
	m.ReturnCode = 0
	m.IsSuccess = false
}

var poolAlibabaWdkOrderSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkOrderSyncAPIResponse)
	},
}

// GetAlibabaWdkOrderSyncAPIResponse 从 sync.Pool 获取 AlibabaWdkOrderSyncAPIResponse
func GetAlibabaWdkOrderSyncAPIResponse() *AlibabaWdkOrderSyncAPIResponse {
	return poolAlibabaWdkOrderSyncAPIResponse.Get().(*AlibabaWdkOrderSyncAPIResponse)
}

// ReleaseAlibabaWdkOrderSyncAPIResponse 将 AlibabaWdkOrderSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkOrderSyncAPIResponse(v *AlibabaWdkOrderSyncAPIResponse) {
	v.Reset()
	poolAlibabaWdkOrderSyncAPIResponse.Put(v)
}
