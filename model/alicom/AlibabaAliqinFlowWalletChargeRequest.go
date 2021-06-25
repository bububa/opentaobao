package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
流量直充 APIRequest
alibaba.aliqin.flow.wallet.charge

流量直充
*/
type AlibabaAliqinFlowWalletChargeRequest struct {
    model.Params

    // 充值号码
    phoneNum   string 

    // 原因
    reason   string 

    // 档位id
    gradeId   string 

    // 唯一流水号
    outRechargeId   string 

    // 渠道id
    channelId   string 

}

func NewAlibabaAliqinFlowWalletChargeRequest() *AlibabaAliqinFlowWalletChargeRequest{
    return &AlibabaAliqinFlowWalletChargeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFlowWalletChargeRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.wallet.charge"
}

func (r AlibabaAliqinFlowWalletChargeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFlowWalletChargeRequest) SetPhoneNum(phoneNum string) error {
    r.phoneNum = phoneNum
    r.Set("phone_num", phoneNum)
    return nil
}

func (r AlibabaAliqinFlowWalletChargeRequest) GetPhoneNum() string {
    return r.phoneNum
}

func (r *AlibabaAliqinFlowWalletChargeRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r AlibabaAliqinFlowWalletChargeRequest) GetReason() string {
    return r.reason
}

func (r *AlibabaAliqinFlowWalletChargeRequest) SetGradeId(gradeId string) error {
    r.gradeId = gradeId
    r.Set("grade_id", gradeId)
    return nil
}

func (r AlibabaAliqinFlowWalletChargeRequest) GetGradeId() string {
    return r.gradeId
}

func (r *AlibabaAliqinFlowWalletChargeRequest) SetOutRechargeId(outRechargeId string) error {
    r.outRechargeId = outRechargeId
    r.Set("out_recharge_id", outRechargeId)
    return nil
}

func (r AlibabaAliqinFlowWalletChargeRequest) GetOutRechargeId() string {
    return r.outRechargeId
}

func (r *AlibabaAliqinFlowWalletChargeRequest) SetChannelId(channelId string) error {
    r.channelId = channelId
    r.Set("channel_id", channelId)
    return nil
}

func (r AlibabaAliqinFlowWalletChargeRequest) GetChannelId() string {
    return r.channelId
}

