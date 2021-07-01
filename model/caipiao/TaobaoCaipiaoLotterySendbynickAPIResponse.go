package caipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家使用nick给买家送彩票 API返回值 
taobao.caipiao.lottery.sendbynick

卖家使用nick给买家送彩票，可以指定彩种和注数。赠送成功，返回true; 以下几种情况情况， 返回false: 注数超过100注、卖家未签署支付宝代扣协议、卖家或者买家信息不存在等。
*/
type TaobaoCaipiaoLotterySendbynickAPIResponse struct {
    model.CommonResponse
    TaobaoCaipiaoLotterySendbynickAPIResponseModel
}

// 卖家使用nick给买家送彩票 成功返回结果
type TaobaoCaipiaoLotterySendbynickAPIResponseModel struct {
    XMLName xml.Name `xml:"caipiao_lottery_sendbynick_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 赠送是否成功，成功为true, 否则为false
    SendResult   bool `json:"send_result,omitempty" xml:"send_result,omitempty"`
}
