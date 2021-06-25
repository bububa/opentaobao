package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务供应链服务单更新 APIResponse
alibaba.servicecenter.spserviceorder.update

服务供应链服务单更新，服务商通过此接口将商品的sn等信息推送到服务单中
*/
type AlibabaServicecenterSpserviceorderUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaServicecenterSpserviceorderUpdateResponse `json:"alibaba_servicecenter_spserviceorder_update_response,omitempty"`
}

type AlibabaServicecenterSpserviceorderUpdateResponse struct {

    // 接口返回model
    Result   *AlibabaServicecenterSpserviceorderUpdateResult `json:"result,omitempty"`

}
