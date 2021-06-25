package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息通知 APIRequest
alibaba.legal.case.common.notice

同步通知给诉讼系统
*/
type AlibabaLegalCaseCommonNoticeRequest struct {
    model.Params

    // 案件id
    caseId   int64 

    // 委托id
    entrustId   int64 

    // adminicular_evidence（补充证据）risk_early_warning（风险预警）
    type   string 

    // 消息模型
    noticeModel   *NoticeModel 

}

func NewAlibabaLegalCaseCommonNoticeRequest() *AlibabaLegalCaseCommonNoticeRequest{
    return &AlibabaLegalCaseCommonNoticeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalCaseCommonNoticeRequest) GetApiMethodName() string {
    return "alibaba.legal.case.common.notice"
}

func (r AlibabaLegalCaseCommonNoticeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalCaseCommonNoticeRequest) SetCaseId(caseId int64) error {
    r.caseId = caseId
    r.Set("case_id", caseId)
    return nil
}

func (r AlibabaLegalCaseCommonNoticeRequest) GetCaseId() int64 {
    return r.caseId
}

func (r *AlibabaLegalCaseCommonNoticeRequest) SetEntrustId(entrustId int64) error {
    r.entrustId = entrustId
    r.Set("entrust_id", entrustId)
    return nil
}

func (r AlibabaLegalCaseCommonNoticeRequest) GetEntrustId() int64 {
    return r.entrustId
}

func (r *AlibabaLegalCaseCommonNoticeRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaLegalCaseCommonNoticeRequest) GetType() string {
    return r.type
}

func (r *AlibabaLegalCaseCommonNoticeRequest) SetNoticeModel(noticeModel *NoticeModel) error {
    r.noticeModel = noticeModel
    r.Set("notice_model", noticeModel)
    return nil
}

func (r AlibabaLegalCaseCommonNoticeRequest) GetNoticeModel() *NoticeModel {
    return r.noticeModel
}

