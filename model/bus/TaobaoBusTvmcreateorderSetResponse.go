package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机创建订单 APIResponse
taobao.bus.tvmcreateorder.set

提供给汽车票线下自助机的创建订单使用
*/
type TaobaoBusTvmcreateorderSetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_tvmcreateorder_set_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // alitripOrderId
    
    AlitripOrderId   string `json:"alitrip_order_id,omitempty" xml:"