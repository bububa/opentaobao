package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商审核通过发货确认 API请求
tmall.channel.trade.deliverorder.agree

供应商通过收货确认单
*/
type TmallChannelTradeDeliverorderAgreeRequest struct {
    model.Params
    // 发货单号
    _mainDeliverOrderNo   int64
    // 同意理由
    _operateDesc   string
}

// 初始化TmallChannelTradeDeliverorderAgreeRequest对象
func NewTmallChannelTradeDeliverorderAgreeRequest() *TmallChannelTradeDeliverorderAgreeRequest{
    return &TmallChannelTradeDeliverorderAgreeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeDeliverorderAgreeRequest) GetApiMethodName() string {
    return "tmall.channel.trade.deliverorder.agree"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeDeliverorderAgreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainDeliverOrderNo Setter
// 发货单号
func (r *TmallChannelTradeDeliverorderAgreeRequest) SetMainDeliverOrderNo(_mainDeliverOrderNo int64) error {
    r._mainDeliverOrderNo = _mainDeliverOrderNo
    r.Set("main_deliver_order_no", _mainDeliverOrderNo)
    return nil
}

// MainDeliverOrderNo Getter
func (r TmallChannelTradeDeliverorderAgreeRequest) GetMainDeliverOrderNo() int64 {
    return r._mainDeliverOrderNo
}
// OperateDesc Setter
// 同意理由
func (r *TmallChannelTradeDeliverorderAgreeRequest) SetOperateDesc(_operateDesc string) error {
    r._operateDesc = _operateDesc
    r.Set("operate_desc", _operateDesc)
    return nil
}

// OperateDesc Getter
func (r TmallChannelTradeDeliverorderAgreeRequest) GetOperateDesc() string {
    return r._operateDesc
}
