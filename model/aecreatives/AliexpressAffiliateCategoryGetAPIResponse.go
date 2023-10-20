package aecreatives

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateCategoryGetAPIResponse AE流量推广类目信息获取API API返回值
// aliexpress.affiliate.category.get
//
// 获取AE流量推广类目的API
type AliexpressAffiliateCategoryGetAPIResponse struct {
	model.CommonResponse
	AliexpressAffiliateCategoryGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressAffiliateCategoryGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressAffiliateCategoryGetAPIResponseModel).Reset()
}

// AliexpressAffiliateCategoryGetAPIResponseModel is AE流量推广类目信息获取API 成功返回结果
type AliexpressAffiliateCategoryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_affiliate_category_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	RespResult *ResponseResult `json:"resp_result,omitempty" xml:"resp_result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressAffiliateCategoryGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespResult = nil
}

var poolAliexpressAffiliateCategoryGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressAffiliateCategoryGetAPIResponse)
	},
}

// GetAliexpressAffiliateCategoryGetAPIResponse 从 sync.Pool 获取 AliexpressAffiliateCategoryGetAPIResponse
func GetAliexpressAffiliateCategoryGetAPIResponse() *AliexpressAffiliateCategoryGetAPIResponse {
	return poolAliexpressAffiliateCategoryGetAPIResponse.Get().(*AliexpressAffiliateCategoryGetAPIResponse)
}

// ReleaseAliexpressAffiliateCategoryGetAPIResponse 将 AliexpressAffiliateCategoryGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressAffiliateCategoryGetAPIResponse(v *AliexpressAffiliateCategoryGetAPIResponse) {
	v.Reset()
	poolAliexpressAffiliateCategoryGetAPIResponse.Put(v)
}
