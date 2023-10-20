package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse 共享库存逆向订单理赔单回传 API返回值
// alibaba.wdkorder.sharestock.insurance.refundcallback
//
// 共享库存逆向订单理赔单回传
type AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse struct {
	model.CommonResponse
	AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponseModel).Reset()
}

// AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponseModel is 共享库存逆向订单理赔单回传 成功返回结果
type AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdkorder_sharestock_insurance_refundcallback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *MaochaoOrderInsuranceRefundCallbackResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse)
	},
}

// GetAlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse 从 sync.Pool 获取 AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse
func GetAlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse() *AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse {
	return poolAlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse.Get().(*AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse)
}

// ReleaseAlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse 将 AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse(v *AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse) {
	v.Reset()
	poolAlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse.Put(v)
}
