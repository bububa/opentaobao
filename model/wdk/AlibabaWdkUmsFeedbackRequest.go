package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
质量反馈（入库辅助）-ERP下发单 API请求
alibaba.wdk.ums.feedback

质量反馈（入库辅助）-ERP下发单
*/
type AlibabaWdkUmsFeedbackRequest struct {
    model.Params
    // 质量反馈请求dto
    erpFeedbackdto   *ErpFeedbackDto
}

// 初始化AlibabaWdkUmsFeedbackRequest对象
func NewAlibabaWdkUmsFeedbackRequest() *AlibabaWdkUmsFeedbackRequest{
    return &AlibabaWdkUmsFeedbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsFeedbackRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.feedback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ErpFeedbackdto Setter
// 质量反馈请求dto
func (r *AlibabaWdkUmsFeedbackRequest) SetErpFeedbackdto(erpFeedbackdto *ErpFeedbackDto) error {
    r.erpFeedbackdto = erpFeedbackdto
    r.Set("erp_feedbackdto", erpFeedbackdto)
    return nil
}

// ErpFeedbackdto Getter
func (r AlibabaWdkUmsFeedbackRequest) GetErpFeedbackdto() *ErpFeedbackDto {
    return r.erpFeedbackdto
}
