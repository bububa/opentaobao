package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过发货单单号获取发货单的详情 APIRequest
tmall.channel.trade.deliverorder.get

通过发货单单号获取发货单的详情
*/
type TmallChannelTradeDeliverorderGetRequest struct {
    model.Params

    // 发货单号
    mainDeliverOrderNo   int64 

    // 是否包含子发货单
    isIncludeSubOrder   bool 

}

func NewTmallChannelTradeDeliverorderGetRequest() *TmallChannelTradeDeliverorderGetRequest{
    return &TmallChannelTradeDeliverorderGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallChannelTradeDeliverorderGetRequest) GetApiMethodName() string {
    return "tmall.channel.trade.deliverorder.get"
}

func (r TmallChannelTradeDeliverorderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallChannelTradeDeliverorderGetRequest) SetMainDeliverOrderNo(mainDeliverOrderNo int64) error {
    r.mainDeliverOrderNo = mainDeliverOrderNo
    r.Set("main_deliver_order_no", mainDeliverOrderNo)
    return nil
}

func (r TmallChannelTradeDeliverorderGetRequest) GetMainDeliverOrderNo() int64 {
    return r.mainDeliverOrderNo
}

func (r *TmallChannelTradeDeliverorderGetRequest) SetIsIncludeSubOrder(isIncludeSubOrder bool) error {
    r.isIncludeSubOrder = isIncludeSubOrder
    r.Set("is_include_sub_order", isIncludeSubOrder)
    return nil
}

func (r TmallChannelTradeDeliverorderGetRequest) GetIsIncludeSubOrder() bool {
    return r.isIncludeSubOrder
}

