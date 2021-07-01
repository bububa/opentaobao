package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalCaseMediateRecordSaveAPIRequest
新增调解结果 API请求
alibaba.legal.case.mediate.record.save

增加调解沟通记录 */
type AlibabaLegalCaseMediateRecordSaveAPIRequest struct {
	model.Params
	// 案件id
	_caseId int64
	// 委托id
	_entrustId int64
	// 记录
	_record *MediateCommunicationModel
}

// New
