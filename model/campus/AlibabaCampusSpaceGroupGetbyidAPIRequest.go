package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusspacegroupgetbyidAPIRequest 根据分组ID查询相关的空间分组信息 API请求
// alibaba.campus.space.group.getbyid
//
// 根据分组ID查询相关的空间分组信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
// HSF方法名称：getById
type AlibabacampusspacegroupgetbyidAPIRequest struct {
	model.Params
	// 用户环境
	_param0 *WorkBenchContext
	// 分组ID
	_param1 int64
}

// NewAlibabacampusspacegroupgetbyidRequest 初始化AlibabacampusspacegroupgetbyidAPIRequest对象
func NewAlibabacampusspacegroupgetbyidRequest() *AlibabacampusspacegroupgetbyidAPIRequest {
	return &AlibabacampusspacegroupgetbyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusspacegroupgetbyidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.group.getbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusspacegroupgetbyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusspacegroupgetbyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 用户环境
func (r *AlibabacampusspacegroupgetbyidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabacampusspacegroupgetbyidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 分组ID
func (r *AlibabacampusspacegroupgetbyidAPIRequest) SetParam1(_param1 int64) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabacampusspacegroupgetbyidAPIRequest) GetParam1() int64 {
	return r._param1
}
