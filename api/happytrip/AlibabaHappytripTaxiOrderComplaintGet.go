package happytrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/happytrip"
)

/* 
投诉详情 
alibaba.happytrip.taxi.order.complaint.get

获取投诉处理进度详情
*/
func AlibabaHappytripTaxiOrderComplaintGet(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderComplaintGetAPIRequest, session string) (*happytrip.AlibabaHappytripTaxiOrderComplaintGetAPIResponse, error) {
    var resp happytrip.AlibabaHappytripTaxiOrderComplaintGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
