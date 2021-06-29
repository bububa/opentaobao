package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商审核拒绝采购申请单 API请求
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

// 初始化TmallChannelTradeApplyorderRefuseRequest对象
func NewTmallChannelTradeApplyorderRefuseRequest() *TmallChannelTradeApplyorderRefuseRequest{
    return &TmallChannelTradeApplyorderRefuseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeApplyorderRefuseRequest) GetApiMethodName() string {
    return "tmall.channel.trade.applyorder.refuse"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeApplyorderRefuseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OperateDesc Setter
// 操作描述
func (r *TmallChannelTradeApplyorderRefuseRequest) SetOperateDesc(operateDesc string) error {
    r.operateDesc = operateDesc
    r.Set("operate_desc", operateDesc)
    return nil
}

// OperateDesc Getter
func (r TmallChannelTradeApplyorderRefuseRequest) GetOperateDesc() string {
    return r.operateDesc
}
// ChannelPurchaseApplyOrderNo Setter
// 采购申请单号
func (r *TmallChannelTradeApplyorderRefuseRequest) SetChannelPurchaseApplyOrderNo(channelPurchaseApplyOrderNo string) error {
    r.channelPurchaseApplyOrderNo = channelPurchaseApplyOrderNo
    r.Set("channel_purchase_apply_order_no", channelPurchaseApplyOrderNo)
    return nil
}

// ChannelPurchaseApplyOrderNo Getter
func (r TmallChannelTradeApplyorderRefuseRequest) GetChannelPurchaseApplyOrderNo() string {
    return r.channelPurchaseApplyOrderNo
}
