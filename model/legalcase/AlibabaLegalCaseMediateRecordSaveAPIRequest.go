package legalcase

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalCaseMediateRecordSaveAPIRequest) Reset() {
	r._caseId = 0
	r._entrustId = 0
	r._record = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseMediateRecordSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.mediate.record.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseMediateRecordSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalCaseMediateRecordSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCaseId is CaseId Setter
// 案件id
func (r *AlibabaLegalCaseMediateRecordSaveAPIRequest) SetCaseId(_caseId int64) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// GetCaseId CaseId Getter
func (r AlibabaLegalCaseMediateRecordSaveAPIRequest) GetCaseId() int64 {
	return r._caseId
}

// SetEntrustId is EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseMediateRecordSaveAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabaLegalCaseMediateRecordSaveAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

// SetRecord is Record Setter
// 记录
func (r *AlibabaLegalCaseMediateRecordSaveAPIRequest) SetRecord(_record *MediateCommunicationModel) error {
	r._record = _record
	r.Set("record", _record)
	return nil
}

// GetRecord Record Getter
func (r AlibabaLegalCaseMediateRecordSaveAPIRequest) GetRecord() *MediateCommunicationModel {
	return r._record
}

var poolAlibabaLegalCaseMediateRecordSaveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalCaseMediateRecordSaveRequest()
	},
}

// GetAlibabaLegalCaseMediateRecordSaveRequest 从 sync.Pool 获取 AlibabaLegalCaseMediateRecordSaveAPIRequest
func GetAlibabaLegalCaseMediateRecordSaveAPIRequest() *AlibabaLegalCaseMediateRecordSaveAPIRequest {
	return poolAlibabaLegalCaseMediateRecordSaveAPIRequest.Get().(*AlibabaLegalCaseMediateRecordSaveAPIRequest)
}

// ReleaseAlibabaLegalCaseMediateRecordSaveAPIRequest 将 AlibabaLegalCaseMediateRecordSaveAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalCaseMediateRecordSaveAPIRequest(v *AlibabaLegalCaseMediateRecordSaveAPIRequest) {
	v.Reset()
	poolAlibabaLegalCaseMediateRecordSaveAPIRequest.Put(v)
}
