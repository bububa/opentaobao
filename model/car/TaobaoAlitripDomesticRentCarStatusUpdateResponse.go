package car

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
航旅国内租车订单状态更新 APIResponse
taobao.alitrip.domestic.rent.car.status.update

航旅国内租车订单状态更新
*/
type TaobaoAlitripDomesticRentCarStatusUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripDomesticRentCarStatusUpdateResponse
}

type TaobaoAlitripDomesticRentCarStatusUpdateResponse struct {
    XMLName xml.Name `xml:"alitrip_domestic_rent_car_status_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 其它数据，预留，暂不使用
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
    // 错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 错误码.code为0时表示成功
    
    MessageCode   int64 `json:"message_code,omitempty" xml:"message_code,omitempty"`

    
}
