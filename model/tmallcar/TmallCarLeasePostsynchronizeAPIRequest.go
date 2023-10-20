package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarleasepostsynchronizeAPIRequest 天猫开新车租后信息同步 API请求
// tmall.car.lease.postsynchronize
//
// 商家同步天猫开新车租后方案
type TmallcarleasepostsynchronizeAPIRequest struct {
	model.Params
	// 租后方案信息
	_schemeDto *CarLeasePostSchemeSynchronizeDto
}

// NewTmallcarleasepostsynchronizeRequest 初始化TmallcarleasepostsynchronizeAPIRequest对象
func NewTmallcarleasepostsynchronizeRequest() *TmallcarleasepostsynchronizeAPIRequest {
	return &TmallcarleasepostsynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarleasepostsynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.postsynchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarleasepostsynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarleasepostsynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemeDto is SchemeDto Setter
// 租后方案信息
func (r *TmallcarleasepostsynchronizeAPIRequest) SetSchemeDto(_schemeDto *CarLeasePostSchemeSynchronizeDto) error {
	r._schemeDto = _schemeDto
	r.Set("scheme_dto", _schemeDto)
	return nil
}

// GetSchemeDto SchemeDto Getter
func (r TmallcarleasepostsynchronizeAPIRequest) GetSchemeDto() *CarLeasePostSchemeSynchronizeDto {
	return r._schemeDto
}
