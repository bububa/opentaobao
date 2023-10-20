package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse 共享库存投保业务售后逆向订单数据获取 API返回值
// alibaba.wdkorder.sharestock.insurance.refundget
//
// 共享库存投保业务售后逆向订单数据获取
type AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkorderSharestockInsuranceRefundgetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkorderSharestockInsuranceRefundgetAPIResponseModel).Reset()
}

// AlibabaWdkorderSharestockInsuranceRefundgetAPIResponseModel is 共享库存投保业务售后逆向订单数据获取 成功返回结果
type AlibabaWdkorderSharestockInsuranceRefundgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdkorder_sharestock_insurance_refundget_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TopBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkorderSharestockInsuranceRefundgetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkorderSharestockInsuranceRefundgetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse)
	},
}

// GetAlibabaWdkorderSharestockInsuranceRefundgetAPIResponse 从 sync.Pool 获取 AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse
func GetAlibabaWdkorderSharestockInsuranceRefundgetAPIResponse() *AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse {
	return poolAlibabaWdkorderSharestockInsuranceRefundgetAPIResponse.Get().(*AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse)
}

// ReleaseAlibabaWdkorderSharestockInsuranceRefundgetAPIResponse 将 AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkorderSharestockInsuranceRefundgetAPIResponse(v *AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse) {
	v.Reset()
	poolAlibabaWdkorderSharestockInsuranceRefundgetAPIResponse.Put(v)
}
