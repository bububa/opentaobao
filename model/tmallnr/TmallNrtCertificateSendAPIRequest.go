package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtcertificatesendAPIRequest 有价礼包发放电子凭证 API请求
// tmall.nrt.certificate.send
//
// 支持有价礼包发放电子凭证
type TmallnrtcertificatesendAPIRequest struct {
	model.Params
	// 发电子凭证DTO
	_nrtCertificateSendDTO *NrtCertificateSendDto
}

// NewTmallnrtcertificatesendRequest 初始化TmallnrtcertificatesendAPIRequest对象
func NewTmallnrtcertificatesendRequest() *TmallnrtcertificatesendAPIRequest {
	return &TmallnrtcertificatesendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtcertificatesendAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.certificate.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtcertificatesendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtcertificatesendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtCertificateSendDTO is NrtCertificateSendDTO Setter
// 发电子凭证DTO
func (r *TmallnrtcertificatesendAPIRequest) SetNrtCertificateSendDTO(_nrtCertificateSendDTO *NrtCertificateSendDto) error {
	r._nrtCertificateSendDTO = _nrtCertificateSendDTO
	r.Set("nrt_certificate_send_d_t_o", _nrtCertificateSendDTO)
	return nil
}

// GetNrtCertificateSendDTO NrtCertificateSendDTO Getter
func (r TmallnrtcertificatesendAPIRequest) GetNrtCertificateSendDTO() *NrtCertificateSendDto {
	return r._nrtCertificateSendDTO
}
