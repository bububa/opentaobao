package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建门店 APIResponse
tmall.servicecenter.servicestore.create

用于创建门店/网点。多个业务共用
*/
type TmallServicecenterServicestoreCreateAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterServicestoreCreateResponse `json:"tmall_servicecenter_servicestore_create_response,omitempty"`
}

type TmallServicecenterServicestoreCreateResponse struct {

    // 方法调用结果
    Result   *TmallServicecenterServicestoreCreateResult `json:"result,omitempty"`

}
