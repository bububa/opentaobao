package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索热词 APIResponse
alibaba.xiami.api.search.hotwords.get

搜索热词
*/
type AlibabaXiamiApiSearchHotwordsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_xiami_api_search_hotwords_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    SearchHotWordsResult   *AlibabaXiamiApiSearchHotwordsGetStruct `json:"search_hot_words_result,omitempty" xml:"