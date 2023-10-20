package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceCampusGetbyidAPIRequest 根据园区id获取园区信息 API请求
// alibaba.campus.space.campus.getbyid
//
// 根据园区id获取园区信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.CampusApiTopService
// HSF方法名称：getCampusById
type AlibabaCampusSpaceCampusGetbyidAPIRequest struct {
	model.Params
	// 园区ID
	_param0 *WorkBenchContext
	// 园区ID
	_param1 int64
}

// NewAlibabaCampusSpaceCampusGetbyidRequest 初始化AlibabaCampusSpaceCampusGetbyidAPIRequest对象
func NewAlibabaCampusSpaceCampusGetbyidRequest() *AlibabaCampusSpaceCampusGetbyidAPIRequest {
	return &AlibabaCampusSpaceCampusGetbyidAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusSpaceCampusGetbyidAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceCampusGetbyidAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.campus.getbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceCampusGetbyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusSpaceCampusGetbyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 园区ID
func (r *AlibabaCampusSpaceCampusGetbyidAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaCampusSpaceCampusGetbyidAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// SetParam1 is Param1 Setter
// 园区ID
func (r *AlibabaCampusSpaceCampusGetbyidAPIRequest) SetParam1(_param1 int64) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaCampusSpaceCampusGetbyidAPIRequest) GetParam1() int64 {
	return r._param1
}

var poolAlibabaCampusSpaceCampusGetbyidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusSpaceCampusGetbyidRequest()
	},
}

// GetAlibabaCampusSpaceCampusGetbyidRequest 从 sync.Pool 获取 AlibabaCampusSpaceCampusGetbyidAPIRequest
func GetAlibabaCampusSpaceCampusGetbyidAPIRequest() *AlibabaCampusSpaceCampusGetbyidAPIRequest {
	return poolAlibabaCampusSpaceCampusGetbyidAPIRequest.Get().(*AlibabaCampusSpaceCampusGetbyidAPIRequest)
}

// ReleaseAlibabaCampusSpaceCampusGetbyidAPIRequest 将 AlibabaCampusSpaceCampusGetbyidAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusSpaceCampusGetbyidAPIRequest(v *AlibabaCampusSpaceCampusGetbyidAPIRequest) {
	v.Reset()
	poolAlibabaCampusSpaceCampusGetbyidAPIRequest.Put(v)
}
