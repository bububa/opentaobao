package tmallnr

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtCertificateSendAPIRequest 有价礼包发放电子凭证 API请求
// tmall.nrt.certificate.send
//
// 支持有价礼包发放电子凭证
type TmallNrtCertificateSendAPIRequest struct {
	model.Params
	// 发电子凭证DTO
	_nrtCertificateSendDTO *NrtCertificateSendDto
}

// NewTmallNrtCertificateSendRequest 初始化TmallNrtCertificateSendAPIRequest对象
func NewTmallNrtCertificateSendRequest() *TmallNrtCertificateSendAPIRequest {
	return &TmallNrtCertificateSendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtCertificateSendAPIRequest) Reset() {
	r._nrtCertificateSendDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtCertificateSendAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.certificate.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtCertificateSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtCertificateSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtCertificateSendDTO is NrtCertificateSendDTO Setter
// 发电子凭证DTO
func (r *TmallNrtCertificateSendAPIRequest) SetNrtCertificateSendDTO(_nrtCertificateSendDTO *NrtCertificateSendDto) error {
	r._nrtCertificateSendDTO = _nrtCertificateSendDTO
	r.Set("nrt_certificate_send_d_t_o", _nrtCertificateSendDTO)
	return nil
}

// GetNrtCertificateSendDTO NrtCertificateSendDTO Getter
func (r TmallNrtCertificateSendAPIRequest) GetNrtCertificateSendDTO() *NrtCertificateSendDto {
	return r._nrtCertificateSendDTO
}

var poolTmallNrtCertificateSendAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtCertificateSendRequest()
	},
}

// GetTmallNrtCertificateSendRequest 从 sync.Pool 获取 TmallNrtCertificateSendAPIRequest
func GetTmallNrtCertificateSendAPIRequest() *TmallNrtCertificateSendAPIRequest {
	return poolTmallNrtCertificateSendAPIRequest.Get().(*TmallNrtCertificateSendAPIRequest)
}

// ReleaseTmallNrtCertificateSendAPIRequest 将 TmallNrtCertificateSendAPIRequest 放入 sync.Pool
func ReleaseTmallNrtCertificateSendAPIRequest(v *TmallNrtCertificateSendAPIRequest) {
	v.Reset()
	poolTmallNrtCertificateSendAPIRequest.Put(v)
}
