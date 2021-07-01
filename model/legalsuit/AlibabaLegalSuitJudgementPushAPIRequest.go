package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalSuitJudgementPushAPIRequest
推送裁判登记信息给集团法务系统 API请求
alibaba.legal.suit.judgement.push

isv推送裁判登记信息给集团法务系统 */
type AlibabaLegalSuitJudgementPushAPIRequest struct {
	model.Params
	// 裁判登记信息
	_refereeRegistrationModel *RefereeRegistrationModel
}

// NewAlibabaLegalSuitJudgementPushRequest 初始化AlibabaLegalSuitJudgementPushAPIRequest对象
func NewAlibabaLegalSuitJudgementPushRequest() *AlibabaLegalSuitJudgementPushAPIRequest {
	return &AlibabaLegalSuitJudgementPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitJudgementPushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.judgement.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitJudgementPushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefereeRegistrationModel Setter
// 裁判登记信息
func (r *AlibabaLegalSuitJudgementPushAPIRequest) SetRefereeRegistrationModel(_refereeRegistrationModel *RefereeRegistrationModel) error {
	r._refereeRegistrationModel = _refereeRegistrationModel
	r.Set("referee_registration_model", _refereeRegistrationModel)
	return nil
}

// Get RefereeRegistrationModel Getter
func (r AlibabaLegalSuitJudgementPushAPIRequest) GetRefereeRegistrationModel() *RefereeRegistrationModel {
	return r._refereeRegistrationModel
}
