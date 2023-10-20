package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitJudgementPushAPIRequest 推送裁判登记信息给集团法务系统 API请求
// alibaba.legal.suit.judgement.push
//
// isv推送裁判登记信息给集团法务系统
type AlibabaLegalSuitJudgementPushAPIRequest struct {
	model.Params
	// 裁判登记信息
	_refereeRegistrationModel *RefereeRegistrationModel
}

// NewAlibabaLegalSuitJudgementPushRequest 初始化AlibabaLegalSuitJudgementPushAPIRequest对象
func NewAlibabaLegalSuitJudgementPushRequest() *AlibabaLegalSuitJudgementPushAPIRequest {
	return &AlibabaLegalSuitJudgementPushAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalSuitJudgementPushAPIRequest) Reset() {
	r._refereeRegistrationModel = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitJudgementPushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.judgement.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitJudgementPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalSuitJudgementPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefereeRegistrationModel is RefereeRegistrationModel Setter
// 裁判登记信息
func (r *AlibabaLegalSuitJudgementPushAPIRequest) SetRefereeRegistrationModel(_refereeRegistrationModel *RefereeRegistrationModel) error {
	r._refereeRegistrationModel = _refereeRegistrationModel
	r.Set("referee_registration_model", _refereeRegistrationModel)
	return nil
}

// GetRefereeRegistrationModel RefereeRegistrationModel Getter
func (r AlibabaLegalSuitJudgementPushAPIRequest) GetRefereeRegistrationModel() *RefereeRegistrationModel {
	return r._refereeRegistrationModel
}

var poolAlibabaLegalSuitJudgementPushAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalSuitJudgementPushRequest()
	},
}

// GetAlibabaLegalSuitJudgementPushRequest 从 sync.Pool 获取 AlibabaLegalSuitJudgementPushAPIRequest
func GetAlibabaLegalSuitJudgementPushAPIRequest() *AlibabaLegalSuitJudgementPushAPIRequest {
	return poolAlibabaLegalSuitJudgementPushAPIRequest.Get().(*AlibabaLegalSuitJudgementPushAPIRequest)
}

// ReleaseAlibabaLegalSuitJudgementPushAPIRequest 将 AlibabaLegalSuitJudgementPushAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalSuitJudgementPushAPIRequest(v *AlibabaLegalSuitJudgementPushAPIRequest) {
	v.Reset()
	poolAlibabaLegalSuitJudgementPushAPIRequest.Put(v)
}
