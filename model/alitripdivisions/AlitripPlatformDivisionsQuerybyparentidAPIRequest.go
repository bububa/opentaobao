package alitripdivisions

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPlatformDivisionsQuerybyparentidAPIRequest 根据父节点id查询下级行政区划数据 API请求
// alitrip.platform.divisions.querybyparentid
//
// 根据行政区划id查询下一层级行政区划数据
type AlitripPlatformDivisionsQuerybyparentidAPIRequest struct {
	model.Params
	// 行政区划父id
	_paramLong int64
}

// NewAlitripPlatformDivisionsQuerybyparentidRequest 初始化AlitripPlatformDivisionsQuerybyparentidAPIRequest对象
func NewAlitripPlatformDivisionsQuerybyparentidRequest() *AlitripPlatformDivisionsQuerybyparentidAPIRequest {
	return &AlitripPlatformDivisionsQuerybyparentidAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripPlatformDivisionsQuerybyparentidAPIRequest) Reset() {
	r._paramLong = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPlatformDivisionsQuerybyparentidAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.divisions.querybyparentid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPlatformDivisionsQuerybyparentidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripPlatformDivisionsQuerybyparentidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamLong is ParamLong Setter
// 行政区划父id
func (r *AlitripPlatformDivisionsQuerybyparentidAPIRequest) SetParamLong(_paramLong int64) error {
	r._paramLong = _paramLong
	r.Set("param_long", _paramLong)
	return nil
}

// GetParamLong ParamLong Getter
func (r AlitripPlatformDivisionsQuerybyparentidAPIRequest) GetParamLong() int64 {
	return r._paramLong
}

var poolAlitripPlatformDivisionsQuerybyparentidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripPlatformDivisionsQuerybyparentidRequest()
	},
}

// GetAlitripPlatformDivisionsQuerybyparentidRequest 从 sync.Pool 获取 AlitripPlatformDivisionsQuerybyparentidAPIRequest
func GetAlitripPlatformDivisionsQuerybyparentidAPIRequest() *AlitripPlatformDivisionsQuerybyparentidAPIRequest {
	return poolAlitripPlatformDivisionsQuerybyparentidAPIRequest.Get().(*AlitripPlatformDivisionsQuerybyparentidAPIRequest)
}

// ReleaseAlitripPlatformDivisionsQuerybyparentidAPIRequest 将 AlitripPlatformDivisionsQuerybyparentidAPIRequest 放入 sync.Pool
func ReleaseAlitripPlatformDivisionsQuerybyparentidAPIRequest(v *AlitripPlatformDivisionsQuerybyparentidAPIRequest) {
	v.Reset()
	poolAlitripPlatformDivisionsQuerybyparentidAPIRequest.Put(v)
}
