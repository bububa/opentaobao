package flightuppc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
机票基础数据城市数据查询 API请求
alitrip.flight.basic.data.city.queryAll

机票基础数据城市数据查询top接口
*/
type AlitripFlightBasicDataCityQueryAllRequest struct {
    model.Params
}

// 初始化AlitripFlightBasicDataCityQueryAllRequest对象
func NewAlitripFlightBasicDataCityQueryAllRequest() *AlitripFlightBasicDataCityQueryAllRequest{
    return &AlitripFlightBasicDataCityQueryAllRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripFlightBasicDataCityQueryAllRequest) GetApiMethodName() string {
    return "alitrip.flight.basic.data.city.queryAll"
}

// IRequest interface 方法, 获取API参数
func (r AlitripFlightBasicDataCityQueryAllRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
