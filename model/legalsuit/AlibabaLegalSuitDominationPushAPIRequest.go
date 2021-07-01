package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalSuitDominationPushAPIRequest
更新或者保存管辖信息 API请求
alibaba.legal.suit.domination.push

ISV推送管辖信息到诉讼平台 */
type AlibabaLegalSuitDominationPushAPIRequest struct {
	model.Params
	// 管辖信息
	_dominationDissentModel *DominationDissentModel
}

// New
