package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
OAID 收件人信息解密接口 APIResponse
taobao.qimen.receiverinfo.query

WMS 调用该接口，通过 OAID 查询解密后的收件人信息
*/
type TaobaoQimenReceiverinfoQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenReceiverinfoQueryResponse `json:"qimen_receiverinfo_query_response,omitempty"` 
    TaobaoQimenReceiverinfoQueryResponse
}

/* model for simplify = false
type TaobaoQimenReceiverinfoQueryResponse struct {

    // 
    
    Response  *struct {
        Response  *Response `json:"response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenReceiverinfoQueryResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
