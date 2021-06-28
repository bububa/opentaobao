package tvpay

import (
    "github.com/bububa/opentaobao/model"
)

/* 
tv支付 APIResponse
taobao.tvpay.access.data.get

在匿名用户支付后尝试为其登陆绑定的淘宝账号
*/
type TaobaoTvpayAccessDataGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTvpayAccessDataGetResponse `json:"tvpay_access_data_get_response,omitempty"` 
    TaobaoTvpayAccessDataGetResponse
}

/* model for simplify = false
type TaobaoTvpayAccessDataGetResponse struct {

    // Top返回对象
    
    Result  *struct {
        TopResultDo  *TopResultDo `json:"top_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoTvpayAccessDataGetResponse struct {

    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty"`

}
