package tvpay

import (
    "github.com/bububa/opentaobao/model"
)

/* 
tv支付授权查询 APIResponse
taobao.tvpay.auth.query

查询该用户在指定设备上是否有支付授权
*/
type TaobaoTvpayAuthQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTvpayAuthQueryResponse `json:"tvpay_auth_query_response,omitempty"` 
    TaobaoTvpayAuthQueryResponse
}

/* model for simplify = false
type TaobaoTvpayAuthQueryResponse struct {

    // Top返回对象
    
    Result  *struct {
        TopResultDo  *TopResultDo `json:"top_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoTvpayAuthQueryResponse struct {

    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty"`

}
