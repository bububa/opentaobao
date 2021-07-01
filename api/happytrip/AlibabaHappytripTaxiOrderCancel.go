package happytrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/happytrip"
)

/* 
取消叫车 
alibaba.happytrip.taxi.order.cancel

取消叫车订单,行程中的订单不能取消
*/
func AlibabaHappytripTaxiOrderCancel(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderCancelAPIRequest, session string) (*happytrip.AlibabaHappytripTaxiOrderCancelAPIResponse, error) {
    var resp happytrip.AlibabaHappytripTaxiOrderCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
