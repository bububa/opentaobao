package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据推广单元id获取关键词 APIResponse
taobao.simba.keyword.findbyadgroupid

根据一个关键词Id列表取得一组关键词
*/
type TaobaoSimbaKeywordFindbyadgroupidAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaKeywordFindbyadgroupidResponse `json:"simba_keyword_findbyadgroupid_response,omitempty"` 
    TaobaoSimbaKeywordFindbyadgroupidResponse
}

/* model for simplify = false
type TaobaoSimbaKeywordFindbyadgroupidResponse struct {

    // 整体的返回值
    
    Results  struct {
        SiriusBidwordDto  []SiriusBidwordDto `json:"sirius_bidword_dto,omitempty"`
    } `json:"results,omitempty"`
    

    // 错误原因
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

type TaobaoSimbaKeywordFindbyadgroupidResponse struct {

    // 整体的返回值
    Results   []SiriusBidwordDto `json:"results,omitempty"`

    // 错误原因
    ErrorMsg   string `json:"error_msg,omitempty"`

}
