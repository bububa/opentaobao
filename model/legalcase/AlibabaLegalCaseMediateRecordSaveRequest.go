package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增调解结果 API请求
alibaba.legal.case.mediate.record.save

增加调解沟通记录
*/
type AlibabaLegalCaseMediateRecordSaveRequest struct {
    model.Params
    // 案件id
    caseId   int64
    // 委托id
    entrustId   int64
    // 记录
    record   *MediateCommunicationModel
}

// 初始化AlibabaLegalCaseMediateRecordSaveRequest对象
func NewAlibabaLegalCaseMediateRecordSaveRequest() *AlibabaLegalCaseMediateRecordSaveRequest{
    return &AlibabaLegalCaseMediateRecordSaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseMediateRecordSaveRequest) GetApiMethodName() string {
    return "alibaba.legal.case.mediate.record.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseMediateRecordSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CaseId Setter
// 案件id
func (r *AlibabaLegalCaseMediateRecordSaveRequest) SetCaseId(caseId int64) error {
    r.caseId = caseId
    r.Set("case_id", caseId)
    return nil
}

// CaseId Getter
func (r AlibabaLegalCaseMediateRecordSaveRequest) GetCaseId() int64 {
    return r.caseId
}
// EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseMediateRecordSaveRequest) SetEntrustId(entrustId int64) error {
    r.entrustId = entrustId
    r.Set("entrust_id", entrustId)
    return nil
}

// EntrustId Getter
func (r AlibabaLegalCaseMediateRecordSaveRequest) GetEntrustId() int64 {
    return r.entrustId
}
// Record Setter
// 记录
func (r *AlibabaLegalCaseMediateRecordSaveRequest) SetRecord(record *MediateCommunicationModel) error {
    r.record = record
    r.Set("record", record)
    return nil
}

// Record Getter
func (r AlibabaLegalCaseMediateRecordSaveRequest) GetRecord() *MediateCommunicationModel {
    return r.record
}
