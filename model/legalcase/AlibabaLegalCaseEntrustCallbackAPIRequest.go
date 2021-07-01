package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalCaseEntrustCallbackAPIRequest
委托回调接口 API请求
alibaba.legal.case.entrust.callback

委托回调接口 */
type AlibabaLegalCaseEntrustCallbackAPIRequest struct {
	model.Params
	// 委托id
	_entrustId int64
	// 案件id
	_caseId int64
}

// New
