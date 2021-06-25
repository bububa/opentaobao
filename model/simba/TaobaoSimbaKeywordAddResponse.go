package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
（新）关键词新增接口 APIResponse
taobao.simba.keyword.add

（新）关键词更新相关接口
*/
type TaobaoSimbaKeywordAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaKeywordAddResponse `json:"taobao_simba_keyword_add_response,omitempty"`
}

type TaobaoSimbaKeywordAddResponse struct {

    // 整体的返回值
    Results   []SiriusBidwordDto `json:"results,omitempty"`

    // 错误原因
    ErrorMsg   string `json:"error_msg,omitempty"`

}
