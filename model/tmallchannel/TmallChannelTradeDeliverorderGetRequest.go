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
type TmallChannelTradeDeliverorderGetRequest struct {
    model.Params
    // 发货单号
    _mainDeliverOrderNo   int64
    // 是否包含子发货单
    _isIncludeSubOrder   bool
}

// 初始化TmallChannelTradeDeliverorderGetRequest对象
func NewTmallChannelTradeDeliverorderGetRequest() *TmallChannelTradeDeliverorderGetRequest{
    return &TmallChannelTradeDeliverorderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeDeliverorderGetRequest) GetApiMethodName() string {
    return "tmall.channel.trade.deliverorder.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeDeliverorderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainDeliverOrderNo Setter
// 发货单号
func (r *TmallChannelTradeDeliverorderGetRequest) SetMainDeliverOrderNo(_mainDeliverOrderNo int64) error {
    r._mainDeliverOrderNo = _mainDeliverOrderNo
    r.Set("main_deliver_order_no", _mainDeliverOrderNo)
    return nil
}

// MainDeliverOrderNo Getter
func (r TmallChannelTradeDeliverorderGetRequest) GetMainDeliverOrderNo() int64 {
    return r._mainDeliverOrderNo
}
// IsIncludeSubOrder Setter
// 是否包含子发货单
func (r *TmallChannelTradeDeliverorderGetRequest) SetIsIncludeSubOrder(_isIncludeSubOrder bool) error {
    r._isIncludeSubOrder = _isIncludeSubOrder
    r.Set("is_include_sub_order", _isIncludeSubOrder)
    return nil
}

// IsIncludeSubOrder Getter
func (r TmallChannelTradeDeliverorderGetRequest) GetIsIncludeSubOrder() bool {
    return r._isIncludeSubOrder
}
