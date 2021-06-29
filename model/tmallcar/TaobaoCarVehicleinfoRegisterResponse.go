package tmallcar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全量车型导入 APIResponse
taobao.car.vehicleinfo.register

全量车型导入
*/
type TaobaoCarVehicleinfoRegisterAPIResponse struct {
    model.CommonResponse
    TaobaoCarVehicleinfoRegisterResponse
}

type TaobaoCarVehicleinfoRegisterResponse struct {
    XMLName xml.Name `xml:"car_vehicleinfo_register_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回素材id
    
    Data   bool `json:"data,omitempty" xml:"data,omitempty"`

    
    // 错误码
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 错误信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 是否成功
    
    ResultSuccess   bool `json:"result_success,omitempty" xml:"result_success,omitempty"`

    
}
