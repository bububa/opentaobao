package caipiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/caipiao"
)

/* 
卖家使用nick给买家送彩票 
taobao.caipiao.lottery.sendbynick

卖家使用nick给买家送彩票，可以指定彩种和注数。赠送成功，返回true; 以下几种情况情况， 返回false: 注数超过100注、卖家未签署支付宝代扣协议、卖家或者买家信息不存在等。
*/
func TaobaoCaipiaoLotterySendbynick(clt *core.SDKClient, req *caipiao.TaobaoCaipiaoLotterySendbynickRequest, session string) (*caipiao.TaobaoCaipiaoLotterySendbynickResponse, error) {
    var resp caipiao.TaobaoCaipiaoLotterySendbynickAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
