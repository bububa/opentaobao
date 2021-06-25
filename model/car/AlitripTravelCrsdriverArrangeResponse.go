package car

import (
    "github.com/bububa/opentaobao/model"
)

/* 
CRS接送机商家派司机接口 APIResponse
alitrip.travel.crsdriver.arrange

提供给CRS接送机商家派司机的API
*/
type AlitripTravelCrsdriverArrangeAPIResponse struct {
    model.CommonResponse
    Response *AlitripTravelCrsdriverArrangeResponse `json:"alitrip_travel_crsdriver_arrange_response,omitempty"`
}

type AlitripTravelCrsdriverArrangeResponse struct {

    // 返回结果message
    Message   string `json:"message,omitempty"`

    // 返回结果code
    MessageCode   int64 `json:"message_code,omitempty"`

}
