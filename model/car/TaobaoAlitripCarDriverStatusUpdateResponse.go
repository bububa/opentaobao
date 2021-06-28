package car

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
司机服务状态更新接口 APIResponse
taobao.alitrip.car.driver.status.update

飞猪用车业务回调接口，用于服务商实时回传更新司机当前服务状态
*/
type TaobaoAlitripCarDriverStatusUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_car_driver_status_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    MessageCode   int64 `json:"message_code,omitempty" xml:"