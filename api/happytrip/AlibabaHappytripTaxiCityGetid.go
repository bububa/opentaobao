package happytrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/happytrip"
)

/* 
获取城市id 
alibaba.happytrip.taxi.city.getid

通过经纬度坐标返回城市id
*/
func AlibabaHappytripTaxiCityGetid(clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiCityGetidAPIRequest, session string) (*happytrip.AlibabaHappytripTaxiCityGetidAPIResponse, error) {
    var resp happytrip.AlibabaHappytripTaxiCityGetidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
