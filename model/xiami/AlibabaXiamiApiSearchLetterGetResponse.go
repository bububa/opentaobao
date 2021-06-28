package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
搜索接口（首字母） APIResponse
alibaba.xiami.api.search.letter.get

搜索接口（首字母）
*/
type AlibabaXiamiApiSearchLetterGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaXiamiApiSearchLetterGetResponse `json:"alibaba_xiami_api_search_letter_get_response,omitempty"` 
    AlibabaXiamiApiSearchLetterGetResponse
}

/* model for simplify = false
type AlibabaXiamiApiSearchLetterGetResponse struct {

    // search_letter_result
    
    SearchLetterResult  *struct {
        AlibabaXiamiApiSearchLetterGetStruct  *AlibabaXiamiApiSearchLetterGetStruct `json:"alibaba_xiami_api_search_letter_get_struct,omitempty"`
    } `json:"search_letter_result,omitempty"`
    

}
*/

type AlibabaXiamiApiSearchLetterGetResponse struct {

    // search_letter_result
    SearchLetterResult   *AlibabaXiamiApiSearchLetterGetStruct `json:"search_letter_result,omitempty"`

}
