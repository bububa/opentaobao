package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家关闭一笔交易 API请求
taobao.trade.close

关闭一笔订单，可以是主订单或子订单。当订单从创建到关闭时间小于10s的时候，会报“CLOSE_TRADE_TOO_FAST”错误。
*/
type TaobaoTradeCloseAPIRequest struct {
    model.Params
    // 主订单或子订单编号。
    _tid   int64
    // 交易关闭原因。可以选择的理由有：1.未及时付款2、买家不想买了3、买家信息填写错误，重新拍4、恶意买家/同行捣乱5、缺货6、买家拍错了7、同城见面交易
    _closeReason   string
}

// 初始化TaobaoTradeCloseAPIRequest对象
func NewTaobaoTradeCloseRequest() *TaobaoTradeCloseAPIRequest{
    return &TaobaoTradeCloseAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeCloseAPIRequest) GetApiMethodName() string {
    return "taobao.trade.close"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeCloseAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 主订单或子订单编号。
func (r *TaobaoTradeCloseAPIRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoTradeCloseAPIRequest) GetTid() int64 {
    return r._tid
}
// CloseReason Setter
// 交易关闭原因。可以选择的理由有：1.未及时付款2、买家不想买了3、买家信息填写错误，重新拍4、恶意买家/同行捣乱5、缺货6、买家拍错了7、同城见面交易
func (r *TaobaoTradeCloseAPIRequest) SetCloseReason(_closeReason string) error {
    r._closeReason = _closeReason
    r.Set("close_reason", _closeReason)
    return nil
}

// CloseReason Getter
func (r TaobaoTradeCloseAPIRequest) GetCloseReason() string {
    return r._closeReason
}
