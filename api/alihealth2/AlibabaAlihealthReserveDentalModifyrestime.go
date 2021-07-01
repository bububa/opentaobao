package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
修改预约时间 
alibaba.alihealth.reserve.dental.modifyrestime

修改预约时间
*/
func AlibabaAlihealthReserveDentalModifyrestime(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthReserveDentalModifyrestimeAPIRequest, session string) (*alihealth2.AlibabaAlihealthReserveDentalModifyrestimeAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthReserveDentalModifyrestimeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
