package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceUnitGetbyidAPIRequest 根据ID查询指定空间单元信息 API请求
// alibaba.campus.space.unit.getbyid
//
// 根据ID查询指定空间单元信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getById
type AlibabaCampusSpaceUnitGetbyidAPIRequest struct {
	model.Params
	// 空间单元ID
	_param0 *WorkBenchContext
	// 空间单元ID
	_param1 int64
}

// NewAlibabaCampusSpaceUnitGetbyidRequest 初始化AlibabaCampusSpaceUnitGetbyidAPIRequest对象
func NewAlibabaCampusSpaceUnitGetbyidRequest() *AlibabaCampusSpaceUnitGetbyidAPIRequest {
	return &AlibabaCampusSpaceUnitGetbyidAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusSpaceUnitGetbyidAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceUnitGetbyidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.unit.getbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceUnitGetbyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusSpaceUnitGetbyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 空间单元ID
func (r *AlibabaCampusSpaceUnitGetbyidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusSpaceUnitGetbyidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 空间单元ID
func (r *AlibabaCampusSpaceUnitGetbyidAPIRequest) SetParam1(_param1 int64) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusSpaceUnitGetbyidAPIRequest) GetParam1() int64 {
	return r._param1
}

var poolAlibabaCampusSpaceUnitGetbyidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusSpaceUnitGetbyidRequest()
	},
}

// GetAlibabaCampusSpaceUnitGetbyidRequest 从 sync.Pool 获取 AlibabaCampusSpaceUnitGetbyidAPIRequest
func GetAlibabaCampusSpaceUnitGetbyidAPIRequest() *AlibabaCampusSpaceUnitGetbyidAPIRequest {
	return poolAlibabaCampusSpaceUnitGetbyidAPIRequest.Get().(*AlibabaCampusSpaceUnitGetbyidAPIRequest)
}

// ReleaseAlibabaCampusSpaceUnitGetbyidAPIRequest 将 AlibabaCampusSpaceUnitGetbyidAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusSpaceUnitGetbyidAPIRequest(v *AlibabaCampusSpaceUnitGetbyidAPIRequest) {
	v.Reset()
	poolAlibabaCampusSpaceUnitGetbyidAPIRequest.Put(v)
}
