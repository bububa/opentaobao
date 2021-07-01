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
type AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest struct {
    model.Params
    // userId
    _userId   int64
    // 测评类型
    _assessType   string
    // 测评结果
    _assessResult   string
    // 测评结果冗余字段
    _refrenceResult   string
}

// 初始化AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest对象
func NewAlibabaAlihealthAlipaypfmAssessResultSyncRequest() *AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest{
    return &AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.alipaypfm.assess.result.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// userId
func (r *AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) GetUserId() int64 {
    return r._userId
}
// AssessType Setter
// 测评类型
func (r *AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) SetAssessType(_assessType string) error {
    r._assessType = _assessType
    r.Set("assess_type", _assessType)
    return nil
}

// AssessType Getter
func (r AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) GetAssessType() string {
    return r._assessType
}
// AssessResult Setter
// 测评结果
func (r *AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) SetAssessResult(_assessResult string) error {
    r._assessResult = _assessResult
    r.Set("assess_result", _assessResult)
    return nil
}

// AssessResult Getter
func (r AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) GetAssessResult() string {
    return r._assessResult
}
// RefrenceResult Setter
// 测评结果冗余字段
func (r *AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) SetRefrenceResult(_refrenceResult string) error {
    r._refrenceResult = _refrenceResult
    r.Set("refrence_result", _refrenceResult)
    return nil
}

// RefrenceResult Getter
func (r AlibabaAlihealthAlipaypfmAssessResultSyncAPIRequest) GetRefrenceResult() string {
    return r._refrenceResult
}
