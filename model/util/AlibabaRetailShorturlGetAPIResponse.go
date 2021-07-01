package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailShorturlGetAPIResponse
短链接获取 API返回值
alibaba.retail.shorturl.get

短链接获取 */
type AlibabaRetailShorturlGetAPIResponse struct {
	model.CommonResponse
	AlibabaRetailShorturlGetAPIResponseModel
}

// AlibabaRetailShorturlGetAPIResponseModel is 短链接获取 成功返回结果
type AlibabaRetailShorturlGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_shorturl_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaRetailShorturlGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
