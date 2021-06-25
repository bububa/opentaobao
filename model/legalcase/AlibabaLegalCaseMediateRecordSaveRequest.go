package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增调解结果 APIRequest
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

func NewAlibabaLegalCaseMediateRecordSaveRequest() *AlibabaLegalCaseMediateRecordSaveRequest{
    return &AlibabaLegalCaseMediateRecordSaveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalCaseMediateRecordSaveRequest) GetApiMethodName() string {
    return "alibaba.legal.case.mediate.record.save"
}

func (r AlibabaLegalCaseMediateRecordSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalCaseMediateRecordSaveRequest) SetCaseId(caseId int64) error {
    r.caseId = caseId
    r.Set("case_id", caseId)
    return nil
}

func (r AlibabaLegalCaseMediateRecordSaveRequest) GetCaseId() int64 {
    return r.caseId
}

func (r *AlibabaLegalCaseMediateRecordSaveRequest) SetEntrustId(entrustId int64) error {
    r.entrustId = entrustId
    r.Set("entrust_id", entrustId)
    return nil
}

func (r AlibabaLegalCaseMediateRecordSaveRequest) GetEntrustId() int64 {
    return r.entrustId
}

func (r *AlibabaLegalCaseMediateRecordSaveRequest) SetRecord(record *MediateCommunicationModel) error {
    r.record = record
    r.Set("record", record)
    return nil
}

func (r AlibabaLegalCaseMediateRecordSaveRequest) GetRecord() *MediateCommunicationModel {
    return r.record
}

