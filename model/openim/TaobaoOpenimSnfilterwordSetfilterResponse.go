package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
关键词过滤 APIResponse
taobao.openim.snfilterword.setfilter

设置openim关键词过滤
*/
type TaobaoOpenimSnfilterwordSetfilterAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenimSnfilterwordSetfilterResponse `json:"taobao_openim_snfilterword_setfilter_response,omitempty"`
}

type TaobaoOpenimSnfilterwordSetfilterResponse struct {

    // 成功
    Errid   int64 `json:"errid,omitempty"`

    // 错误原因
    Errmsg   string `json:"errmsg,omitempty"`

}
