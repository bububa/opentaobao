package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
ISV 新增/修改复诊预约信息 
alibaba.alihealth.booking.reserve.rise

ISV 新增/修改复诊预约信息
*/
func AlibabaAlihealthBookingReserveRise(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBookingReserveRiseAPIRequest, session string) (*alihealth2.AlibabaAlihealthBookingReserveRiseAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthBookingReserveRiseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
