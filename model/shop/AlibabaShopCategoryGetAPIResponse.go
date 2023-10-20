package shop

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabashopcategorygetAPIResponse 指定店铺分类信息查询接口 API返回值
// alibaba.shop.category.get
//
// 按照卖家身份查询指定分类信息
type AlibabashopcategorygetAPIResponse struct {
	model.CommonResponse
	AlibabashopcategorygetAPIResponseModel
}

// AlibabashopcategorygetAPIResponseModel is 指定店铺分类信息查询接口 成功返回结果
type AlibabashopcategorygetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_shop_category_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分类返回结果
	Result *AlibabashopcategorygetResult `json:"result,omitempty" xml:"result,omitempty"`
}
