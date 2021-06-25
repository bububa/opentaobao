package car

import (
    "github.com/bububa/opentaobao/model"
)

/* 
CRS接送机商家服务完成接口 APIResponse
alitrip.travel.crsorder.complete

提供给CRS接送机商家的服务完成回调接口
*/
type AlitripTravelCrsorderCompleteAPIResponse struct {
    model.CommonResponse
    Response *AlitripTravelCrsorderCompleteResponse `json:"alitrip_travel_crsorder_complete_response,omitempty"`
}

type AlitripTravelCrsorderCompleteResponse struct {

    // 返回结果message
    Message   string `json:"message,omitempty"`

    // 返回结果code
    MessageCode   int64 `json:"message_code,omitempty"`

}
