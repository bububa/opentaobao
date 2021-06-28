package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索接口（首字母） APIResponse
alibaba.xiami.api.search.letter.get

搜索接口（首字母）
*/
type AlibabaXiamiApiSearchLetterGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_xiami_api_search_letter_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // search_letter_result
    
    SearchLetterResult   *AlibabaXiamiApiSearchLetterGetStruct `json:"search_letter_result,omitempty" xml:"