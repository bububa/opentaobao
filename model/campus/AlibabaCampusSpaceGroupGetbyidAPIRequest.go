package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceGroupGetbyidAPIRequest 根据分组ID查询相关的空间分组信息 API请求
// alibaba.campus.space.group.getbyid
//
// 根据分组ID查询相关的空间分组信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
// HSF方法名称：getById
type AlibabaCampusSpaceGroupGetbyidAPIRequest struct {
	model.Params
	// 用户环境
	_param0 *WorkBenchContext
	// 分组ID
	_param1 int64
}

// NewAlibabaCampusSpaceGroupGetbyidRequest 初始化AlibabaCampusSpaceGroupGetbyidAPIRequest对象
func NewAlibabaCampusSpaceGroupGetbyidRequest() *AlibabaCampusSpaceGroupGetbyidAPIRequest {
	return &AlibabaCampusSpaceGroupGetbyidAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusSpaceGroupGetbyidAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceGroupGetbyidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.group.getbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceGroupGetbyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusSpaceGroupGetbyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 用户环境
func (r *AlibabaCampusSpaceGroupGetbyidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusSpaceGroupGetbyidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 分组ID
func (r *AlibabaCampusSpaceGroupGetbyidAPIRequest) SetParam1(_param1 int64) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusSpaceGroupGetbyidAPIRequest) GetParam1() int64 {
	return r._param1
}

var poolAlibabaCampusSpaceGroupGetbyidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusSpaceGroupGetbyidRequest()
	},
}

// GetAlibabaCampusSpaceGroupGetbyidRequest 从 sync.Pool 获取 AlibabaCampusSpaceGroupGetbyidAPIRequest
func GetAlibabaCampusSpaceGroupGetbyidAPIRequest() *AlibabaCampusSpaceGroupGetbyidAPIRequest {
	return poolAlibabaCampusSpaceGroupGetbyidAPIRequest.Get().(*AlibabaCampusSpaceGroupGetbyidAPIRequest)
}

// ReleaseAlibabaCampusSpaceGroupGetbyidAPIRequest 将 AlibabaCampusSpaceGroupGetbyidAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusSpaceGroupGetbyidAPIRequest(v *AlibabaCampusSpaceGroupGetbyidAPIRequest) {
	v.Reset()
	poolAlibabaCampusSpaceGroupGetbyidAPIRequest.Put(v)
}
