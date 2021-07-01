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

// New
