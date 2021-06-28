package elife

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询交易结果 APIResponse
taobao.elife.lifecard.query

卖家在交易状态不明的情况下, 查询交易结果.
*/
type TaobaoElifeLifecardQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoElifeLifecardQueryResponse `json:"elife_lifecard_query_response,omitempty"` 
    TaobaoElifeLifecardQueryResponse
}

/* model for simplify = false
type TaobaoElifeLifecardQueryResponse struct {

    // resultMsg
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // amount
    
    Amount   int64 `json:"amount,omitempty"`
    

    // successed
    
    Successed   bool `json:"successed,omitempty"`
    

    // resultCode
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // inflateAmount
    
    InflateAmount   int64 `json:"inflate_amount,omitempty"`
    

}
*/

type TaobaoElifeLifecardQueryResponse struct {

    // resultMsg
    ResultMsg   string `json:"result_msg,omitempty"`

    // amount
    Amount   int64 `json:"amount,omitempty"`

    // successed
    Successed   bool `json:"successed,omitempty"`

    // resultCode
    ResultCode   string `json:"result_code,omitempty"`

    // inflateAmount
    InflateAmount   int64 `json:"inflate_amount,omitempty"`

}
