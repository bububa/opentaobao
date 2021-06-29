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
    _phoneNum   string
    // 原因
    _reason   string
    // 档位id
    _gradeId   string
    // 唯一流水号
    _outRechargeId   string
    // 渠道id（运营分配）
    _channelId   string
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
func (r *AlibabaAliqinFlowWalletChargeRuleRequest) SetPhoneNum(_phoneNum string) error {
    r._phoneNum = _phoneNum
    r.Set("phone_num", _phoneNum)
    return nil
}

// PhoneNum Getter
func (r AlibabaAliqinFlowWalletChargeRuleRequest) GetPhoneNum() string {
    return r._phoneNum
}
// Reason Setter
// 原因
func (r *AlibabaAliqinFlowWalletChargeRuleRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r AlibabaAliqinFlowWalletChargeRuleRequest) GetReason() string {
    return r._reason
}
// GradeId Setter
// 档位id
func (r *AlibabaAliqinFlowWalletChargeRuleRequest) SetGradeId(_gradeId string) error {
    r._gradeId = _gradeId
    r.Set("grade_id", _gradeId)
    return nil
}

// GradeId Getter
func (r AlibabaAliqinFlowWalletChargeRuleRequest) GetGradeId() string {
    return r._gradeId
}
// OutRechargeId Setter
// 唯一流水号
func (r *AlibabaAliqinFlowWalletChargeRuleRequest) SetOutRechargeId(_outRechargeId string) error {
    r._outRechargeId = _outRechargeId
    r.Set("out_recharge_id", _outRechargeId)
    return nil
}

// OutRechargeId Getter
func (r AlibabaAliqinFlowWalletChargeRuleRequest) GetOutRechargeId() string {
    return r._outRechargeId
}
// ChannelId Setter
// 渠道id（运营分配）
func (r *AlibabaAliqinFlowWalletChargeRuleRequest) SetChannelId(_channelId string) error {
    r._channelId = _channelId
    r.Set("channel_id", _channelId)
    return nil
}

// ChannelId Getter
func (r AlibabaAliqinFlowWalletChargeRuleRequest) GetChannelId() string {
    return r._channelId
}
