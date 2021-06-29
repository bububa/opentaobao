package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商审核同意采购申请单 API请求
tmall.channel.trade.applyorder.agree

供应商审核同意采购申请单
*/
type TmallChannelTradeApplyorderAgreeRequest struct {
    model.Params
    // 操作描述
    _operateDesc   string
    // 采购申请单号
    _channelPurchaseApplyOrderNo   string
}

// 初始化TmallChannelTradeApplyorderAgreeRequest对象
func NewTmallChannelTradeApplyorderAgreeRequest() *TmallChannelTradeApplyorderAgreeRequest{
    return &TmallChannelTradeApplyorderAgreeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeApplyorderAgreeRequest) GetApiMethodName() string {
    return "tmall.channel.trade.applyorder.agree"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeApplyorderAgreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OperateDesc Setter
// 操作描述
func (r *TmallChannelTradeApplyorderAgreeRequest) SetOperateDesc(_operateDesc string) error {
    r._operateDesc = _operateDesc
    r.Set("operate_desc", _operateDesc)
    return nil
}

// OperateDesc Getter
func (r TmallChannelTradeApplyorderAgreeRequest) GetOperateDesc() string {
    return r._operateDesc
}
// ChannelPurchaseApplyOrderNo Setter
// 采购申请单号
func (r *TmallChannelTradeApplyorderAgreeRequest) SetChannelPurchaseApplyOrderNo(_channelPurchaseApplyOrderNo string) error {
    r._channelPurchaseApplyOrderNo = _channelPurchaseApplyOrderNo
    r.Set("channel_purchase_apply_order_no", _channelPurchaseApplyOrderNo)
    return nil
}

// ChannelPurchaseApplyOrderNo Getter
func (r TmallChannelTradeApplyorderAgreeRequest) GetChannelPurchaseApplyOrderNo() string {
    return r._channelPurchaseApplyOrderNo
}
