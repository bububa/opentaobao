package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceUnitGetlistAPIRequest 多条件查询空间单元信息 API请求
// alibaba.campus.space.unit.getlist
//
// 多条件查询空间单元信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getList
type AlibabaCampusSpaceUnitGetlistAPIRequest struct {
	model.Params
	// 查询条件封装
	_param0 *WorkBenchContext
	// 查询参数封装
	_param1 *SpaceUnitQuery
}

// NewAlibabaCampusSpaceUnitGetlistRequest 初始化AlibabaCampusSpaceUnitGetlistAPIRequest对象
func NewAlibabaCampusSpaceUnitGetlistRequest() *AlibabaCampusSpaceUnitGetlistAPIRequest {
	return &AlibabaCampusSpaceUnitGetlistAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusSpaceUnitGetlistAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetlistAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.unit.getlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusSpaceUnitGetlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询条件封装
func (r *AlibabaCampusSpaceUnitGetlistAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusSpaceUnitGetlistAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 查询参数封装
func (r *AlibabaCampusSpaceUnitGetlistAPIRequest) SetParam1(_param1 *SpaceUnitQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusSpaceUnitGetlistAPIRequest) GetParam1() *SpaceUnitQuery {
	return r._param1
}

var poolAlibabaCampusSpaceUnitGetlistAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusSpaceUnitGetlistRequest()
	},
}

// GetAlibabaCampusSpaceUnitGetlistRequest 从 sync.Pool 获取 AlibabaCampusSpaceUnitGetlistAPIRequest
func GetAlibabaCampusSpaceUnitGetlistAPIRequest() *AlibabaCampusSpaceUnitGetlistAPIRequest {
	return poolAlibabaCampusSpaceUnitGetlistAPIRequest.Get().(*AlibabaCampusSpaceUnitGetlistAPIRequest)
}

// ReleaseAlibabaCampusSpaceUnitGetlistAPIRequest 将 AlibabaCampusSpaceUnitGetlistAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusSpaceUnitGetlistAPIRequest(v *AlibabaCampusSpaceUnitGetlistAPIRequest) {
	v.Reset()
	poolAlibabaCampusSpaceUnitGetlistAPIRequest.Put(v)
}
