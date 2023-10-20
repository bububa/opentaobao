package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinflowwalletchargeruleAPIRequest 流量钱包直充（根据号码归属地省份路由） API请求
// alibaba.aliqin.flow.wallet.charge.rule
//
// 流量钱包直充（根据号码归属地省份路由）
type AlibabaaliqinflowwalletchargeruleAPIRequest struct {
	model.Params
	// 号码
	_phoneNum string
	// 原因
	_reason string
	// 档位id
	_gradeId string
	// 唯一流水号
	_outRechargeId string
	// 渠道id（运营分配）
	_channelId string
}

// NewAlibabaaliqinflowwalletchargeruleRequest 初始化AlibabaaliqinflowwalletchargeruleAPIRequest对象
func NewAlibabaaliqinflowwalletchargeruleRequest() *AlibabaaliqinflowwalletchargeruleAPIRequest {
	return &AlibabaaliqinflowwalletchargeruleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinflowwalletchargeruleAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.wallet.charge.rule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinflowwalletchargeruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinflowwalletchargeruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhoneNum is PhoneNum Setter
// 号码
func (r *AlibabaaliqinflowwalletchargeruleAPIRequest) SetPhoneNum(_phoneNum string) error {
	r._phoneNum = _phoneNum
	r.Set("phone_num", _phoneNum)
	return nil
}

// GetPhoneNum PhoneNum Getter
func (r AlibabaaliqinflowwalletchargeruleAPIRequest) GetPhoneNum() string {
	return r._phoneNum
}

// SetReason is Reason Setter
// 原因
func (r *AlibabaaliqinflowwalletchargeruleAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r AlibabaaliqinflowwalletchargeruleAPIRequest) GetReason() string {
	return r._reason
}

// SetGradeId is GradeId Setter
// 档位id
func (r *AlibabaaliqinflowwalletchargeruleAPIRequest) SetGradeId(_gradeId string) error {
	r._gradeId = _gradeId
	r.Set("grade_id", _gradeId)
	return nil
}

// GetGradeId GradeId Getter
func (r AlibabaaliqinflowwalletchargeruleAPIRequest) GetGradeId() string {
	return r._gradeId
}

// SetOutRechargeId is OutRechargeId Setter
// 唯一流水号
func (r *AlibabaaliqinflowwalletchargeruleAPIRequest) SetOutRechargeId(_outRechargeId string) error {
	r._outRechargeId = _outRechargeId
	r.Set("out_recharge_id", _outRechargeId)
	return nil
}

// GetOutRechargeId OutRechargeId Getter
func (r AlibabaaliqinflowwalletchargeruleAPIRequest) GetOutRechargeId() string {
	return r._outRechargeId
}

// SetChannelId is ChannelId Setter
// 渠道id（运营分配）
func (r *AlibabaaliqinflowwalletchargeruleAPIRequest) SetChannelId(_channelId string) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r AlibabaaliqinflowwalletchargeruleAPIRequest) GetChannelId() string {
	return r._channelId
}
