package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
查询供应商订单 
alibaba.tianji.supplier.order.query

查询供应商订单
*/
func AlibabaTianjiSupplierOrderQuery(clt *core.SDKClient, req *alicom.AlibabaTianjiSupplierOrderQueryRequest, session string) (*alicom.AlibabaTianjiSupplierOrderQueryAPIResponse, error) {
    var resp alicom.AlibabaTianjiSupplierOrderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
