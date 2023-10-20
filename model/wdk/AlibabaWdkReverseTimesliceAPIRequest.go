package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkreversetimesliceAPIRequest 逆向取货时间片查询 API请求
// alibaba.wdk.reverse.timeslice
//
// 逆向取货时间片查询
type AlibabawdkreversetimesliceAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramQueryTimeSliceReq *QueryTimeSliceReq
}

// NewAlibabawdkreversetimesliceRequest 初始化AlibabawdkreversetimesliceAPIRequest对象
func NewAlibabawdkreversetimesliceRequest() *AlibabawdkreversetimesliceAPIRequest {
	return &AlibabawdkreversetimesliceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkreversetimesliceAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.timeslice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkreversetimesliceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkreversetimesliceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQueryTimeSliceReq is ParamQueryTimeSliceReq Setter
// 系统自动生成
func (r *AlibabawdkreversetimesliceAPIRequest) SetParamQueryTimeSliceReq(_paramQueryTimeSliceReq *QueryTimeSliceReq) error {
	r._paramQueryTimeSliceReq = _paramQueryTimeSliceReq
	r.Set("param_query_time_slice_req", _paramQueryTimeSliceReq)
	return nil
}

// GetParamQueryTimeSliceReq ParamQueryTimeSliceReq Getter
func (r AlibabawdkreversetimesliceAPIRequest) GetParamQueryTimeSliceReq() *QueryTimeSliceReq {
	return r._paramQueryTimeSliceReq
}
