package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantproductEditAPIResponse 商家产品服务-编辑产品 API返回值
// alibaba.wdk.merchantproduct.edit
//
// 商家产品服务-编辑产品
type AlibabaWdkMerchantproductEditAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMerchantproductEditAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMerchantproductEditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMerchantproductEditAPIResponseModel).Reset()
}

// AlibabaWdkMerchantproductEditAPIResponseModel is 商家产品服务-编辑产品 成功返回结果
type AlibabaWdkMerchantproductEditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_merchantproduct_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品编辑返回结果
	Result *AlibabaWdkMerchantproductEditApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMerchantproductEditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMerchantproductEditAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMerchantproductEditAPIResponse)
	},
}

// GetAlibabaWdkMerchantproductEditAPIResponse 从 sync.Pool 获取 AlibabaWdkMerchantproductEditAPIResponse
func GetAlibabaWdkMerchantproductEditAPIResponse() *AlibabaWdkMerchantproductEditAPIResponse {
	return poolAlibabaWdkMerchantproductEditAPIResponse.Get().(*AlibabaWdkMerchantproductEditAPIResponse)
}

// ReleaseAlibabaWdkMerchantproductEditAPIResponse 将 AlibabaWdkMerchantproductEditAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMerchantproductEditAPIResponse(v *AlibabaWdkMerchantproductEditAPIResponse) {
	v.Reset()
	poolAlibabaWdkMerchantproductEditAPIResponse.Put(v)
}
