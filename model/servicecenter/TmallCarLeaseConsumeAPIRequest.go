package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseConsumeAPIRequest 汽车租赁核销 API请求
// tmall.car.lease.consume
//
// 租赁公司回传信息，核销
type TmallCarLeaseConsumeAPIRequest struct {
	model.Params
	// 核销请求
	_cosumeCodeReqDTO *CosumeCodeReqDto
}

// NewTmallCarLeaseConsumeRequest 初始化TmallCarLeaseConsumeAPIRequest对象
func NewTmallCarLeaseConsumeRequest() *TmallCarLeaseConsumeAPIRequest {
	return &TmallCarLeaseConsumeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarLeaseConsumeAPIRequest) Reset() {
	r._cosumeCodeReqDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseConsumeAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseConsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarLeaseConsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCosumeCodeReqDTO is CosumeCodeReqDTO Setter
// 核销请求
func (r *TmallCarLeaseConsumeAPIRequest) SetCosumeCodeReqDTO(_cosumeCodeReqDTO *CosumeCodeReqDto) error {
	r._cosumeCodeReqDTO = _cosumeCodeReqDTO
	r.Set("cosume_code_req_d_t_o", _cosumeCodeReqDTO)
	return nil
}

// GetCosumeCodeReqDTO CosumeCodeReqDTO Getter
func (r TmallCarLeaseConsumeAPIRequest) GetCosumeCodeReqDTO() *CosumeCodeReqDto {
	return r._cosumeCodeReqDTO
}

var poolTmallCarLeaseConsumeAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarLeaseConsumeRequest()
	},
}

// GetTmallCarLeaseConsumeRequest 从 sync.Pool 获取 TmallCarLeaseConsumeAPIRequest
func GetTmallCarLeaseConsumeAPIRequest() *TmallCarLeaseConsumeAPIRequest {
	return poolTmallCarLeaseConsumeAPIRequest.Get().(*TmallCarLeaseConsumeAPIRequest)
}

// ReleaseTmallCarLeaseConsumeAPIRequest 将 TmallCarLeaseConsumeAPIRequest 放入 sync.Pool
func ReleaseTmallCarLeaseConsumeAPIRequest(v *TmallCarLeaseConsumeAPIRequest) {
	v.Reset()
	poolTmallCarLeaseConsumeAPIRequest.Put(v)
}
