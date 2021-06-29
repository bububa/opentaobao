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
    _phoneNum   string
    // 原因
    _reason   string
    // 档位id
    _gradeId   string
    // 唯一流水号
    _outRechargeId   string
    // 渠道id
    _channelId   string
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
func (r *AlibabaAliqinFlowWalletChargeRequest) SetPhoneNum(_phoneNum string) error {
    r._phoneNum = _phoneNum
    r.Set("phone_num", _phoneNum)
    return nil
}

// PhoneNum Getter
func (r AlibabaAliqinFlowWalletChargeRequest) GetPhoneNum() string {
    return r._phoneNum
}
// Reason Setter
// 原因
func (r *AlibabaAliqinFlowWalletChargeRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r AlibabaAliqinFlowWalletChargeRequest) GetReason() string {
    return r._reason
}
// GradeId Setter
// 档位id
func (r *AlibabaAliqinFlowWalletChargeRequest) SetGradeId(_gradeId string) error {
    r._gradeId = _gradeId
    r.Set("grade_id", _gradeId)
    return nil
}

// GradeId Getter
func (r AlibabaAliqinFlowWalletChargeRequest) GetGradeId() string {
    return r._gradeId
}
// OutRechargeId Setter
// 唯一流水号
func (r *AlibabaAliqinFlowWalletChargeRequest) SetOutRechargeId(_outRechargeId string) error {
    r._outRechargeId = _outRechargeId
    r.Set("out_recharge_id", _outRechargeId)
    return nil
}

// OutRechargeId Getter
func (r AlibabaAliqinFlowWalletChargeRequest) GetOutRechargeId() string {
    return r._outRechargeId
}
// ChannelId Setter
// 渠道id
func (r *AlibabaAliqinFlowWalletChargeRequest) SetChannelId(_channelId string) error {
    r._channelId = _channelId
    r.Set("channel_id", _channelId)
    return nil
}

// ChannelId Getter
func (r AlibabaAliqinFlowWalletChargeRequest) GetChannelId() string {
    return r._channelId
}
