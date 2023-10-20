package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkCouponContractCreateAPIResponse 营销券合同创建接口 API返回值
// alibaba.wdk.coupon.contract.create
//
// 营销券合同创建接口
type AlibabaWdkCouponContractCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkCouponContractCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkCouponContractCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkCouponContractCreateAPIResponseModel).Reset()
}

// AlibabaWdkCouponContractCreateAPIResponseModel is 营销券合同创建接口 成功返回结果
type AlibabaWdkCouponContractCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_coupon_contract_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaWdkCouponContractCreateApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkCouponContractCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkCouponContractCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponContractCreateAPIResponse)
	},
}

// GetAlibabaWdkCouponContractCreateAPIResponse 从 sync.Pool 获取 AlibabaWdkCouponContractCreateAPIResponse
func GetAlibabaWdkCouponContractCreateAPIResponse() *AlibabaWdkCouponContractCreateAPIResponse {
	return poolAlibabaWdkCouponContractCreateAPIResponse.Get().(*AlibabaWdkCouponContractCreateAPIResponse)
}

// ReleaseAlibabaWdkCouponContractCreateAPIResponse 将 AlibabaWdkCouponContractCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkCouponContractCreateAPIResponse(v *AlibabaWdkCouponContractCreateAPIResponse) {
	v.Reset()
	poolAlibabaWdkCouponContractCreateAPIResponse.Put(v)
}
