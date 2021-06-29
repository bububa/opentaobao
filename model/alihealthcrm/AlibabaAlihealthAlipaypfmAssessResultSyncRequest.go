package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户测评结果回传接口 APIRequest
alibaba.alihealth.alipaypfm.assess.result.sync

用户测评结果回传接口
*/
type AlibabaAlihealthAlipaypfmAssessResultSyncRequest struct {
    model.Params

    // userId
    userId   int64 

    // 测评类型
    assessType   string 

    // 测评结果
    assessResult   string 

    // 测评结果冗余字段
    refrenceResult   string 

}

func NewAlibabaAlihealthAlipaypfmAssessResultSyncRequest() *AlibabaAlihealthAlipaypfmAssessResultSyncRequest{
    return &AlibabaAlihealthAlipaypfmAssessResultSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthAlipaypfmAssessResultSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alipaypfm.assess.result.sync"
}

func (r AlibabaAlihealthAlipaypfmAssessResultSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthAlipaypfmAssessResultSyncRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAlihealthAlipaypfmAssessResultSyncRequest) GetUserId() int64 {
    return r.userId
}

func (r *AlibabaAlihealthAlipaypfmAssessResultSyncRequest) SetAssessType(assessType string) error {
    r.assessType = assessType
    r.Set("assess_type", assessType)
    return nil
}

func (r AlibabaAlihealthAlipaypfmAssessResultSyncRequest) GetAssessType() string {
    return r.assessType
}

func (r *AlibabaAlihealthAlipaypfmAssessResultSyncRequest) SetAssessResult(assessResult string) error {
    r.assessResult = assessResult
    r.Set("assess_result", assessResult)
    return nil
}

func (r AlibabaAlihealthAlipaypfmAssessResultSyncRequest) GetAssessResult() string {
    return r.assessResult
}

func (r *AlibabaAlihealthAlipaypfmAssessResultSyncRequest) SetRefrenceResult(refrenceResult string) error {
    r.refrenceResult = refrenceResult
    r.Set("refrence_result", refrenceResult)
    return nil
}

func (r AlibabaAlihealthAlipaypfmAssessResultSyncRequest) GetRefrenceResult() string {
    return r.refrenceResult
}

