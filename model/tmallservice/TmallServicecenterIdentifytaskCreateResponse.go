package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务商创建核销单 APIResponse
tmall.servicecenter.identifytask.create

服务商调用该接口进行创建核销单操作
*/
type TmallServicecenterIdentifytaskCreateAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterIdentifytaskCreateResponse `json:"tmall_servicecenter_identifytask_create_response,omitempty"`
}

type TmallServicecenterIdentifytaskCreateResponse struct {

    // -
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
