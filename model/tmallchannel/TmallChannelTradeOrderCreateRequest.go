package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建渠道分销单 API请求
tmall.channel.trade.order.create

创建渠道分销单
*/
type TmallChannelTradeOrderCreateAPIRequest struct {
    model.Params
    // 入参
    _param0   *TopChannelPurchaseOrderCreateParam
}

// 初始化TmallChannelTradeOrderCreateAPIRequest对象
func NewTmallChannelTradeOrderCreateRequest() *TmallChannelTradeOrderCreateAPIRequest{
    return &TmallChannelTradeOrderCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeOrderCreateAPIRequest) GetApiMethodName() string {
    return "tmall.channel.trade.order.create"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeOrderCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 入参
func (r *TmallChannelTradeOrderCreateAPIRequest) SetParam0(_param0 *TopChannelPurchaseOrderCreateParam) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TmallChannelTradeOrderCreateAPIRequest) GetParam0() *TopChannelPurchaseOrderCreateParam {
    return r._param0
}
