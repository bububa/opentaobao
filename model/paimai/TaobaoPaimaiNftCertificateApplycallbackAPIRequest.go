package paimai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPaimaiNftCertificateApplycallbackAPIRequest 数字藏品版权证书申请结果回调 API请求
// taobao.paimai.nft.certificate.applycallback
//
// 数字藏品版权证书申请结果回调
type TaobaoPaimaiNftCertificateApplycallbackAPIRequest struct {
	model.Params
	// dto
	_nftCertificateApplyCallbackDto *NftCertificateApplyCallbackDto
}

// NewTaobaoPaimaiNftCertificateApplycallbackRequest 初始化TaobaoPaimaiNftCertificateApplycallbackAPIRequest对象
func NewTaobaoPaimaiNftCertificateApplycallbackRequest() *TaobaoPaimaiNftCertificateApplycallbackAPIRequest {
	return &TaobaoPaimaiNftCertificateApplycallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPaimaiNftCertificateApplycallbackAPIRequest) GetApiMethodName() string {
	return "taobao.paimai.nft.certificate.applycallback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPaimaiNftCertificateApplycallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPaimaiNftCertificateApplycallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNftCertificateApplyCallbackDto is NftCertificateApplyCallbackDto Setter
// dto
func (r *TaobaoPaimaiNftCertificateApplycallbackAPIRequest) SetNftCertificateApplyCallbackDto(_nftCertificateApplyCallbackDto *NftCertificateApplyCallbackDto) error {
	r._nftCertificateApplyCallbackDto = _nftCertificateApplyCallbackDto
	r.Set("nft_certificate_apply_callback_dto", _nftCertificateApplyCallbackDto)
	return nil
}

// GetNftCertificateApplyCallbackDto NftCertificateApplyCallbackDto Getter
func (r TaobaoPaimaiNftCertificateApplycallbackAPIRequest) GetNftCertificateApplyCallbackDto() *NftCertificateApplyCallbackDto {
	return r._nftCertificateApplyCallbackDto
}
