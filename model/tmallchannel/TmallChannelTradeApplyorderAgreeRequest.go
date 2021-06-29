package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商审核同意采购申请单 APIRequest
tmall.channel.trade.applyorder.agree

供应商审核同意采购申请单
*/
type TmallChannelTradeApplyorderAgreeRequest struct {
    model.Params

    // 操作描述
    operateDesc   string 

    // 采购申请单号
    channelPurchaseApplyOrderNo   string 

}

func NewTmallChannelTradeApplyorderAgreeRequest() *TmallChannelTradeApplyorderAgreeRequest{
    return &TmallChannelTradeApplyorderAgreeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallChannelTradeApplyorderAgreeRequest) GetApiMethodName() string {
    return "tmall.channel.trade.applyorder.agree"
}

func (r TmallChannelTradeApplyorderAgreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallChannelTradeApplyorderAgreeRequest) SetOperateDesc(operateDesc string) error {
    r.operateDesc = operateDesc
    r.Set("operate_desc", operateDesc)
    return nil
}

func (r TmallChannelTradeApplyorderAgreeRequest) GetOperateDesc() string {
    return r.operateDesc
}

func (r *TmallChannelTradeApplyorderAgreeRequest) SetChannelPurchaseApplyOrderNo(channelPurchaseApplyOrderNo string) error {
    r.channelPurchaseApplyOrderNo = channelPurchaseApplyOrderNo
    r.Set("channel_purchase_apply_order_no", channelPurchaseApplyOrderNo)
    return nil
}

func (r TmallChannelTradeApplyorderAgreeRequest) GetChannelPurchaseApplyOrderNo() string {
    return r.channelPurchaseApplyOrderNo
}

