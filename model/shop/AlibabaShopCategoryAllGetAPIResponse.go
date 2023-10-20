package shop

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaShopCategoryAllGetAPIResponse 全部店铺分类信息查询接口 API返回值
// alibaba.shop.category.all.get
//
// 按照卖家身份查询全部分类信息
type AlibabaShopCategoryAllGetAPIResponse struct {
	model.CommonResponse
	AlibabaShopCategoryAllGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaShopCategoryAllGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaShopCategoryAllGetAPIResponseModel).Reset()
}

// AlibabaShopCategoryAllGetAPIResponseModel is 全部店铺分类信息查询接口 成功返回结果
type AlibabaShopCategoryAllGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_shop_category_all_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分类返回结果
	Result *AlibabaShopCategoryAllGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaShopCategoryAllGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaShopCategoryAllGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaShopCategoryAllGetAPIResponse)
	},
}

// GetAlibabaShopCategoryAllGetAPIResponse 从 sync.Pool 获取 AlibabaShopCategoryAllGetAPIResponse
func GetAlibabaShopCategoryAllGetAPIResponse() *AlibabaShopCategoryAllGetAPIResponse {
	return poolAlibabaShopCategoryAllGetAPIResponse.Get().(*AlibabaShopCategoryAllGetAPIResponse)
}

// ReleaseAlibabaShopCategoryAllGetAPIResponse 将 AlibabaShopCategoryAllGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaShopCategoryAllGetAPIResponse(v *AlibabaShopCategoryAllGetAPIResponse) {
	v.Reset()
	poolAlibabaShopCategoryAllGetAPIResponse.Put(v)
}
