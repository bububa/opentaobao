package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingversiongenerateAPIResponse 生成发布使用的版本号 API返回值
// alibaba.hm.marketing.version.generate
//
// 生成发布使用的版本号
type AlibabahmmarketingversiongenerateAPIResponse struct {
	model.CommonResponse
	AlibabahmmarketingversiongenerateAPIResponseModel
}

// AlibabahmmarketingversiongenerateAPIResponseModel is 生成发布使用的版本号 成功返回结果
type AlibabahmmarketingversiongenerateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_version_generate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
