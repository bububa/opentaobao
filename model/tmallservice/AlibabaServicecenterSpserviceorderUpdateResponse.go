package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务供应链服务单更新 API返回值 
alibaba.servicecenter.spserviceorder.update

服务供应链服务单更新，服务商通过此接口将商品的sn等信息推送到服务单中
*/
type AlibabaServicecenterSpserviceorderUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaServicecenterSpserviceorderUpdateResponse
}

// 服务供应链服务单更新 成功返回结果
type AlibabaServicecenterSpserviceorderUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_servicecenter_spserviceorder_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaServicecenterSpserviceorderUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
