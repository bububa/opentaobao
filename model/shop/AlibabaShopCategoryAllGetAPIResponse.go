package shop

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabashopcategoryallgetAPIResponse 全部店铺分类信息查询接口 API返回值
// alibaba.shop.category.all.get
//
// 按照卖家身份查询全部分类信息
type AlibabashopcategoryallgetAPIResponse struct {
	model.CommonResponse
	AlibabashopcategoryallgetAPIResponseModel
}

// AlibabashopcategoryallgetAPIResponseModel is 全部店铺分类信息查询接口 成功返回结果
type AlibabashopcategoryallgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_shop_category_all_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分类返回结果
	Result *AlibabashopcategoryallgetResult `json:"result,omitempty" xml:"result,omitempty"`
}
