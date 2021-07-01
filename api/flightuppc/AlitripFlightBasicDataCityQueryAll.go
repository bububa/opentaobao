package flightuppc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/flightuppc"
)

/* 
机票基础数据城市数据查询 
alitrip.flight.basic.data.city.queryAll

机票基础数据城市数据查询top接口
*/
func AlitripFlightBasicDataCityQueryAll(clt *core.SDKClient, req *flightuppc.AlitripFlightBasicDataCityQueryAllAPIRequest, session string) (*flightuppc.AlitripFlightBasicDataCityQueryAllAPIResponse, error) {
    var resp flightuppc.AlitripFlightBasicDataCityQueryAllAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
