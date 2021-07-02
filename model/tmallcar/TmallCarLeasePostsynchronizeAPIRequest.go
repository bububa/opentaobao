package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeasePostsynchronizeAPIRequest 天猫开新车租后信息同步 API请求
// tmall.car.lease.postsynchronize
//
// 商家同步天猫开新车租后方案
type TmallCarLeasePostsynchronizeAPIRequest struct {
	model.Params
	// 租后方案信息
	_schemeDto *CarLeasePostSchemeSynchronizeDto
}

// NewTmallCarLeasePostsynchronizeRequest 初始化TmallCarLeasePostsynchronizeAPIRequest对象
func NewTmallCarLeasePostsynchronizeRequest() *TmallCarLeasePostsynchronizeAPIRequest {
	return &TmallCarLeasePostsynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeasePostsynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.postsynchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeasePostsynchronizeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSchemeDto is SchemeDto Setter
// 租后方案信息
func (r *TmallCarLeasePostsynchronizeAPIRequest) SetSchemeDto(_schemeDto *CarLeasePostSchemeSynchronizeDto) error {
	r._schemeDto = _schemeDto
	r.Set("scheme_dto", _schemeDto)
	return nil
}

// GetSchemeDto SchemeDto Getter
func (r TmallCarLeasePostsynchronizeAPIRequest) GetSchemeDto() *CarLeasePostSchemeSynchronizeDto {
	return r._schemeDto
}
