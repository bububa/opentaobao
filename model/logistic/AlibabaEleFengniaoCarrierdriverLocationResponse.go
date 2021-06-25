package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询骑手当前位置 APIResponse
alibaba.ele.fengniao.carrierdriver.location

查询骑手当前位置
*/
type AlibabaEleFengniaoCarrierdriverLocationAPIResponse struct {
    model.CommonResponse
    Response *AlibabaEleFengniaoCarrierdriverLocationResponse `json:"alibaba_ele_fengniao_carrierdriver_location_response,omitempty"`
}

type AlibabaEleFengniaoCarrierdriverLocationResponse struct {

    // location
    Location   *Location `json:"location,omitempty"`

    // 运单状态
    State   int64 `json:"state,omitempty"`

    // 骑手电话
    CarrierDriverPhone   string `json:"carrier_driver_phone,omitempty"`

    // 骑手姓名
    CarrierDriverName   string `json:"carrier_driver_name,omitempty"`

    // 运单状态变化时间点
    OccurredAt   int64 `json:"occurred_at,omitempty"`

}
