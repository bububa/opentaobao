package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmvouchertemplatelistAPIRequest 获取优惠券模版列表 API请求
// alibaba.alsc.crm.voucher.template.list
//
// 获取优惠券模版列表
type AlibabaalsccrmvouchertemplatelistAPIRequest struct {
	model.Params
	// 获取优惠模版规则请求参数
	_voucherTemplateOpenReq *VoucherTemplateOpenReq
}

// NewAlibabaalsccrmvouchertemplatelistRequest 初始化AlibabaalsccrmvouchertemplatelistAPIRequest对象
func NewAlibabaalsccrmvouchertemplatelistRequest() *AlibabaalsccrmvouchertemplatelistAPIRequest {
	return &AlibabaalsccrmvouchertemplatelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmvouchertemplatelistAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.voucher.template.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmvouchertemplatelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmvouchertemplatelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVoucherTemplateOpenReq is VoucherTemplateOpenReq Setter
// 获取优惠模版规则请求参数
func (r *AlibabaalsccrmvouchertemplatelistAPIRequest) SetVoucherTemplateOpenReq(_voucherTemplateOpenReq *VoucherTemplateOpenReq) error {
	r._voucherTemplateOpenReq = _voucherTemplateOpenReq
	r.Set("voucher_template_open_req", _voucherTemplateOpenReq)
	return nil
}

// GetVoucherTemplateOpenReq VoucherTemplateOpenReq Getter
func (r AlibabaalsccrmvouchertemplatelistAPIRequest) GetVoucherTemplateOpenReq() *VoucherTemplateOpenReq {
	return r._voucherTemplateOpenReq
}
