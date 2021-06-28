package drug

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drug"
)

/* 
上传订单同城快递单号 
alibaba.health.nr.logistics.deliveryno.update

上传订单同城快递单号
*/
func AlibabaHealthNrLogisticsDeliverynoUpdate(clt *core.SDKClient, req *drug.AlibabaHealthNrLogisticsDeliverynoUpdateRequest, session string) (*drug.AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse, error) {
    var resp drug.AlibabaHealthNrLogisticsDeliverynoUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
