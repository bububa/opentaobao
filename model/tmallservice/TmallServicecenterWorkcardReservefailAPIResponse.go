package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
预约失败 API返回值 
tmall.servicecenter.workcard.reservefail

服务商调用该接口回传工单预约失败
*/
type TmallServicecenterWorkcardReservefailAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkcardReservefailAPIResponseModel
}

// 预约失败 成功返回结果
type TmallServicecenterWorkcardReservefailAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_reservefail_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // -
    Result   *TmallServicecenterWorkcardReservefailResult `json:"result,omitempty" xml:"result,omitempty"`
}
