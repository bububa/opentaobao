package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
天机供应商发货 
alibaba.tianji.supplier.order.delivery

天机供应商发货
*/
func AlibabaTianjiSupplierOrderDelivery(clt *core.SDKClient, req *alicom.AlibabaTianjiSupplierOrderDeliveryRequest, session string) (*alicom.AlibabaTianjiSupplierOrderDeliveryAPIResponse, error) {
    var resp alicom.AlibabaTianjiSupplierOrderDeliveryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
