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
type AlibabaAliqinFlowWalletQueryChargeRequest struct {
    model.Params
    // 唯一流水号
    outRechargeId   string
    // 渠道id
    channelId   string
}

// 初始化AlibabaAliqinFlowWalletQueryChargeRequest对象
func NewAlibabaAliqinFlowWalletQueryChargeRequest() *AlibabaAliqinFlowWalletQueryChargeRequest{
    return &AlibabaAliqinFlowWalletQueryChargeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletQueryChargeRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.wallet.query.charge"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletQueryChargeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutRechargeId Setter
// 唯一流水号
func (r *AlibabaAliqinFlowWalletQueryChargeRequest) SetOutRechargeId(outRechargeId string) error {
    r.outRechargeId = outRechargeId
    r.Set("out_recharge_id", outRechargeId)
    return nil
}

// OutRechargeId Getter
func (r AlibabaAliqinFlowWalletQueryChargeRequest) GetOutRechargeId() string {
    return r.outRechargeId
}
// ChannelId Setter
// 渠道id
func (r *AlibabaAliqinFlowWalletQueryChargeRequest) SetChannelId(channelId string) error {
    r.channelId = channelId
    r.Set("channel_id", channelId)
    return nil
}

// ChannelId Getter
func (r AlibabaAliqinFlowWalletQueryChargeRequest) GetChannelId() string {
    return r.channelId
}
