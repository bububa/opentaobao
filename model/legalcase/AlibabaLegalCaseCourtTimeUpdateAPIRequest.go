package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalCaseCourtTimeUpdateAPIRequest
开庭时间变更 API请求
alibaba.legal.case.court.time.update

修改案件的开庭时间 */
type AlibabaLegalCaseCourtTimeUpdateAPIRequest struct {
	model.Params
	// 案件id
	_caseId int64
	// 委托id
	_entrustId int64
	// 开庭时间
	_courtTime string
}

// New
