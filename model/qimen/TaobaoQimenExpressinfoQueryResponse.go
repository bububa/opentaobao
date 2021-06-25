package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
配送公司信息查询接口 APIResponse
taobao.qimen.expressinfo.query

配送公司信息查询
*/
type TaobaoQimenExpressinfoQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenExpressinfoQueryResponse `json:"taobao_qimen_expressinfo_query_response,omitempty"`
}

type TaobaoQimenExpressinfoQueryResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
