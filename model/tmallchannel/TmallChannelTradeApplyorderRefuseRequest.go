package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商审核拒绝采购申请单 APIRequest
tmall.channel.trade.applyorder.refuse

供应商审核拒绝采购申请单
*/
type TmallChannelTradeApplyorderRefuseRequest struct {
    model.Params

    // 操作描述
    operateDesc   string 

    // 采购申请单号
    channelPurchaseApplyOrderNo   string 

}

func NewTmallChannelTradeApplyorderRefuseRequest() *TmallChannelTradeApplyorderRefuseRequest{
    return &TmallChannelTradeApplyorderRefuseRequest{
        Params: model.NewParams(),
    }
}

func (r TmallChannelTradeApplyorderRefuseRequest) GetApiMethodName() string {
    return "tmall.channel.trade.applyorder.refuse"
}

func (r TmallChannelTradeApplyorderRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallChannelTradeApplyorderRefuseRequest) SetOperateDesc(operateDesc string) error {
    r.operateDesc = operateDesc
    r.Set("operate_desc", operateDesc)
    return nil
}

func (r TmallChannelTradeApplyorderRefuseRequest) GetOperateDesc() string {
    return r.operateDesc
}

func (r *TmallChannelTradeApplyorderRefuseRequest) SetChannelPurchaseApplyOrderNo(channelPurchaseApplyOrderNo string) error {
    r.channelPurchaseApplyOrderNo = channelPurchaseApplyOrderNo
    r.Set("channel_purchase_apply_order_no", channelPurchaseApplyOrderNo)
    return nil
}

func (r TmallChannelTradeApplyorderRefuseRequest) GetChannelPurchaseApplyOrderNo() string {
    return r.channelPurchaseApplyOrderNo
}

