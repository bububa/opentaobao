package caipiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
卖家使用nick给买家送彩票 APIResponse
taobao.caipiao.lottery.sendbynick

卖家使用nick给买家送彩票，可以指定彩种和注数。赠送成功，返回true; 以下几种情况情况， 返回false: 注数超过100注、卖家未签署支付宝代扣协议、卖家或者买家信息不存在等。
*/
type TaobaoCaipiaoLotterySendbynickAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCaipiaoLotterySendbynickResponse `json:"taobao_caipiao_lottery_sendbynick_response,omitempty"`
}

type TaobaoCaipiaoLotterySendbynickResponse struct {

    // 赠送是否成功，成功为true, 否则为false
    SendResult   bool `json:"send_result,omitempty"`

}
