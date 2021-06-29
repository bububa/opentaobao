package happytrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/happytrip"
)

/* 
订单指派 
alibaba.happytrip.taxi.order.assign

通知供应商订单指派成功
*/
func AlibabaHappytripTaxiOrderAssign(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderAssignRequest, session string) (*happytrip.AlibabaHappytripTaxiOrderAssignAPIResponse, error) {
    var resp happytrip.AlibabaHappytripTaxiOrderAssignAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
