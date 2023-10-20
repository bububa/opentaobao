package shop

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaShopCategoryGetAPIResponse 指定店铺分类信息查询接口 API返回值
// alibaba.shop.category.get
//
// 按照卖家身份查询指定分类信息
type AlibabaShopCategoryGetAPIResponse struct {
	model.CommonResponse
	AlibabaShopCategoryGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaShopCategoryGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaShopCategoryGetAPIResponseModel).Reset()
}

// AlibabaShopCategoryGetAPIResponseModel is 指定店铺分类信息查询接口 成功返回结果
type AlibabaShopCategoryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_shop_category_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分类返回结果
	Result *AlibabaShopCategoryGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaShopCategoryGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaShopCategoryGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaShopCategoryGetAPIResponse)
	},
}

// GetAlibabaShopCategoryGetAPIResponse 从 sync.Pool 获取 AlibabaShopCategoryGetAPIResponse
func GetAlibabaShopCategoryGetAPIResponse() *AlibabaShopCategoryGetAPIResponse {
	return poolAlibabaShopCategoryGetAPIResponse.Get().(*AlibabaShopCategoryGetAPIResponse)
}

// ReleaseAlibabaShopCategoryGetAPIResponse 将 AlibabaShopCategoryGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaShopCategoryGetAPIResponse(v *AlibabaShopCategoryGetAPIResponse) {
	v.Reset()
	poolAlibabaShopCategoryGetAPIResponse.Put(v)
}
