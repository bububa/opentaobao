package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔扩展码查询 APIResponse
taobao.jst.sms.extendcode.query

聚石塔扩展码查询
*/
type TaobaoJstSmsExtendcodeQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoJstSmsExtendcodeQueryResponse `json:"taobao_jst_sms_extendcode_query_response,omitempty"`
}

type TaobaoJstSmsExtendcodeQueryResponse struct {

    // 系统异常
    ResponseCode   string `json:"response_code,omitempty"`

    // 成功
    ResponseSuccess   bool `json:"response_success,omitempty"`

    // 请求id
    ResponseId   string `json:"response_id,omitempty"`

    // 成功
    Module   string `json:"module,omitempty"`

    // 系统异常
    Message   string `json:"message,omitempty"`

}
