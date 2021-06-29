package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询采购申请单详情 APIRequest
tmall.channel.trade.applyorder.get

通过采购申请单ID获取单据详情
*/
type TmallChannelTradeApplyorderGetRequest struct {
    model.Params

    // 采购申请单单号
    channelPurchaseApplyOrderNo   string 

}

func NewTmallChannelTradeApplyorderGetRequest() *TmallChannelTradeApplyorderGetRequest{
    return &TmallChannelTradeApplyorderGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallChannelTradeApplyorderGetRequest) GetApiMethodName() string {
    return "tmall.channel.trade.applyorder.get"
}

func (r TmallChannelTradeApplyorderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallChannelTradeApplyorderGetRequest) SetChannelPurchaseApplyOrderNo(channelPurchaseApplyOrderNo string) error {
    r.channelPurchaseApplyOrderNo = channelPurchaseApplyOrderNo
    r.Set("channel_purchase_apply_order_no", channelPurchaseApplyOrderNo)
    return nil
}

func (r TmallChannelTradeApplyorderGetRequest) GetChannelPurchaseApplyOrderNo() string {
    return r.channelPurchaseApplyOrderNo
}

