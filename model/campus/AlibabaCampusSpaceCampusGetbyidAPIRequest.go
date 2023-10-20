package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspacecampusgetbyidAPIRequest 根据园区id获取园区信息 API请求
// alibaba.campus.space.campus.getbyid
//
// 根据园区id获取园区信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.CampusApiTopService
// HSF方法名称：getCampusById
type AlibabacampusspacecampusgetbyidAPIRequest struct {
	model.Params
	// 园区ID
	_param0 *WorkBenchContext
	// 园区ID
	_param1 int64
}

// NewAlibabacampusspacecampusgetbyidRequest 初始化AlibabacampusspacecampusgetbyidAPIRequest对象
func NewAlibabacampusspacecampusgetbyidRequest() *AlibabacampusspacecampusgetbyidAPIRequest {
	return &AlibabacampusspacecampusgetbyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusspacecampusgetbyidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.campus.getbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusspacecampusgetbyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusspacecampusgetbyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 园区ID
func (r *AlibabacampusspacecampusgetbyidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabacampusspacecampusgetbyidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 园区ID
func (r *AlibabacampusspacecampusgetbyidAPIRequest) SetParam1(_param1 int64) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabacampusspacecampusgetbyidAPIRequest) GetParam1() int64 {
	return r._param1
}
