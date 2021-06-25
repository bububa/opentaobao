package car

import (
    "github.com/bububa/opentaobao/model"
)

/* 
司机服务状态更新接口 APIResponse
taobao.alitrip.car.driver.status.update

飞猪用车业务回调接口，用于服务商实时回传更新司机当前服务状态
*/
type TaobaoAlitripCarDriverStatusUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAlitripCarDriverStatusUpdateResponse `json:"taobao_alitrip_car_driver_status_update_response,omitempty"`
}

type TaobaoAlitripCarDriverStatusUpdateResponse struct {

    // 错误码
    MessageCode   int64 `json:"message_code,omitempty"`

    // 错误信息
    Message   string `json:"message,omitempty"`

    // 其它数据，预留，暂不使用
    Data   string `json:"data,omitempty"`

}
