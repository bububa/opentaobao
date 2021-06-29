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
type AlibabaTmallgenieScpPlanFeedbackOemUploadRequest struct {
    model.Params
    // 扩展参数
    requestExtendJson   string
}

// 初始化AlibabaTmallgenieScpPlanFeedbackOemUploadRequest对象
func NewAlibabaTmallgenieScpPlanFeedbackOemUploadRequest() *AlibabaTmallgenieScpPlanFeedbackOemUploadRequest{
    return &AlibabaTmallgenieScpPlanFeedbackOemUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanFeedbackOemUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.feedback.oem.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanFeedbackOemUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanFeedbackOemUploadRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

// RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanFeedbackOemUploadRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}
