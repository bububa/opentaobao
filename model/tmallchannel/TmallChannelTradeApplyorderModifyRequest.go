package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商修改申请单 API请求
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

// 初始化TmallChannelTradeApplyorderModifyRequest对象
func NewTmallChannelTradeApplyorderModifyRequest() *TmallChannelTradeApplyorderModifyRequest{
    return &TmallChannelTradeApplyorderModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeApplyorderModifyRequest) GetApiMethodName() string {
    return "tmall.channel.trade.applyorder.modify"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeApplyorderModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelPurchaseApplyOrderNo Setter
// 采购申请单号
func (r *TmallChannelTradeApplyorderModifyRequest) SetChannelPurchaseApplyOrderNo(channelPurchaseApplyOrderNo string) error {
    r.channelPurchaseApplyOrderNo = channelPurchaseApplyOrderNo
    r.Set("channel_purchase_apply_order_no", channelPurchaseApplyOrderNo)
    return nil
}

// ChannelPurchaseApplyOrderNo Getter
func (r TmallChannelTradeApplyorderModifyRequest) GetChannelPurchaseApplyOrderNo() string {
    return r.channelPurchaseApplyOrderNo
}
// ApplyOrderRelateItemModifyParamList Setter
// 修改关联的的宝贝信息
func (r *TmallChannelTradeApplyorderModifyRequest) SetApplyOrderRelateItemModifyParamList(applyOrderRelateItemModifyParamList []TopChannelApplyOrderRelateItemModifyParam) error {
    r.applyOrderRelateItemModifyParamList = applyOrderRelateItemModifyParamList
    r.Set("apply_order_relate_item_modify_param_list", applyOrderRelateItemModifyParamList)
    return nil
}

// ApplyOrderRelateItemModifyParamList Getter
func (r TmallChannelTradeApplyorderModifyRequest) GetApplyOrderRelateItemModifyParamList() []TopChannelApplyOrderRelateItemModifyParam {
    return r.applyOrderRelateItemModifyParamList
}
