package flightuppc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
机票基础数据城市数据查询 APIRequest
alitrip.flight.basic.data.city.queryAll

机票基础数据城市数据查询top接口
*/
type AlitripFlightBasicDataCityQueryAllRequest struct {
    model.Params

}

func NewAlitripFlightBasicDataCityQueryAllRequest() *AlitripFlightBasicDataCityQueryAllRequest{
    return &AlitripFlightBasicDataCityQueryAllRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripFlightBasicDataCityQueryAllRequest) GetApiMethodName() string {
    return "alitrip.flight.basic.data.city.queryAll"
}

func (r AlitripFlightBasicDataCityQueryAllRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


