package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspaceunitgetlistbygroupidAPIRequest 根据分组ID查询相应的空间单元 API请求
// alibaba.campus.space.unit.getlistbygroupid
//
// 根据分组ID查询相应的空间单元
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getListByGroupId
type AlibabacampusspaceunitgetlistbygroupidAPIRequest struct {
	model.Params
	// 分组ID
	_param0 *WorkBenchContext
	// 分组ID
	_param1 int64
}

// NewAlibabacampusspaceunitgetlistbygroupidRequest 初始化AlibabacampusspaceunitgetlistbygroupidAPIRequest对象
func NewAlibabacampusspaceunitgetlistbygroupidRequest() *AlibabacampusspaceunitgetlistbygroupidAPIRequest {
	return &AlibabacampusspaceunitgetlistbygroupidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusspaceunitgetlistbygroupidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.unit.getlistbygroupid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusspaceunitgetlistbygroupidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusspaceunitgetlistbygroupidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 分组ID
func (r *AlibabacampusspaceunitgetlistbygroupidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabacampusspaceunitgetlistbygroupidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 分组ID
func (r *AlibabacampusspaceunitgetlistbygroupidAPIRequest) SetParam1(_param1 int64) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabacampusspaceunitgetlistbygroupidAPIRequest) GetParam1() int64 {
	return r._param1
}
