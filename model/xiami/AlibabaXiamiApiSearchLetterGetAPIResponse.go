package xiami

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXiamiApiSearchLetterGetAPIResponse 搜索接口（首字母） API返回值
// alibaba.xiami.api.search.letter.get
//
// 搜索接口（首字母）
type AlibabaXiamiApiSearchLetterGetAPIResponse struct {
	model.CommonResponse
	AlibabaXiamiApiSearchLetterGetAPIResponseModel
}

// AlibabaXiamiApiSearchLetterGetAPIResponseModel is 搜索接口（首字母） 成功返回结果
type AlibabaXiamiApiSearchLetterGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_xiami_api_search_letter_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// search_letter_result
	SearchLetterResult *AlibabaXiamiApiSearchLetterGetStruct `json:"search_letter_result,omitempty" xml:"search_letter_result,omitempty"`
}
