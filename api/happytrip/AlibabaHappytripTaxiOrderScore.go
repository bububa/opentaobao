package happytrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/happytrip"
)

/* 
订单打分和评价 
alibaba.happytrip.taxi.order.score

对司机进行评分，只有订单结束后，才能进行。
*/
func AlibabaHappytripTaxiOrderScore(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiOrderScoreAPIRequest, session string) (*happytrip.AlibabaHappytripTaxiOrderScoreAPIResponse, error) {
    var resp happytrip.AlibabaHappytripTaxiOrderScoreAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
