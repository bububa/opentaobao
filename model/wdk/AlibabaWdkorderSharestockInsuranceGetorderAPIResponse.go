package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockInsuranceGetorderAPIResponse 共享库存订单投保消息获取 API返回值
// alibaba.wdkorder.sharestock.insurance.getorder
//
// 共享库存订单投保消息获取
type AlibabaWdkorderSharestockInsuranceGetorderAPIResponse struct {
	model.CommonResponse
	AlibabaWdkorderSharestockInsuranceGetorderAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkorderSharestockInsuranceGetorderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkorderSharestockInsuranceGetorderAPIResponseModel).Reset()
}

// AlibabaWdkorderSharestockInsuranceGetorderAPIResponseModel is 共享库存订单投保消息获取 成功返回结果
type AlibabaWdkorderSharestockInsuranceGetorderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdkorder_sharestock_insurance_getorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *MaochaoOrderInsuranceQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkorderSharestockInsuranceGetorderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkorderSharestockInsuranceGetorderAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkorderSharestockInsuranceGetorderAPIResponse)
	},
}

// GetAlibabaWdkorderSharestockInsuranceGetorderAPIResponse 从 sync.Pool 获取 AlibabaWdkorderSharestockInsuranceGetorderAPIResponse
func GetAlibabaWdkorderSharestockInsuranceGetorderAPIResponse() *AlibabaWdkorderSharestockInsuranceGetorderAPIResponse {
	return poolAlibabaWdkorderSharestockInsuranceGetorderAPIResponse.Get().(*AlibabaWdkorderSharestockInsuranceGetorderAPIResponse)
}

// ReleaseAlibabaWdkorderSharestockInsuranceGetorderAPIResponse 将 AlibabaWdkorderSharestockInsuranceGetorderAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkorderSharestockInsuranceGetorderAPIResponse(v *AlibabaWdkorderSharestockInsuranceGetorderAPIResponse) {
	v.Reset()
	poolAlibabaWdkorderSharestockInsuranceGetorderAPIResponse.Put(v)
}
