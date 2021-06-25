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
    Response *TaobaoTvpayAccessDataGetResponse `json:"taobao_tvpay_access_data_get_response,omitempty"`
}

type TaobaoTvpayAccessDataGetResponse struct {

    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty"`

}
