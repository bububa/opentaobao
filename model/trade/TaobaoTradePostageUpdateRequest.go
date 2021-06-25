package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改交易邮费价格 APIRequest
taobao.trade.postage.update

修改订单邮费接口，通过传入订单编号和邮费价格，修改订单的邮费，返回修改时间modified,邮费post_fee,总费用total_fee。
<br/> <span style="color:red"> API取消加邮费功能通知：http://open.taobao.com/support/announcement_detail.htm?tid=24750</span>
*/
type TaobaoTradePostageUpdateRequest struct {
    model.Params

    // 主订单编号
    tid   int64 

    // 邮费价格(邮费单位是元）
    postFee   string 

}

func NewTaobaoTradePostageUpdateRequest() *TaobaoTradePostageUpdateRequest{
    return &TaobaoTradePostageUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradePostageUpdateRequest) GetApiMethodName() string {
    return "taobao.trade.postage.update"
}

func (r TaobaoTradePostageUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradePostageUpdateRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoTradePostageUpdateRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoTradePostageUpdateRequest) SetPostFee(postFee string) error {
    r.postFee = postFee
    r.Set("post_fee", postFee)
    return nil
}

func (r TaobaoTradePostageUpdateRequest) GetPostFee() string {
    return r.postFee
}

