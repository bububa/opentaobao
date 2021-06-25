package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询流量充值状态 APIRequest
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

func NewAlibabaAliqinFlowWalletQueryChargeRequest() *AlibabaAliqinFlowWalletQueryChargeRequest{
    return &AlibabaAliqinFlowWalletQueryChargeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFlowWalletQueryChargeRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.wallet.query.charge"
}

func (r AlibabaAliqinFlowWalletQueryChargeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFlowWalletQueryChargeRequest) SetOutRechargeId(outRechargeId string) error {
    r.outRechargeId = outRechargeId
    r.Set("out_recharge_id", outRechargeId)
    return nil
}

func (r AlibabaAliqinFlowWalletQueryChargeRequest) GetOutRechargeId() string {
    return r.outRechargeId
}

func (r *AlibabaAliqinFlowWalletQueryChargeRequest) SetChannelId(channelId string) error {
    r.channelId = channelId
    r.Set("channel_id", channelId)
    return nil
}

func (r AlibabaAliqinFlowWalletQueryChargeRequest) GetChannelId() string {
    return r.channelId
}

