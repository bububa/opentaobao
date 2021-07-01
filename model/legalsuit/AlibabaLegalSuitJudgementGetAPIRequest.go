package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalSuitJudgementGetAPIRequest
获取裁判登记信息 API请求
alibaba.legal.suit.judgement.get

供ISV供应商获取集团法务系统的裁判登记信息 */
type AlibabaLegalSuitJudgementGetAPIRequest struct {
	model.Params
	// 案件id
	_suitId int64
}

// New
