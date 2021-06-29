package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
15-供应商反馈（原料）同步接口 API请求
alibaba.tmallgenie.scp.plan.feedback.raw.upload

供应商反馈（原料）同步接口
*/
type AlibabaTmallgenieScpPlanFeedbackRawUploadRequest struct {
    model.Params
    // 扩展参数
    _requestExtendJson   string
}

// 初始化AlibabaTmallgenieScpPlanFeedbackRawUploadRequest对象
func NewAlibabaTmallgenieScpPlanFeedbackRawUploadRequest() *AlibabaTmallgenieScpPlanFeedbackRawUploadRequest{
    return &AlibabaTmallgenieScpPlanFeedbackRawUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanFeedbackRawUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.feedback.raw.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanFeedbackRawUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanFeedbackRawUploadRequest) SetRequestExtendJson(_requestExtendJson string) error {
    r._requestExtendJson = _requestExtendJson
    r.Set("request_extend_json", _requestExtendJson)
    return nil
}

// RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanFeedbackRawUploadRequest) GetRequestExtendJson() string {
    return r._requestExtendJson
}
