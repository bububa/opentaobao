package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
流量直充 API请求
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

// 初始化AlibabaAliqinFlowWalletChargeRequest对象
func NewAlibabaAliqinFlowWalletChargeRequest() *AlibabaAliqinFlowWalletChargeRequest{
    return &AlibabaAliqinFlowWalletChargeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletChargeRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.wallet.charge"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletChargeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PhoneNum Setter
// 充值号码
func (r *AlibabaAliqinFlowWalletChargeRequest) SetPhoneNum(phoneNum string) error {
    r.phoneNum = phoneNum
    r.Set("phone_num", phoneNum)
    return nil
}

// PhoneNum Getter
func (r AlibabaAliqinFlowWalletChargeRequest) GetPhoneNum() string {
    return r.phoneNum
}
// Reason Setter
// 原因
func (r *AlibabaAliqinFlowWalletChargeRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

// Reason Getter
func (r AlibabaAliqinFlowWalletChargeRequest) GetReason() string {
    return r.reason
}
// GradeId Setter
// 档位id
func (r *AlibabaAliqinFlowWalletChargeRequest) SetGradeId(gradeId string) error {
    r.gradeId = gradeId
    r.Set("grade_id", gradeId)
    return nil
}

// GradeId Getter
func (r AlibabaAliqinFlowWalletChargeRequest) GetGradeId() string {
    return r.gradeId
}
// OutRechargeId Setter
// 唯一流水号
func (r *AlibabaAliqinFlowWalletChargeRequest) SetOutRechargeId(outRechargeId string) error {
    r.outRechargeId = outRechargeId
    r.Set("out_recharge_id", outRechargeId)
    return nil
}

// OutRechargeId Getter
func (r AlibabaAliqinFlowWalletChargeRequest) GetOutRechargeId() string {
    return r.outRechargeId
}
// ChannelId Setter
// 渠道id
func (r *AlibabaAliqinFlowWalletChargeRequest) SetChannelId(channelId string) error {
    r.channelId = channelId
    r.Set("channel_id", channelId)
    return nil
}

// ChannelId Getter
func (r AlibabaAliqinFlowWalletChargeRequest) GetChannelId() string {
    return r.channelId
}
