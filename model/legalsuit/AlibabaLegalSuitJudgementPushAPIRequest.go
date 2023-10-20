package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalsuitjudgementpushAPIRequest 推送裁判登记信息给集团法务系统 API请求
// alibaba.legal.suit.judgement.push
//
// isv推送裁判登记信息给集团法务系统
type AlibabalegalsuitjudgementpushAPIRequest struct {
	model.Params
	// 裁判登记信息
	_refereeRegistrationModel *RefereeRegistrationModel
}

// NewAlibabalegalsuitjudgementpushRequest 初始化AlibabalegalsuitjudgementpushAPIRequest对象
func NewAlibabalegalsuitjudgementpushRequest() *AlibabalegalsuitjudgementpushAPIRequest {
	return &AlibabalegalsuitjudgementpushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalsuitjudgementpushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.judgement.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalsuitjudgementpushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalsuitjudgementpushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefereeRegistrationModel is RefereeRegistrationModel Setter
// 裁判登记信息
func (r *AlibabalegalsuitjudgementpushAPIRequest) SetRefereeRegistrationModel(_refereeRegistrationModel *RefereeRegistrationModel) error {
	r._refereeRegistrationModel = _refereeRegistrationModel
	r.Set("referee_registration_model", _refereeRegistrationModel)
	return nil
}

// GetRefereeRegistrationModel RefereeRegistrationModel Getter
func (r AlibabalegalsuitjudgementpushAPIRequest) GetRefereeRegistrationModel() *RefereeRegistrationModel {
	return r._refereeRegistrationModel
}
