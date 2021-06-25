package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
（新）根据一堆关键词ids获取关键词 APIResponse
taobao.simba.keyword.findbyids

根据一个关键词Id列表取得一组关键词
*/
type TaobaoSimbaKeywordFindbyidsAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaKeywordFindbyidsResponse `json:"taobao_simba_keyword_findbyids_response,omitempty"`
}

type TaobaoSimbaKeywordFindbyidsResponse struct {

    // 整体的返回值
    Results   []SiriusBidwordDto `json:"results,omitempty"`

    // 错误原因
    ErrorMsg   string `json:"error_msg,omitempty"`

}
