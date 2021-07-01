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
type AlibabaWdkUmsFeedbackAPIRequest struct {
    model.Params
    // 质量反馈请求dto
    _erpFeedbackdto   *ErpFeedbackDTO
}

// 初始化AlibabaWdkUmsFeedbackAPIRequest对象
func NewAlibabaWdkUmsFeedbackRequest() *AlibabaWdkUmsFeedbackAPIRequest{
    return &AlibabaWdkUmsFeedbackAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkUmsFeedbackAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.ums.feedback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkUmsFeedbackAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ErpFeedbackdto Setter
// 质量反馈请求dto
func (r *AlibabaWdkUmsFeedbackAPIRequest) SetErpFeedbackdto(_erpFeedbackdto *ErpFeedbackDTO) error {
    r._erpFeedbackdto = _erpFeedbackdto
    r.Set("erp_feedbackdto", _erpFeedbackdto)
    return nil
}

// ErpFeedbackdto Getter
func (r AlibabaWdkUmsFeedbackAPIRequest) GetErpFeedbackdto() *ErpFeedbackDTO {
    return r._erpFeedbackdto
}
