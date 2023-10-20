package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceGroupGetlistAPIRequest 多条件查询空间分组信息 API请求
// alibaba.campus.space.group.getlist
//
// 多条件查询空间分组信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
// HSF方法名称：getList
type AlibabaCampusSpaceGroupGetlistAPIRequest struct {
	model.Params
	// 查询条件封装
	_param0 *WorkBenchContext
	// 查询参数封装
	_param1 *SpaceGroupQuery
}

// NewAlibabaCampusSpaceGroupGetlistRequest 初始化AlibabaCampusSpaceGroupGetlistAPIRequest对象
func NewAlibabaCampusSpaceGroupGetlistRequest() *AlibabaCampusSpaceGroupGetlistAPIRequest {
	return &AlibabaCampusSpaceGroupGetlistAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusSpaceGroupGetlistAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGroupGetlistAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.group.getlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGroupGetlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusSpaceGroupGetlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询条件封装
func (r *AlibabaCampusSpaceGroupGetlistAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusSpaceGroupGetlistAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 查询参数封装
func (r *AlibabaCampusSpaceGroupGetlistAPIRequest) SetParam1(_param1 *SpaceGroupQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusSpaceGroupGetlistAPIRequest) GetParam1() *SpaceGroupQuery {
	return r._param1
}

var poolAlibabaCampusSpaceGroupGetlistAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusSpaceGroupGetlistRequest()
	},
}

// GetAlibabaCampusSpaceGroupGetlistRequest 从 sync.Pool 获取 AlibabaCampusSpaceGroupGetlistAPIRequest
func GetAlibabaCampusSpaceGroupGetlistAPIRequest() *AlibabaCampusSpaceGroupGetlistAPIRequest {
	return poolAlibabaCampusSpaceGroupGetlistAPIRequest.Get().(*AlibabaCampusSpaceGroupGetlistAPIRequest)
}

// ReleaseAlibabaCampusSpaceGroupGetlistAPIRequest 将 AlibabaCampusSpaceGroupGetlistAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusSpaceGroupGetlistAPIRequest(v *AlibabaCampusSpaceGroupGetlistAPIRequest) {
	v.Reset()
	poolAlibabaCampusSpaceGroupGetlistAPIRequest.Put(v)
}
