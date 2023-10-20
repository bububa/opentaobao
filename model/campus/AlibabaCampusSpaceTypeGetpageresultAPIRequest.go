package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspacetypegetpageresultAPIRequest 分页查询空间类别接口 API请求
// alibaba.campus.space.type.getpageresult
//
// 分页查询空间类别接口
// HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
// HSF方法名称：getPageResult
type AlibabacampusspacetypegetpageresultAPIRequest struct {
	model.Params
	// 环境参数
	_param0 *WorkBenchContext
	// 查询参数
	_param1 *SpaceTypeQuery
}

// NewAlibabacampusspacetypegetpageresultRequest 初始化AlibabacampusspacetypegetpageresultAPIRequest对象
func NewAlibabacampusspacetypegetpageresultRequest() *AlibabacampusspacetypegetpageresultAPIRequest {
	return &AlibabacampusspacetypegetpageresultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusspacetypegetpageresultAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.type.getpageresult"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusspacetypegetpageresultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusspacetypegetpageresultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 环境参数
func (r *AlibabacampusspacetypegetpageresultAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabacampusspacetypegetpageresultAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 查询参数
func (r *AlibabacampusspacetypegetpageresultAPIRequest) SetParam1(_param1 *SpaceTypeQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabacampusspacetypegetpageresultAPIRequest) GetParam1() *SpaceTypeQuery {
	return r._param1
}
