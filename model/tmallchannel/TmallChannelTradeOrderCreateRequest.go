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
type TmallChannelTradeOrderCreateRequest struct {
    model.Params
    // 入参
    _param0   *TopChannelPurchaseOrderCreateParam
}

// 初始化TmallChannelTradeOrderCreateRequest对象
func NewTmallChannelTradeOrderCreateRequest() *TmallChannelTradeOrderCreateRequest{
    return &TmallChannelTradeOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeOrderCreateRequest) GetApiMethodName() string {
    return "tmall.channel.trade.order.create"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 入参
func (r *TmallChannelTradeOrderCreateRequest) SetParam0(_param0 *TopChannelPurchaseOrderCreateParam) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TmallChannelTradeOrderCreateRequest) GetParam0() *TopChannelPurchaseOrderCreateParam {
    return r._param0
}
