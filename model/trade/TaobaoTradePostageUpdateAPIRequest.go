package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改交易邮费价格 API请求
taobao.trade.postage.update

修改订单邮费接口，通过传入订单编号和邮费价格，修改订单的邮费，返回修改时间modified,邮费post_fee,总费用total_fee。
<br/> <span style="color:red"> API取消加邮费功能通知：http://open.taobao.com/support/announcement_detail.htm?tid=24750</span>
*/
type TaobaoTradePostageUpdateAPIRequest struct {
    model.Params
    // 主订单编号
    _tid   int64
    // 邮费价格(邮费单位是元）
    _postFee   string
}

// 初始化TaobaoTradePostageUpdateAPIRequest对象
func NewTaobaoTradePostageUpdateRequest() *TaobaoTradePostageUpdateAPIRequest{
    return &TaobaoTradePostageUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradePostageUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.trade.postage.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradePostageUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 主订单编号
func (r *TaobaoTradePostageUpdateAPIRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoTradePostageUpdateAPIRequest) GetTid() int64 {
    return r._tid
}
// PostFee Setter
// 邮费价格(邮费单位是元）
func (r *TaobaoTradePostageUpdateAPIRequest) SetPostFee(_postFee string) error {
    r._postFee = _postFee
    r.Set("post_fee", _postFee)
    return nil
}

// PostFee Getter
func (r TaobaoTradePostageUpdateAPIRequest) GetPostFee() string {
    return r._postFee
}
