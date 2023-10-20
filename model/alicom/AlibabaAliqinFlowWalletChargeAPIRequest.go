package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinflowwalletchargeAPIRequest 流量直充 API请求
// alibaba.aliqin.flow.wallet.charge
//
// 流量直充
type AlibabaaliqinflowwalletchargeAPIRequest struct {
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

// NewAlibabaaliqinflowwalletchargeRequest 初始化AlibabaaliqinflowwalletchargeAPIRequest对象
func NewAlibabaaliqinflowwalletchargeRequest() *AlibabaaliqinflowwalletchargeAPIRequest {
	return &AlibabaaliqinflowwalletchargeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinflowwalletchargeAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.wallet.charge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinflowwalletchargeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinflowwalletchargeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhoneNum is PhoneNum Setter
// 充值号码
func (r *AlibabaaliqinflowwalletchargeAPIRequest) SetPhoneNum(_phoneNum string) error {
	r._phoneNum = _phoneNum
	r.Set("phone_num", _phoneNum)
	return nil
}

// GetPhoneNum PhoneNum Getter
func (r AlibabaaliqinflowwalletchargeAPIRequest) GetPhoneNum() string {
	return r._phoneNum
}

// SetReason is Reason Setter
// 原因
func (r *AlibabaaliqinflowwalletchargeAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r AlibabaaliqinflowwalletchargeAPIRequest) GetReason() string {
	return r._reason
}

// SetGradeId is GradeId Setter
// 档位id
func (r *AlibabaaliqinflowwalletchargeAPIRequest) SetGradeId(_gradeId string) error {
	r._gradeId = _gradeId
	r.Set("grade_id", _gradeId)
	return nil
}

// GetGradeId GradeId Getter
func (r AlibabaaliqinflowwalletchargeAPIRequest) GetGradeId() string {
	return r._gradeId
}

// SetOutRechargeId is OutRechargeId Setter
// 唯一流水号
func (r *AlibabaaliqinflowwalletchargeAPIRequest) SetOutRechargeId(_outRechargeId string) error {
	r._outRechargeId = _outRechargeId
	r.Set("out_recharge_id", _outRechargeId)
	return nil
}

// GetOutRechargeId OutRechargeId Getter
func (r AlibabaaliqinflowwalletchargeAPIRequest) GetOutRechargeId() string {
	return r._outRechargeId
}

// SetChannelId is ChannelId Setter
// 渠道id
func (r *AlibabaaliqinflowwalletchargeAPIRequest) SetChannelId(_channelId string) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r AlibabaaliqinflowwalletchargeAPIRequest) GetChannelId() string {
	return r._channelId
}
