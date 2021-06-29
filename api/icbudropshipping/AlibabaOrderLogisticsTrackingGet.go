package icbudropshipping

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbudropshipping"
)

/* 
阿里巴巴订单物流轨迹查询 
alibaba.order.logistics.tracking.get

阿里巴巴订单物流轨迹查询
*/
func AlibabaOrderLogisticsTrackingGet(clt *core.SDKClient, req *icbudropshipping.AlibabaOrderLogisticsTrackingGetRequest, session string) (*icbudropshipping.AlibabaOrderLogisticsTrackingGetAPIResponse, error) {
    var resp icbudropshipping.AlibabaOrderLogisticsTrackingGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
