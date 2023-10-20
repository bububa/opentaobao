package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkShopQueryAPIResponse 门店查询接口 API返回值
// alibaba.wdk.shop.query
//
// 根据门店code查询门店信息
type AlibabaWdkShopQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkShopQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkShopQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkShopQueryAPIResponseModel).Reset()
}

// AlibabaWdkShopQueryAPIResponseModel is 门店查询接口 成功返回结果
type AlibabaWdkShopQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_shop_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabaWdkShopQueryApiResults `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkShopQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkShopQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkShopQueryAPIResponse)
	},
}

// GetAlibabaWdkShopQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkShopQueryAPIResponse
func GetAlibabaWdkShopQueryAPIResponse() *AlibabaWdkShopQueryAPIResponse {
	return poolAlibabaWdkShopQueryAPIResponse.Get().(*AlibabaWdkShopQueryAPIResponse)
}

// ReleaseAlibabaWdkShopQueryAPIResponse 将 AlibabaWdkShopQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkShopQueryAPIResponse(v *AlibabaWdkShopQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkShopQueryAPIResponse.Put(v)
}
