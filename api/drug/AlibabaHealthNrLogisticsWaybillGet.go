package drug

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drug"
)

/* 
电子面单查询接口 
alibaba.health.nr.logistics.waybill.get

商家登录后根据订单号查询物流单号及电子面单信息
*/
func AlibabaHealthNrLogisticsWaybillGet(clt *core.SDKClient, req *drug.AlibabaHealthNrLogisticsWaybillGetRequest, session string) (*drug.AlibabaHealthNrLogisticsWaybillGetAPIResponse, error) {
    var resp drug.AlibabaHealthNrLogisticsWaybillGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
