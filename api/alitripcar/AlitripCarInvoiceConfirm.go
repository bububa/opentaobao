package alitripcar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alitripcar"
)

/* 
发票确认接口 
alitrip.car.invoice.confirm

飞猪发票回调接口
*/
func AlitripCarInvoiceConfirm(clt *core.SDKClient, req *alitripcar.AlitripCarInvoiceConfirmRequest, session string) (*alitripcar.AlitripCarInvoiceConfirmAPIResponse, error) {
    var resp alitripcar.AlitripCarInvoiceConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
