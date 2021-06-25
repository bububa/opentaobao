package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改门店信息 APIResponse
tmall.servicecenter.servicestore.update

用于修改门店/网点信息。多个业务共用
*/
type TmallServicecenterServicestoreUpdateAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterServicestoreUpdateResponse `json:"tmall_servicecenter_servicestore_update_response,omitempty"`
}

type TmallServicecenterServicestoreUpdateResponse struct {

    // 方法调用结果
    Result   *TmallServicecenterServicestoreUpdateResult `json:"result,omitempty"`

}
