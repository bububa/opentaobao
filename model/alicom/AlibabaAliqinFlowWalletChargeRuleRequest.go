package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
流量钱包直充（根据号码归属地省份路由） API请求
alibaba.aliqin.flow.wallet.charge.rule

流量钱包直充（根据号码归属地省份路由）
*/
type AlibabaAliqinFlowWalletChargeRuleRequest struct {
    model.Params
    // 号码
    phoneNum   string
    // 原因
    reason   string
    // 档位id
    gradeId   string
    // 唯一流水号
    outRechargeId   string
    // 渠道id（运营分配）
    channelId   string
}

// 初始化AlibabaAliqinFlowWalletChargeRuleRequest对象
func NewAlibabaAliqinFlowWalletChargeRuleRequest() *AlibabaAliqinFlowWalletChargeRuleRequest{
    return &AlibabaAliqinFlowWalletChargeRuleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletChargeRuleRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.wallet.charge.rule"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletChargeRuleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PhoneNum Setter
// 号码
func (r *AlibabaAliqinFlowWalletChargeRuleRequest) SetPhoneNum(phoneNum string) error {
    r.phoneNum = phoneNum
    r.Set("phone_num", phoneNum)
    return nil
}

// PhoneNum Getter
func (r AlibabaAliqinFlowWalletChargeRuleRequest) GetPhoneNum() string {
    return r.phoneNum
}
// Reason Setter
// 原因
func (r *AlibabaAliqinFlowWalletChargeRuleRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

// Reason Getter
func (r AlibabaAliqinFlowWalletChargeRuleRequest) GetReason() string {
    return r.reason
}
// GradeId Setter
// 档位id
func (r *AlibabaAliqinFlowWalletChargeRuleRequest) SetGradeId(gradeId string) error {
    r.gradeId = gradeId
    r.Set("grade_id", gradeId)
    return nil
}

// GradeId Getter
func (r AlibabaAliqinFlowWalletChargeRuleRequest) GetGradeId() string {
    return r.gradeId
}
// OutRechargeId Setter
// 唯一流水号
func (r *AlibabaAliqinFlowWalletChargeRuleRequest) SetOutRechargeId(outRechargeId string) error {
    r.outRechargeId = outRechargeId
    r.Set("out_recharge_id", outRechargeId)
    return nil
}

// OutRechargeId Getter
func (r AlibabaAliqinFlowWalletChargeRuleRequest) GetOutRechargeId() string {
    return r.outRechargeId
}
// ChannelId Setter
// 渠道id（运营分配）
func (r *AlibabaAliqinFlowWalletChargeRuleRequest) SetChannelId(channelId string) error {
    r.channelId = channelId
    r.Set("channel_id", channelId)
    return nil
}

// ChannelId Getter
func (r AlibabaAliqinFlowWalletChargeRuleRequest) GetChannelId() string {
    return r.channelId
}
