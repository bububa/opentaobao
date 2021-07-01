package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
14-供应商反馈（OEM）同步接口 API请求
alibaba.tmallgenie.scp.plan.feedback.oem.upload

供应商反馈（OEM）同步接口
*/
type AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest struct {
    model.Params
    // 扩展参数
    _requestExtendJson   string
}

// 初始化AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanFeedbackOemUploadRequest() *AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest{
    return &AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.feedback.oem.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
    r._requestExtendJson = _requestExtendJson
    r.Set("request_extend_json", _requestExtendJson)
    return nil
}

// RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest) GetRequestExtendJson() string {
    return r._requestExtendJson
}
