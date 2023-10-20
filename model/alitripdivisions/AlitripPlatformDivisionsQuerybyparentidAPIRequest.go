package alitripdivisions

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripplatformdivisionsquerybyparentidAPIRequest 根据父节点id查询下级行政区划数据 API请求
// alitrip.platform.divisions.querybyparentid
//
// 根据行政区划id查询下一层级行政区划数据
type AlitripplatformdivisionsquerybyparentidAPIRequest struct {
	model.Params
	// 行政区划父id
	_paramLong int64
}

// NewAlitripplatformdivisionsquerybyparentidRequest 初始化AlitripplatformdivisionsquerybyparentidAPIRequest对象
func NewAlitripplatformdivisionsquerybyparentidRequest() *AlitripplatformdivisionsquerybyparentidAPIRequest {
	return &AlitripplatformdivisionsquerybyparentidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripplatformdivisionsquerybyparentidAPIRequest) GetApiMethodName() string {
	return "alitrip.platform.divisions.querybyparentid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripplatformdivisionsquerybyparentidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripplatformdivisionsquerybyparentidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamLong is ParamLong Setter
// 行政区划父id
func (r *AlitripplatformdivisionsquerybyparentidAPIRequest) SetParamLong(_paramLong int64) error {
	r._paramLong = _paramLong
	r.Set("param_long", _paramLong)
	return nil
}

// GetParamLong ParamLong Getter
func (r AlitripplatformdivisionsquerybyparentidAPIRequest) GetParamLong() int64 {
	return r._paramLong
}
