package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商拒绝收货确认单 API请求
tmall.channel.trade.deliverorder.reject

供应商拒绝收货确认单
*/
type TmallChannelTradeDeliverorderRejectRequest struct {
    model.Params
    // 发货单号
    _mainDeliverOrderNo   int64
    // 拒绝原因
    _operateDesc   string
}

// 初始化TmallChannelTradeDeliverorderRejectRequest对象
func NewTmallChannelTradeDeliverorderRejectRequest() *TmallChannelTradeDeliverorderRejectRequest{
    return &TmallChannelTradeDeliverorderRejectRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeDeliverorderRejectRequest) GetApiMethodName() string {
    return "tmall.channel.trade.deliverorder.reject"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeDeliverorderRejectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainDeliverOrderNo Setter
// 发货单号
func (r *TmallChannelTradeDeliverorderRejectRequest) SetMainDeliverOrderNo(_mainDeliverOrderNo int64) error {
    r._mainDeliverOrderNo = _mainDeliverOrderNo
    r.Set("main_deliver_order_no", _mainDeliverOrderNo)
    return nil
}

// MainDeliverOrderNo Getter
func (r TmallChannelTradeDeliverorderRejectRequest) GetMainDeliverOrderNo() int64 {
    return r._mainDeliverOrderNo
}
// OperateDesc Setter
// 拒绝原因
func (r *TmallChannelTradeDeliverorderRejectRequest) SetOperateDesc(_operateDesc string) error {
    r._operateDesc = _operateDesc
    r.Set("operate_desc", _operateDesc)
    return nil
}

// OperateDesc Getter
func (r TmallChannelTradeDeliverorderRejectRequest) GetOperateDesc() string {
    return r._operateDesc
}