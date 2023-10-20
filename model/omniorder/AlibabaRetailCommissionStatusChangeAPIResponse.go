package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailCommissionStatusChangeAPIResponse 分佣状态变更 API返回值
// alibaba.retail.commission.status.change
//
// 分佣系统，分佣状态变更接口
type AlibabaRetailCommissionStatusChangeAPIResponse struct {
	model.CommonResponse
	AlibabaRetailCommissionStatusChangeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailCommissionStatusChangeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailCommissionStatusChangeAPIResponseModel).Reset()
}

// AlibabaRetailCommissionStatusChangeAPIResponseModel is 分佣状态变更 成功返回结果
type AlibabaRetailCommissionStatusChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_commission_status_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回描述信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回的执行状态吗
	SCode string `json:"s_code,omitempty" xml:"s_code,omitempty"`
	// 返回的数据实体
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否执行成功
	SSuccess bool `json:"s_success,omitempty" xml:"s_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailCommissionStatusChangeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.SCode = ""
	m.Data = false
	m.SSuccess = false
}

var poolAlibabaRetailCommissionStatusChangeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailCommissionStatusChangeAPIResponse)
	},
}

// GetAlibabaRetailCommissionStatusChangeAPIResponse 从 sync.Pool 获取 AlibabaRetailCommissionStatusChangeAPIResponse
func GetAlibabaRetailCommissionStatusChangeAPIResponse() *AlibabaRetailCommissionStatusChangeAPIResponse {
	return poolAlibabaRetailCommissionStatusChangeAPIResponse.Get().(*AlibabaRetailCommissionStatusChangeAPIResponse)
}

// ReleaseAlibabaRetailCommissionStatusChangeAPIResponse 将 AlibabaRetailCommissionStatusChangeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailCommissionStatusChangeAPIResponse(v *AlibabaRetailCommissionStatusChangeAPIResponse) {
	v.Reset()
	poolAlibabaRetailCommissionStatusChangeAPIResponse.Put(v)
}
