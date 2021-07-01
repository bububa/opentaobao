package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过发货单单号获取发货单的详情 API请求
tmall.channel.trade.deliverorder.get

通过发货单单号获取发货单的详情
*/
type TmallChannelTradeDeliverorderGetAPIRequest struct {
    model.Params
    // 发货单号
    _mainDeliverOrderNo   int64
    // 是否包含子发货单
    _isIncludeSubOrder   bool
}

// 初始化TmallChannelTradeDeliverorderGetAPIRequest对象
func NewTmallChannelTradeDeliverorderGetRequest() *TmallChannelTradeDeliverorderGetAPIRequest{
    return &TmallChannelTradeDeliverorderGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeDeliverorderGetAPIRequest) GetApiMethodName() string {
    return "tmall.channel.trade.deliverorder.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeDeliverorderGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainDeliverOrderNo Setter
// 发货单号
func (r *TmallChannelTradeDeliverorderGetAPIRequest) SetMainDeliverOrderNo(_mainDeliverOrderNo int64) error {
    r._mainDeliverOrderNo = _mainDeliverOrderNo
    r.Set("main_deliver_order_no", _mainDeliverOrderNo)
    return nil
}

// MainDeliverOrderNo Getter
func (r TmallChannelTradeDeliverorderGetAPIRequest) GetMainDeliverOrderNo() int64 {
    return r._mainDeliverOrderNo
}
// IsIncludeSubOrder Setter
// 是否包含子发货单
func (r *TmallChannelTradeDeliverorderGetAPIRequest) SetIsIncludeSubOrder(_isIncludeSubOrder bool) error {
    r._isIncludeSubOrder = _isIncludeSubOrder
    r.Set("is_include_sub_order", _isIncludeSubOrder)
    return nil
}

// IsIncludeSubOrder Getter
func (r TmallChannelTradeDeliverorderGetAPIRequest) GetIsIncludeSubOrder() bool {
    return r._isIncludeSubOrder
}
