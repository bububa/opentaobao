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
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_domestic_rent_car_status_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 其它数据，预留，暂不使用
    
    Data   string `json:"data,omitempty" xml:"