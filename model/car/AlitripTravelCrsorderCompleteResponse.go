package car

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
CRS接送机商家服务完成接口 APIResponse
alitrip.travel.crsorder.complete

提供给CRS接送机商家的服务完成回调接口
*/
type AlitripTravelCrsorderCompleteAPIResponse struct {
    model.CommonResponse
    AlitripTravelCrsorderCompleteResponse
}

type AlitripTravelCrsorderCompleteResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_crsorder_complete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 返回结果code
    
    MessageCode   int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`

    
}
