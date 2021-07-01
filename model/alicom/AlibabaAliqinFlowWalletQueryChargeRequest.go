package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询流量充值状态 API请求
alibaba.aliqin.flow.wallet.query.charge

查询流量充值状态
*/
type AlibabaAliqinFlowWalletQueryChargeAPIRequest struct {
    model.Params
    // 唯一流水号
    _outRechargeId   string
    // 渠道id
    _channelId   string
}

// 初始化AlibabaAliqinFlowWalletQueryChargeAPIRequest对象
func NewAlibabaAliqinFlowWalletQueryChargeRequest() *AlibabaAliqinFlowWalletQueryChargeAPIRequest{
    return &AlibabaAliqinFlowWalletQueryChargeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletQueryChargeAPIRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.wallet.query.charge"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletQueryChargeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutRechargeId Setter
// 唯一流水号
func (r *AlibabaAliqinFlowWalletQueryChargeAPIRequest) SetOutRechargeId(_outRechargeId string) error {
    r._outRechargeId = _outRechargeId
    r.Set("out_recharge_id", _outRechargeId)
    return nil
}

// OutRechargeId Getter
func (r AlibabaAliqinFlowWalletQueryChargeAPIRequest) GetOutRechargeId() string {
    return r._outRechargeId
}
// ChannelId Setter
// 渠道id
func (r *AlibabaAliqinFlowWalletQueryChargeAPIRequest) SetChannelId(_channelId string) error {
    r._channelId = _channelId
    r.Set("channel_id", _channelId)
    return nil
}

// ChannelId Getter
func (r AlibabaAliqinFlowWalletQueryChargeAPIRequest) GetChannelId() string {
    return r._channelId
}
