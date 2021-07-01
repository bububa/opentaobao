package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalCaseStandpointSavestandpointAPIRequest
新增反馈口径 API请求
alibaba.legal.case.standpoint.savestandpoint

新增反馈口径 ,从外部接受反馈的口径 */
type AlibabaLegalCaseStandpointSavestandpointAPIRequest struct {
	model.Params
	// 答辩口径
	_defenseCaliber string
	// 口径描述
	_standpointDesc string
	// 案件id
	_suitId int64
	// 委托id
	_entrustId int64
	// 提交人
	_submitPeople string
}

// New
