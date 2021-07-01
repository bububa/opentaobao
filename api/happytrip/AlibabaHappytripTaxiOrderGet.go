package happytrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/happytrip"
)

/* 
订单详情 
alibaba.happytrip.taxi.order.get

获取订单状态及详情
*/
func AlibabaHappytripTaxiOrderGet(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderGetAPIRequest, session string) (*happytrip.AlibabaHappytripTaxiOrderGetAPIResponse, error) {
    var resp happytrip.AlibabaHappytripTaxiOrderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
