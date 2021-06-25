package car

import (
    "github.com/bububa/opentaobao/model"
)

/* 
航旅国内租车订单状态更新 APIResponse
taobao.alitrip.domestic.rent.car.status.update

航旅国内租车订单状态更新
*/
type TaobaoAlitripDomesticRentCarStatusUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAlitripDomesticRentCarStatusUpdateResponse `json:"taobao_alitrip_domestic_rent_car_status_update_response,omitempty"`
}

type TaobaoAlitripDomesticRentCarStatusUpdateResponse struct {

    // 其它数据，预留，暂不使用
    Data   string `json:"data,omitempty"`

    // 错误信息
    Message   string `json:"message,omitempty"`

    // 错误码.code为0时表示成功
    MessageCode   int64 `json:"message_code,omitempty"`

}
