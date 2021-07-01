package ascpchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpchannel"
)

/* 
ERP调用打印面单取号接口 
alibaba.ascp.uop.supplier.waybill.query

ERP调用打印面单取号接口
*/
func AlibabaAscpUopSupplierWaybillQuery(clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopSupplierWaybillQueryAPIRequest, session string) (*ascpchannel.AlibabaAscpUopSupplierWaybillQueryAPIResponse, error) {
    var resp ascpchannel.AlibabaAscpUopSupplierWaybillQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
