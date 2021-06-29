package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索热词 API返回值 
alibaba.xiami.api.search.hotwords.get

搜索热词
*/
type AlibabaXiamiApiSearchHotwordsGetAPIResponse struct {
    model.CommonResponse
    AlibabaXiamiApiSearchHotwordsGetResponse
}

// 搜索热词 成功返回结果
type AlibabaXiamiApiSearchHotwordsGetResponse struct {
    XMLName xml.Name `xml:"alibaba_xiami_api_search_hotwords_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    SearchHotWordsResult   *AlibabaXiamiApiSearchHotwordsGetStruct `json:"search_hot_words_result,omitempty" xml:"search_hot_words_result,omitempty"`
}
