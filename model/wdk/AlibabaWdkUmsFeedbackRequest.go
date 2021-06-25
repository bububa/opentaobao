package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
质量反馈（入库辅助）-ERP下发单 APIRequest
alibaba.wdk.ums.feedback

质量反馈（入库辅助）-ERP下发单
*/
type AlibabaWdkUmsFeedbackRequest struct {
    model.Params

    // 质量反馈请求dto
    erpFeedbackdto   *ErpFeedbackDto 

}

func NewAlibabaWdkUmsFeedbackRequest() *AlibabaWdkUmsFeedbackRequest{
    return &AlibabaWdkUmsFeedbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkUmsFeedbackRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.feedback"
}

func (r AlibabaWdkUmsFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkUmsFeedbackRequest) SetErpFeedbackdto(erpFeedbackdto *ErpFeedbackDto) error {
    r.erpFeedbackdto = erpFeedbackdto
    r.Set("erp_feedbackdto", erpFeedbackdto)
    return nil
}

func (r AlibabaWdkUmsFeedbackRequest) GetErpFeedbackdto() *ErpFeedbackDto {
    return r.erpFeedbackdto
}

