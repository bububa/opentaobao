package shop

import (
	"encoding/xml"

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

// AlibabaShopCategoryAllGetAPIResponseModel is 全部店铺分类信息查询接口 成功返回结果
type AlibabaShopCategoryAllGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_shop_category_all_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分类返回结果
	Result *AlibabaShopCategoryAllGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
