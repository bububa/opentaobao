package alihealthlab

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthlab"
)

/* 
预约单核销接口 
alibaba.alihealth.reservation.order.verify

预约单核销
*/
func AlibabaAlihealthReservationOrderVerify(clt *core.SDKClient, req *alihealthlab.AlibabaAlihealthReservationOrderVerifyAPIRequest, session string) (*alihealthlab.AlibabaAlihealthReservationOrderVerifyAPIResponse, error) {
    var resp alihealthlab.AlibabaAlihealthReservationOrderVerifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
