package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDatabankOpenOneserviceGetdataAPIRequest 瓴羊DaaS消费者运营CGP取数接口 API请求
// alibaba.databank.open.oneservice.getdata
//
// 瓴羊DaaS消费者运营CGP取数接口
type AlibabaDatabankOpenOneserviceGetdataAPIRequest struct {
	model.Params
	// 参数
	_dataIndicatorQueryParam *DataIndicatorQueryParam
}

// NewAlibabaDatabankOpenOneserviceGetdataRequest 初始化AlibabaDatabankOpenOneserviceGetdataAPIRequest对象
func NewAlibabaDatabankOpenOneserviceGetdataRequest() *AlibabaDatabankOpenOneserviceGetdataAPIRequest {
	return &AlibabaDatabankOpenOneserviceGetdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDatabankOpenOneserviceGetdataAPIRequest) GetApiMethodName() string {
	return "alibaba.databank.open.oneservice.getdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDatabankOpenOneserviceGetdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDatabankOpenOneserviceGetdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDataIndicatorQueryParam is DataIndicatorQueryParam Setter
// 参数
func (r *AlibabaDatabankOpenOneserviceGetdataAPIRequest) SetDataIndicatorQueryParam(_dataIndicatorQueryParam *DataIndicatorQueryParam) error {
	r._dataIndicatorQueryParam = _dataIndicatorQueryParam
	r.Set("data_indicator_query_param", _dataIndicatorQueryParam)
	return nil
}

// GetDataIndicatorQueryParam DataIndicatorQueryParam Getter
func (r AlibabaDatabankOpenOneserviceGetdataAPIRequest) GetDataIndicatorQueryParam() *DataIndicatorQueryParam {
	return r._dataIndicatorQueryParam
}
