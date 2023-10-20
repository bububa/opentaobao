package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspaceunitgetbyidAPIRequest 根据ID查询指定空间单元信息 API请求
// alibaba.campus.space.unit.getbyid
//
// 根据ID查询指定空间单元信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getById
type AlibabacampusspaceunitgetbyidAPIRequest struct {
	model.Params
	// 空间单元ID
	_param0 *WorkBenchContext
	// 空间单元ID
	_param1 int64
}

// NewAlibabacampusspaceunitgetbyidRequest 初始化AlibabacampusspaceunitgetbyidAPIRequest对象
func NewAlibabacampusspaceunitgetbyidRequest() *AlibabacampusspaceunitgetbyidAPIRequest {
	return &AlibabacampusspaceunitgetbyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusspaceunitgetbyidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.unit.getbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusspaceunitgetbyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusspaceunitgetbyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 空间单元ID
func (r *AlibabacampusspaceunitgetbyidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabacampusspaceunitgetbyidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 空间单元ID
func (r *AlibabacampusspaceunitgetbyidAPIRequest) SetParam1(_param1 int64) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabacampusspaceunitgetbyidAPIRequest) GetParam1() int64 {
	return r._param1
}
