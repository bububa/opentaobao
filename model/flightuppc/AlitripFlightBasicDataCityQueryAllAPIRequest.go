package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightbasicdatacityqueryAllAPIRequest 机票基础数据城市数据查询 API请求
// alitrip.flight.basic.data.city.queryAll
//
// 机票基础数据城市数据查询top接口
type AlitripflightbasicdatacityqueryAllAPIRequest struct {
	model.Params
}

// NewAlitripflightbasicdatacityqueryAllRequest 初始化AlitripflightbasicdatacityqueryAllAPIRequest对象
func NewAlitripflightbasicdatacityqueryAllRequest() *AlitripflightbasicdatacityqueryAllAPIRequest {
	return &AlitripflightbasicdatacityqueryAllAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripflightbasicdatacityqueryAllAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.basic.data.city.queryAll"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripflightbasicdatacityqueryAllAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripflightbasicdatacityqueryAllAPIRequest) GetRawParams() model.Params {
	return r.Params
}
