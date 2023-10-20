package paimai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopaimainftcertificateapplycallbackAPIRequest 数字藏品版权证书申请结果回调 API请求
// taobao.paimai.nft.certificate.applycallback
//
// 数字藏品版权证书申请结果回调
type TaobaopaimainftcertificateapplycallbackAPIRequest struct {
	model.Params
	// dto
	_nftCertificateApplyCallbackDto *NftCertificateApplyCallbackDto
}

// NewTaobaopaimainftcertificateapplycallbackRequest 初始化TaobaopaimainftcertificateapplycallbackAPIRequest对象
func NewTaobaopaimainftcertificateapplycallbackRequest() *TaobaopaimainftcertificateapplycallbackAPIRequest {
	return &TaobaopaimainftcertificateapplycallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopaimainftcertificateapplycallbackAPIRequest) GetApiMethodName() string {
	return "taobao.paimai.nft.certificate.applycallback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopaimainftcertificateapplycallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopaimainftcertificateapplycallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNftCertificateApplyCallbackDto is NftCertificateApplyCallbackDto Setter
// dto
func (r *TaobaopaimainftcertificateapplycallbackAPIRequest) SetNftCertificateApplyCallbackDto(_nftCertificateApplyCallbackDto *NftCertificateApplyCallbackDto) error {
	r._nftCertificateApplyCallbackDto = _nftCertificateApplyCallbackDto
	r.Set("nft_certificate_apply_callback_dto", _nftCertificateApplyCallbackDto)
	return nil
}

// GetNftCertificateApplyCallbackDto NftCertificateApplyCallbackDto Getter
func (r TaobaopaimainftcertificateapplycallbackAPIRequest) GetNftCertificateApplyCallbackDto() *NftCertificateApplyCallbackDto {
	return r._nftCertificateApplyCallbackDto
}
