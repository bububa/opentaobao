package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户测评结果回传接口 API请求
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

// 初始化AlibabaAlihealthAlipaypfmAssessResultSyncRequest对象
func NewAlibabaAlihealthAlipaypfmAssessResultSyncRequest() *AlibabaAlihealthAlipaypfmAssessResultSyncRequest{
    return &AlibabaAlihealthAlipaypfmAssessResultSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlipaypfmAssessResultSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alipaypfm.assess.result.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlipaypfmAssessResultSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// userId
func (r *AlibabaAlihealthAlipaypfmAssessResultSyncRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaAlihealthAlipaypfmAssessResultSyncRequest) GetUserId() int64 {
    return r.userId
}
// AssessType Setter
// 测评类型
func (r *AlibabaAlihealthAlipaypfmAssessResultSyncRequest) SetAssessType(assessType string) error {
    r.assessType = assessType
    r.Set("assess_type", assessType)
    return nil
}

// AssessType Getter
func (r AlibabaAlihealthAlipaypfmAssessResultSyncRequest) GetAssessType() string {
    return r.assessType
}
// AssessResult Setter
// 测评结果
func (r *AlibabaAlihealthAlipaypfmAssessResultSyncRequest) SetAssessResult(assessResult string) error {
    r.assessResult = assessResult
    r.Set("assess_result", assessResult)
    return nil
}

// AssessResult Getter
func (r AlibabaAlihealthAlipaypfmAssessResultSyncRequest) GetAssessResult() string {
    return r.assessResult
}
// RefrenceResult Setter
// 测评结果冗余字段
func (r *AlibabaAlihealthAlipaypfmAssessResultSyncRequest) SetRefrenceResult(refrenceResult string) error {
    r.refrenceResult = refrenceResult
    r.Set("refrence_result", refrenceResult)
    return nil
}

// RefrenceResult Getter
func (r AlibabaAlihealthAlipaypfmAssessResultSyncRequest) GetRefrenceResult() string {
    return r.refrenceResult
}
