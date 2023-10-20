package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssocialcurrencygetAPIResponse 币种获取接口 API返回值
// aliexpress.social.currency.get
//
// 获取目前AE社交支持的币种
type AliexpresssocialcurrencygetAPIResponse struct {
	model.CommonResponse
	AliexpresssocialcurrencygetAPIResponseModel
}

// AliexpresssocialcurrencygetAPIResponseModel is 币种获取接口 成功返回结果
type AliexpresssocialcurrencygetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_social_currency_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 包类型
	Result *ItemPickPagingResult `json:"result,omitempty" xml:"result,omitempty"`
}
