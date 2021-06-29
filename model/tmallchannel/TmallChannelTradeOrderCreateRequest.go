package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建渠道分销单 APIRequest
tmall.channel.trade.order.create

创建渠道分销单
*/
type TmallChannelTradeOrderCreateRequest struct {
    model.Params

    // 入参
    param0   *TopChannelPurchaseOrderCreateParam 

}

func NewTmallChannelTradeOrderCreateRequest() *TmallChannelTradeOrderCreateRequest{
    return &TmallChannelTradeOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallChannelTradeOrderCreateRequest) GetApiMethodName() string {
    return "tmall.channel.trade.order.create"
}

func (r TmallChannelTradeOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallChannelTradeOrderCreateRequest) SetParam0(param0 *TopChannelPurchaseOrderCreateParam) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TmallChannelTradeOrderCreateRequest) GetParam0() *TopChannelPurchaseOrderCreateParam {
    return r.param0
}

