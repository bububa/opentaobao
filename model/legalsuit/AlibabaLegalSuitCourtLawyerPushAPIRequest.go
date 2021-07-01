package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalSuitCourtLawyerPushAPIRequest
推荐律师接口 API请求
alibaba.legal.suit.court.lawyer.push

为诉讼系统推荐律师 */
type AlibabaLegalSuitCourtLawyerPushAPIRequest struct {
	model.Params
	// 委托ID
	_entrustId int64
	// 案件ID
	_suitId int64
	// 推荐律师模型
	_lawyersModel *LawyersModel
}

// New
