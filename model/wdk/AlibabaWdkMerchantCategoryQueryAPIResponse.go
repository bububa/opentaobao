package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMerchantCategoryQueryAPIResponse 三江erp对接类目查询接口 API返回值
// alibaba.wdk.merchant.category.query
//
// 三江erp对接类目查询接口
type AlibabaWdkMerchantCategoryQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMerchantCategoryQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMerchantCategoryQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMerchantCategoryQueryAPIResponseModel).Reset()
}

// AlibabaWdkMerchantCategoryQueryAPIResponseModel is 三江erp对接类目查询接口 成功返回结果
type AlibabaWdkMerchantCategoryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_merchant_category_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorCode
	Errorcode string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`
	// errorDesc
	Errordesc string `json:"errordesc,omitempty" xml:"errordesc,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// success
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMerchantCategoryQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errorcode = ""
	m.Errordesc = ""
	m.Result = ""
	m.Suc = false
}

var poolAlibabaWdkMerchantCategoryQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMerchantCategoryQueryAPIResponse)
	},
}

// GetAlibabaWdkMerchantCategoryQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkMerchantCategoryQueryAPIResponse
func GetAlibabaWdkMerchantCategoryQueryAPIResponse() *AlibabaWdkMerchantCategoryQueryAPIResponse {
	return poolAlibabaWdkMerchantCategoryQueryAPIResponse.Get().(*AlibabaWdkMerchantCategoryQueryAPIResponse)
}

// ReleaseAlibabaWdkMerchantCategoryQueryAPIResponse 将 AlibabaWdkMerchantCategoryQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMerchantCategoryQueryAPIResponse(v *AlibabaWdkMerchantCategoryQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkMerchantCategoryQueryAPIResponse.Put(v)
}
