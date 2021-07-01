package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalCaseEntrustGetAPIRequest
委托 API请求
alibaba.legal.case.entrust.get

获取委托案件的基本信息 */
type AlibabaLegalCaseEntrustGetAPIRequest struct {
	model.Params
	// 委托id
	_entrustId int64
}

// New
