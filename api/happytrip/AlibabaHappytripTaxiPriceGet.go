package happytrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/happytrip"
)

/* 
获取价格预估信息 
alibaba.happytrip.taxi.price.get

打车价格预估
*/
func AlibabaHappytripTaxiPriceGet(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiPriceGetRequest, session string) (*happytrip.AlibabaHappytripTaxiPriceGetAPIResponse, error) {
    var resp happytrip.AlibabaHappytripTaxiPriceGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
