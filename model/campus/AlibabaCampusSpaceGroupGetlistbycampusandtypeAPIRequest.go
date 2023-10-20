package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest 根据园区id及TypeId获取空间分组 API请求
// alibaba.campus.space.group.getlistbycampusandtype
//
// 根据园区id及TypeId获取空间分组
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
// HSF方法名称：getListByCampusAndType
type AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *WorkBenchContext
	// 查询参数封装
	_param1 *SpaceGroupQuery
}

// NewAlibabaCampusSpaceGroupGetlistbycampusandtypeRequest 初始化AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest对象
func NewAlibabaCampusSpaceGroupGetlistbycampusandtypeRequest() *AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest {
	return &AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.group.getlistbycampusandtype"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 查询参数封装
func (r *AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest) SetParam1(_param1 *SpaceGroupQuery) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest) GetParam1() *SpaceGroupQuery {
	return r._param1
}

var poolAlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusSpaceGroupGetlistbycampusandtypeRequest()
	},
}

// GetAlibabaCampusSpaceGroupGetlistbycampusandtypeRequest 从 sync.Pool 获取 AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest
func GetAlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest() *AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest {
	return poolAlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest.Get().(*AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest)
}

// ReleaseAlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest 将 AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest(v *AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest) {
	v.Reset()
	poolAlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest.Put(v)
}
