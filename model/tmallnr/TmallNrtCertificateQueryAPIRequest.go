package tmallnr

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtCertificateQueryAPIRequest 批量查询电子凭证信息 API请求
// tmall.nrt.certificate.query
//
// 批量查询电子凭证信息
type TmallNrtCertificateQueryAPIRequest struct {
	model.Params
	// 查询电子凭证实例DTO
	_nrtCertificateInstanceQueryDTO *NrtCertificateInstanceQueryDto
}

// NewTmallNrtCertificateQueryRequest 初始化TmallNrtCertificateQueryAPIRequest对象
func NewTmallNrtCertificateQueryRequest() *TmallNrtCertificateQueryAPIRequest {
	return &TmallNrtCertificateQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtCertificateQueryAPIRequest) Reset() {
	r._nrtCertificateInstanceQueryDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtCertificateQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.certificate.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtCertificateQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtCertificateQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtCertificateInstanceQueryDTO is NrtCertificateInstanceQueryDTO Setter
// 查询电子凭证实例DTO
func (r *TmallNrtCertificateQueryAPIRequest) SetNrtCertificateInstanceQueryDTO(_nrtCertificateInstanceQueryDTO *NrtCertificateInstanceQueryDto) error {
	r._nrtCertificateInstanceQueryDTO = _nrtCertificateInstanceQueryDTO
	r.Set("nrt_certificate_instance_query_d_t_o", _nrtCertificateInstanceQueryDTO)
	return nil
}

// GetNrtCertificateInstanceQueryDTO NrtCertificateInstanceQueryDTO Getter
func (r TmallNrtCertificateQueryAPIRequest) GetNrtCertificateInstanceQueryDTO() *NrtCertificateInstanceQueryDto {
	return r._nrtCertificateInstanceQueryDTO
}

var poolTmallNrtCertificateQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtCertificateQueryRequest()
	},
}

// GetTmallNrtCertificateQueryRequest 从 sync.Pool 获取 TmallNrtCertificateQueryAPIRequest
func GetTmallNrtCertificateQueryAPIRequest() *TmallNrtCertificateQueryAPIRequest {
	return poolTmallNrtCertificateQueryAPIRequest.Get().(*TmallNrtCertificateQueryAPIRequest)
}

// ReleaseTmallNrtCertificateQueryAPIRequest 将 TmallNrtCertificateQueryAPIRequest 放入 sync.Pool
func ReleaseTmallNrtCertificateQueryAPIRequest(v *TmallNrtCertificateQueryAPIRequest) {
	v.Reset()
	poolTmallNrtCertificateQueryAPIRequest.Put(v)
}
