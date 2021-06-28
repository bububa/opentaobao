package tvpay

import (
    "github.com/bububa/opentaobao/model"
)

/* 
tv支付获取应用信息 APIResponse
taobao.tvpay.appinfo.get

tv支付获取应用信息
*/
type TaobaoTvpayAppinfoGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTvpayAppinfoGetResponse `json:"tvpay_appinfo_get_response,omitempty"` 
    TaobaoTvpayAppinfoGetResponse
}

/* model for simplify = false
type TaobaoTvpayAppinfoGetResponse struct {

    // Top返回对象
    
    Result  *struct {
        TopResultDo  *TopResultDo `json:"top_result_do,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoTvpayAppinfoGetResponse struct {

    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty"`

}
