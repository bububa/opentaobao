package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
修改预约 
alibaba.alihealth.booking.reserve.modify

消费医疗统一预约平台，取消预约
*/
func AlibabaAlihealthBookingReserveModify(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBookingReserveModifyRequest, session string) (*alihealth2.AlibabaAlihealthBookingReserveModifyAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthBookingReserveModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
