package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商修改申请单 APIRequest
tmall.channel.trade.applyorder.modify

上游供应商修改申请单, 目前只允许修改价格+件数且sku数量必须完全一致
*/
type TmallChannelTradeApplyorderModifyRequest struct {
    model.Params

    // 采购申请单号
    channelPurchaseApplyOrderNo   string 

    // 修改关联的的宝贝信息
    applyOrderRelateItemModifyParamList   []TopChannelApplyOrderRelateItemModifyParam 

}

func NewTmallChannelTradeApplyorderModifyRequest() *TmallChannelTradeApplyorderModifyRequest{
    return &TmallChannelTradeApplyorderModifyRequest{
        Params: model.NewParams(),
    }
}

func (r TmallChannelTradeApplyorderModifyRequest) GetApiMethodName() string {
    return "tmall.channel.trade.applyorder.modify"
}

func (r TmallChannelTradeApplyorderModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallChannelTradeApplyorderModifyRequest) SetChannelPurchaseApplyOrderNo(channelPurchaseApplyOrderNo string) error {
    r.channelPurchaseApplyOrderNo = channelPurchaseApplyOrderNo
    r.Set("channel_purchase_apply_order_no", channelPurchaseApplyOrderNo)
    return nil
}

func (r TmallChannelTradeApplyorderModifyRequest) GetChannelPurchaseApplyOrderNo() string {
    return r.channelPurchaseApplyOrderNo
}

func (r *TmallChannelTradeApplyorderModifyRequest) SetApplyOrderRelateItemModifyParamList(applyOrderRelateItemModifyParamList []TopChannelApplyOrderRelateItemModifyParam) error {
    r.applyOrderRelateItemModifyParamList = applyOrderRelateItemModifyParamList
    r.Set("apply_order_relate_item_modify_param_list", applyOrderRelateItemModifyParamList)
    return nil
}

func (r TmallChannelTradeApplyorderModifyRequest) GetApplyOrderRelateItemModifyParamList() []TopChannelApplyOrderRelateItemModifyParam {
    return r.applyOrderRelateItemModifyParamList
}

