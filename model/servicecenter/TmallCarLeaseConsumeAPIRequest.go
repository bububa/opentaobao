package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarleaseconsumeAPIRequest 汽车租赁核销 API请求
// tmall.car.lease.consume
//
// 租赁公司回传信息，核销
type TmallcarleaseconsumeAPIRequest struct {
	model.Params
	// 核销请求
	_cosumeCodeReqDTO *CosumeCodeReqDto
}

// NewTmallcarleaseconsumeRequest 初始化TmallcarleaseconsumeAPIRequest对象
func NewTmallcarleaseconsumeRequest() *TmallcarleaseconsumeAPIRequest {
	return &TmallcarleaseconsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarleaseconsumeAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarleaseconsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarleaseconsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCosumeCodeReqDTO is CosumeCodeReqDTO Setter
// 核销请求
func (r *TmallcarleaseconsumeAPIRequest) SetCosumeCodeReqDTO(_cosumeCodeReqDTO *CosumeCodeReqDto) error {
	r._cosumeCodeReqDTO = _cosumeCodeReqDTO
	r.Set("cosume_code_req_d_t_o", _cosumeCodeReqDTO)
	return nil
}

// GetCosumeCodeReqDTO CosumeCodeReqDTO Getter
func (r TmallcarleaseconsumeAPIRequest) GetCosumeCodeReqDTO() *CosumeCodeReqDto {
	return r._cosumeCodeReqDTO
}
