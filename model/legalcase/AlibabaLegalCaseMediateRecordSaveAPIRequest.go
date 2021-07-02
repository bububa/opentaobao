package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseMediateRecordSaveAPIRequest 新增调解结果 API请求
// alibaba.legal.case.mediate.record.save
//
// 增加调解沟通记录
type AlibabaLegalCaseMediateRecordSaveAPIRequest struct {
	model.Params
	// 案件id
	_caseId int64
	// 委托id
	_entrustId int64
	// 记录
	_record *MediateCommunicationModel
}

// NewAlibabaLegalCaseMediateRecordSaveRequest 初始化AlibabaLegalCaseMediateRecordSaveAPIRequest对象
func NewAlibabaLegalCaseMediateRecordSaveRequest() *AlibabaLegalCaseMediateRecordSaveAPIRequest {
	return &AlibabaLegalCaseMediateRecordSaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseMediateRecordSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.mediate.record.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseMediateRecordSaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CaseId Setter
// 案件id
func (r *AlibabaLegalCaseMediateRecordSaveAPIRequest) SetCaseId(_caseId int64) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// Get CaseId Getter
func (r AlibabaLegalCaseMediateRecordSaveAPIRequest) GetCaseId() int64 {
	return r._caseId
}

// Set is EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseMediateRecordSaveAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// Get EntrustId Getter
func (r AlibabaLegalCaseMediateRecordSaveAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

// Set is Record Setter
// 记录
func (r *AlibabaLegalCaseMediateRecordSaveAPIRequest) SetRecord(_record *MediateCommunicationModel) error {
	r._record = _record
	r.Set("record", _record)
	return nil
}

// Get Record Getter
func (r AlibabaLegalCaseMediateRecordSaveAPIRequest) GetRecord() *MediateCommunicationModel {
	return r._record
}
