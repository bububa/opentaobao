package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
15-供应商反馈（原料）同步接口 APIRequest
alibaba.tmallgenie.scp.plan.feedback.raw.upload

供应商反馈（原料）同步接口
*/
type AlibabaTmallgenieScpPlanFeedbackRawUploadRequest struct {
    model.Params

    // 扩展参数
    requestExtendJson   string 

}

func NewAlibabaTmallgenieScpPlanFeedbackRawUploadRequest() *AlibabaTmallgenieScpPlanFeedbackRawUploadRequest{
    return &AlibabaTmallgenieScpPlanFeedbackRawUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanFeedbackRawUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.feedback.raw.upload"
}

func (r AlibabaTmallgenieScpPlanFeedbackRawUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanFeedbackRawUploadRequest) SetRequestExtendJson(requestExtendJson string) error {
    r.requestExtendJson = requestExtendJson
    r.Set("request_extend_json", requestExtendJson)
    return nil
}

func (r AlibabaTmallgenieScpPlanFeedbackRawUploadRequest) GetRequestExtendJson() string {
    return r.requestExtendJson
}

