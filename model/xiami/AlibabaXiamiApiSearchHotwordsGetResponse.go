package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
搜索热词 APIResponse
alibaba.xiami.api.search.hotwords.get

搜索热词
*/
type AlibabaXiamiApiSearchHotwordsGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaXiamiApiSearchHotwordsGetResponse `json:"alibaba_xiami_api_search_hotwords_get_response,omitempty"`
}

type AlibabaXiamiApiSearchHotwordsGetResponse struct {

    // 返回结果
    SearchHotWordsResult   *AlibabaXiamiApiSearchHotwordsGetStruct `json:"search_hot_words_result,omitempty"`

}
