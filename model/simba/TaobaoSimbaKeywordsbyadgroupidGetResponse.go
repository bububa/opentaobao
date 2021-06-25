package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广组的所有关键词 APIResponse
taobao.simba.keywordsbyadgroupid.get

取得一个推广组的所有关键词
*/
type TaobaoSimbaKeywordsbyadgroupidGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaKeywordsbyadgroupidGetResponse `json:"taobao_simba_keywordsbyadgroupid_get_response,omitempty"`
}

type TaobaoSimbaKeywordsbyadgroupidGetResponse struct {

    // 取得的关键词列表
    Keywords   []Keyword `json:"keywords,omitempty"`

}
