package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspacegroupgetlistAPIRequest 多条件查询空间分组信息 API请求
// alibaba.campus.space.group.getlist
//
// 多条件查询空间分组信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
// HSF方法名称：getList
type AlibabacampusspacegroupgetlistAPIRequest struct {
	model.Params
	// 查询条件封装
	_param0 *WorkBenchContext
	// 查询参数封装
	_param1 *SpaceGroupQuery
}

// NewAlibabacampusspacegroupgetlistRequest 初始化AlibabacampusspacegroupgetlistAPIRequest对象
func NewAlibabacampusspacegroupgetlistRequest() *AlibabacampusspacegroupgetlistAPIRequest {
	return &AlibabacampusspacegroupgetlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusspacegroupgetlistAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.group.getlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusspacegroupgetlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusspacegroupgetlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询条件封装
func (r *AlibabacampusspacegroupgetlistAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabacampusspacegroupgetlistAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 查询参数封装
func (r *AlibabacampusspacegroupgetlistAPIRequest) SetParam1(_param1 *SpaceGroupQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabacampusspacegroupgetlistAPIRequest) GetParam1() *SpaceGroupQuery {
	return r._param1
}
