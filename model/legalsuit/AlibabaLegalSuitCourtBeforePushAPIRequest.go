package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalSuitCourtBeforePushAPIRequest
更新或保存庭前信息 API请求
alibaba.legal.suit.court.before.push

更新或者保存庭前信息 */
type AlibabaLegalSuitCourtBeforePushAPIRequest struct {
	model.Params
	// 庭前信息
	_beforeCourtModel *BeforeCourtModel
}

// New
