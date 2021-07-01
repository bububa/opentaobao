package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalSuitCourtOpenPushAPIRequest
开庭信息推送接口 API请求
alibaba.legal.suit.court.open.push

供ISV推送开庭信息给集团诉讼 */
type AlibabaLegalSuitCourtOpenPushAPIRequest struct {
	model.Params
	// 开庭信息
	_courtInfoModel *CourtInfoModel
}

// New
