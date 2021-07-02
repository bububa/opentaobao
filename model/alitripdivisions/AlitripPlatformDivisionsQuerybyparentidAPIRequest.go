package alitripdivisions

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPlatformDivisionsQuerybyparentidAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.divisions.querybyparentid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPlatformDivisionsQuerybyparentidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamLong Setter
// 行政区划父id
func (r *AlitripPlatformDivisionsQuerybyparentidAPIRequest) SetParamLong(_paramLong int64) error {
	r._paramLong = _paramLong
	r.Set("param_long", _paramLong)
	return nil
}

// Get ParamLong Getter
func (r AlitripPlatformDivisionsQuerybyparentidAPIRequest) GetParamLong() int64 {
	return r._paramLong
}
