package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalcasemediaterecordsaveAPIRequest 新增调解结果 API请求
// alibaba.legal.case.mediate.record.save
//
// 增加调解沟通记录
type AlibabalegalcasemediaterecordsaveAPIRequest struct {
	model.Params
	// 案件id
	_caseId int64
	// 委托id
	_entrustId int64
	// 记录
	_record *MediateCommunicationModel
}

// NewAlibabalegalcasemediaterecordsaveRequest 初始化AlibabalegalcasemediaterecordsaveAPIRequest对象
func NewAlibabalegalcasemediaterecordsaveRequest() *AlibabalegalcasemediaterecordsaveAPIRequest {
	return &AlibabalegalcasemediaterecordsaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalcasemediaterecordsaveAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.mediate.record.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalcasemediaterecordsaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalcasemediaterecordsaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCaseId is CaseId Setter
// 案件id
func (r *AlibabalegalcasemediaterecordsaveAPIRequest) SetCaseId(_caseId int64) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// GetCaseId CaseId Getter
func (r AlibabalegalcasemediaterecordsaveAPIRequest) GetCaseId() int64 {
	return r._caseId
}

// SetEntrustId is EntrustId Setter
// 委托id
func (r *AlibabalegalcasemediaterecordsaveAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabalegalcasemediaterecordsaveAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

// SetRecord is Record Setter
// 记录
func (r *AlibabalegalcasemediaterecordsaveAPIRequest) SetRecord(_record *MediateCommunicationModel) error {
	r._record = _record
	r.Set("record", _record)
	return nil
}

// GetRecord Record Getter
func (r AlibabalegalcasemediaterecordsaveAPIRequest) GetRecord() *MediateCommunicationModel {
	return r._record
}
