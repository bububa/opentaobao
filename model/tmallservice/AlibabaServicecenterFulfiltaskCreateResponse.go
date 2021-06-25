package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
合单生成核销单 APIResponse
alibaba.servicecenter.fulfiltask.create

服务对工单进行合单，合单的结果是生成核销单
*/
type AlibabaServicecenterFulfiltaskCreateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaServicecenterFulfiltaskCreateResponse `json:"alibaba_servicecenter_fulfiltask_create_response,omitempty"`
}

type AlibabaServicecenterFulfiltaskCreateResponse struct {

    // 接口返回model
    Result   *AlibabaServicecenterFulfiltaskCreateResult `json:"result,omitempty"`

}
