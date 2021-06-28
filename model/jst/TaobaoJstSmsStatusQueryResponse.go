package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔公众号状态查询 APIResponse
taobao.jst.sms.status.query

聚石塔公众号状态查询
*/
type TaobaoJstSmsStatusQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJstSmsStatusQueryResponse `json:"jst_sms_status_query_response,omitempty"` 
    TaobaoJstSmsStatusQueryResponse
}

/* model for simplify = false
type TaobaoJstSmsStatusQueryResponse struct {

    // 返回值
    
    Result  *struct {
        SmsResponse  *SmsResponse `json:"sms_response,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoJstSmsStatusQueryResponse struct {

    // 返回值
    Result   *SmsResponse `json:"result,omitempty"`

}
