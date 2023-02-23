package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightBasicDataCityQueryAllAPIRequest 机票基础数据城市数据查询 API请求
// alitrip.flight.basic.data.city.queryAll
//
// 机票基础数据城市数据查询top接口
type AlitripFlightBasicDataCityQueryAllAPIRequest struct {
	model.Params
}

// NewAlitripFlightBasicDataCityQueryAllRequest 初始化AlitripFlightBasicDataCityQueryAllAPIRequest对象
func NewAlitripFlightBasicDataCityQueryAllRequest() *AlitripFlightBasicDataCityQueryAllAPIRequest {
	return &AlitripFlightBasicDataCityQueryAllAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripFlightBasicDataCityQueryAllAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.basic.data.city.queryAll"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripFlightBasicDataCityQueryAllAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripFlightBasicDataCityQueryAllAPIRequest) GetRawParams() model.Params {
	return r.Params
}
