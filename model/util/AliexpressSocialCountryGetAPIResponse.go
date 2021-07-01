package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSocialCountryGetAPIResponse
获取国家列表 API返回值
aliexpress.social.country.get

获取目前AE支持的国家列表 */
type AliexpressSocialCountryGetAPIResponse struct {
	model.CommonResponse
	AliexpressSocialCountryGetAPIResponseModel
}

// AliexpressSocialCountryGetAPIResponseModel is 获取国家列表 成功返回结果
type AliexpressSocialCountryGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_social_country_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ItemPickPagingResult
	Result *ItemPickPagingResult `json:"result,omitempty" xml:"result,omitempty"`
}
