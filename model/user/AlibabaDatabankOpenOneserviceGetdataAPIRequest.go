package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadatabankopenoneservicegetdataAPIRequest 瓴羊DaaS消费者运营CGP取数接口 API请求
// alibaba.databank.open.oneservice.getdata
//
// 瓴羊DaaS消费者运营CGP取数接口
type AlibabadatabankopenoneservicegetdataAPIRequest struct {
	model.Params
	// 参数
	_dataIndicatorQueryParam *DataIndicatorQueryParam
}

// NewAlibabadatabankopenoneservicegetdataRequest 初始化AlibabadatabankopenoneservicegetdataAPIRequest对象
func NewAlibabadatabankopenoneservicegetdataRequest() *AlibabadatabankopenoneservicegetdataAPIRequest {
	return &AlibabadatabankopenoneservicegetdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadatabankopenoneservicegetdataAPIRequest) GetApiMethodName() string {
	return "alibaba.databank.open.oneservice.getdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadatabankopenoneservicegetdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadatabankopenoneservicegetdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDataIndicatorQueryParam is DataIndicatorQueryParam Setter
// 参数
func (r *AlibabadatabankopenoneservicegetdataAPIRequest) SetDataIndicatorQueryParam(_dataIndicatorQueryParam *DataIndicatorQueryParam) error {
	r._dataIndicatorQueryParam = _dataIndicatorQueryParam
	r.Set("data_indicator_query_param", _dataIndicatorQueryParam)
	return nil
}

// GetDataIndicatorQueryParam DataIndicatorQueryParam Getter
func (r AlibabadatabankopenoneservicegetdataAPIRequest) GetDataIndicatorQueryParam() *DataIndicatorQueryParam {
	return r._dataIndicatorQueryParam
}
