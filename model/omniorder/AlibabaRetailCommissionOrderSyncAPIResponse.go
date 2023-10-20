package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailCommissionOrderSyncAPIResponse 分佣数据传输 API返回值
// alibaba.retail.commission.order.sync
//
// 同步分佣结果
type AlibabaRetailCommissionOrderSyncAPIResponse struct {
	model.CommonResponse
	AlibabaRetailCommissionOrderSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailCommissionOrderSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailCommissionOrderSyncAPIResponseModel).Reset()
}

// AlibabaRetailCommissionOrderSyncAPIResponseModel is 分佣数据传输 成功返回结果
type AlibabaRetailCommissionOrderSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_commission_order_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回的执行状态吗
	SCode string `json:"s_code,omitempty" xml:"s_code,omitempty"`
	// 返回的数据实体
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否执行成功
	SSuccess bool `json:"s_success,omitempty" xml:"s_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailCommissionOrderSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.SCode = ""
	m.Data = false
	m.SSuccess = false
}

var poolAlibabaRetailCommissionOrderSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailCommissionOrderSyncAPIResponse)
	},
}

// GetAlibabaRetailCommissionOrderSyncAPIResponse 从 sync.Pool 获取 AlibabaRetailCommissionOrderSyncAPIResponse
func GetAlibabaRetailCommissionOrderSyncAPIResponse() *AlibabaRetailCommissionOrderSyncAPIResponse {
	return poolAlibabaRetailCommissionOrderSyncAPIResponse.Get().(*AlibabaRetailCommissionOrderSyncAPIResponse)
}

// ReleaseAlibabaRetailCommissionOrderSyncAPIResponse 将 AlibabaRetailCommissionOrderSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailCommissionOrderSyncAPIResponse(v *AlibabaRetailCommissionOrderSyncAPIResponse) {
	v.Reset()
	poolAlibabaRetailCommissionOrderSyncAPIResponse.Put(v)
}
