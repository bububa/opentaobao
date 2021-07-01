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
type AlitripFlightBasicDataCityQueryAllAPIRequest struct {
    model.Params
}

// 初始化AlitripFlightBasicDataCityQueryAllAPIRequest对象
func NewAlitripFlightBasicDataCityQueryAllRequest() *AlitripFlightBasicDataCityQueryAllAPIRequest{
    return &AlitripFlightBasicDataCityQueryAllAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripFlightBasicDataCityQueryAllAPIRequest) GetApiMethodName() string {
    return "alitrip.flight.basic.data.city.queryAll"
}

// IRequest interface 方法, 获取API参数
func (r AlitripFlightBasicDataCityQueryAllAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
