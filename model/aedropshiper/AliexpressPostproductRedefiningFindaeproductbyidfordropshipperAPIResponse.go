package aedropshiper

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse Dropshipper查找商品信息接口 API返回值
// aliexpress.postproduct.redefining.findaeproductbyidfordropshipper
//
// 提供给Dropshipper的通过商品ID查找商品信息的接口，只有特定买家可以使用
type AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse struct {
	model.CommonResponse
	AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponseModel).Reset()
}

// AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponseModel is Dropshipper查找商品信息接口 成功返回结果
type AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_postproduct_redefining_findaeproductbyidfordropshipper_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AeopFindProductResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse)
	},
}

// GetAliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse 从 sync.Pool 获取 AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse
func GetAliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse() *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse {
	return poolAliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse.Get().(*AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse)
}

// ReleaseAliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse 将 AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse 保存到 sync.Pool
func ReleaseAliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse(v *AliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse) {
	v.Reset()
	poolAliexpressPostproductRedefiningFindaeproductbyidfordropshipperAPIResponse.Put(v)
}
