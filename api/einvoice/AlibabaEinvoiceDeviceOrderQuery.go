package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
查询税控设备加盘订购单详情 
alibaba.einvoice.device.order.query

查询税控设备订购单详情
*/
func AlibabaEinvoiceDeviceOrderQuery(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceDeviceOrderQueryRequest, session string) (*einvoice.AlibabaEinvoiceDeviceOrderQueryAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceDeviceOrderQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
