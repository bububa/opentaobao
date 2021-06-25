package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家关闭一笔交易 APIRequest
taobao.trade.close

关闭一笔订单，可以是主订单或子订单。当订单从创建到关闭时间小于10s的时候，会报“CLOSE_TRADE_TOO_FAST”错误。
*/
type TaobaoTradeCloseRequest struct {
    model.Params

    // 主订单或子订单编号。
    tid   int64 

    // 交易关闭原因。可以选择的理由有：1.未及时付款2、买家不想买了3、买家信息填写错误，重新拍4、恶意买家/同行捣乱5、缺货6、买家拍错了7、同城见面交易
    closeReason   string 

}

func NewTaobaoTradeCloseRequest() *TaobaoTradeCloseRequest{
    return &TaobaoTradeCloseRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTradeCloseRequest) GetApiMethodName() string {
    return "taobao.trade.close"
}

func (r TaobaoTradeCloseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTradeCloseRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoTradeCloseRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoTradeCloseRequest) SetCloseReason(closeReason string) error {
    r.closeReason = closeReason
    r.Set("close_reason", closeReason)
    return nil
}

func (r TaobaoTradeCloseRequest) GetCloseReason() string {
    return r.closeReason
}

