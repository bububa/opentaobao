package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
（新）关键词更新相关接口 APIResponse
taobao.simba.keyword.update

（新）关键词更新相关接口
*/
type TaobaoSimbaKeywordUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaKeywordUpdateResponse `json:"taobao_simba_keyword_update_response,omitempty"`
}

type TaobaoSimbaKeywordUpdateResponse struct {

    // 整体的返回值
    Results   []SiriusBidwordDto `json:"results,omitempty"`

    // 错误原因
    ErrorMsg   string `json:"error_msg,omitempty"`

}
