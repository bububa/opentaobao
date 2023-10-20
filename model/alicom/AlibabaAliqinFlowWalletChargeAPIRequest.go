package alicom

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowWalletChargeAPIRequest 流量直充 API请求
// alibaba.aliqin.flow.wallet.charge
//
// 流量直充
type AlibabaAliqinFlowWalletChargeAPIRequest struct {
	model.Params
	// 充值号码
	_phoneNum string
	// 原因
	_reason string
	// 档位id
	_gradeId string
	// 唯一流水号
	_outRechargeId string
	// 渠道id
	_channelId string
}

// NewAlibabaAliqinFlowWalletChargeRequest 初始化AlibabaAliqinFlowWalletChargeAPIRequest对象
func NewAlibabaAliqinFlowWalletChargeRequest() *AlibabaAliqinFlowWalletChargeAPIRequest {
	return &AlibabaAliqinFlowWalletChargeAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFlowWalletChargeAPIRequest) Reset() {
	r._phoneNum = ""
	r._reason = ""
	r._gradeId = ""
	r._outRechargeId = ""
	r._channelId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletChargeAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.wallet.charge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletChargeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFlowWalletChargeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhoneNum is PhoneNum Setter
// 充值号码
func (r *AlibabaAliqinFlowWalletChargeAPIRequest) SetPhoneNum(_phoneNum string) error {
	r._phoneNum = _phoneNum
	r.Set("phone_num", _phoneNum)
	return nil
}

// GetPhoneNum PhoneNum Getter
func (r AlibabaAliqinFlowWalletChargeAPIRequest) GetPhoneNum() string {
	return r._phoneNum
}

// SetReason is Reason Setter
// 原因
func (r *AlibabaAliqinFlowWalletChargeAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r AlibabaAliqinFlowWalletChargeAPIRequest) GetReason() string {
	return r._reason
}

// SetGradeId is GradeId Setter
// 档位id
func (r *AlibabaAliqinFlowWalletChargeAPIRequest) SetGradeId(_gradeId string) error {
	r._gradeId = _gradeId
	r.Set("grade_id", _gradeId)
	return nil
}

// GetGradeId GradeId Getter
func (r AlibabaAliqinFlowWalletChargeAPIRequest) GetGradeId() string {
	return r._gradeId
}

// SetOutRechargeId is OutRechargeId Setter
// 唯一流水号
func (r *AlibabaAliqinFlowWalletChargeAPIRequest) SetOutRechargeId(_outRechargeId string) error {
	r._outRechargeId = _outRechargeId
	r.Set("out_recharge_id", _outRechargeId)
	return nil
}

// GetOutRechargeId OutRechargeId Getter
func (r AlibabaAliqinFlowWalletChargeAPIRequest) GetOutRechargeId() string {
	return r._outRechargeId
}

// SetChannelId is ChannelId Setter
// 渠道id
func (r *AlibabaAliqinFlowWalletChargeAPIRequest) SetChannelId(_channelId string) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r AlibabaAliqinFlowWalletChargeAPIRequest) GetChannelId() string {
	return r._channelId
}

var poolAlibabaAliqinFlowWalletChargeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFlowWalletChargeRequest()
	},
}

// GetAlibabaAliqinFlowWalletChargeRequest 从 sync.Pool 获取 AlibabaAliqinFlowWalletChargeAPIRequest
func GetAlibabaAliqinFlowWalletChargeAPIRequest() *AlibabaAliqinFlowWalletChargeAPIRequest {
	return poolAlibabaAliqinFlowWalletChargeAPIRequest.Get().(*AlibabaAliqinFlowWalletChargeAPIRequest)
}

// ReleaseAlibabaAliqinFlowWalletChargeAPIRequest 将 AlibabaAliqinFlowWalletChargeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFlowWalletChargeAPIRequest(v *AlibabaAliqinFlowWalletChargeAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFlowWalletChargeAPIRequest.Put(v)
}
