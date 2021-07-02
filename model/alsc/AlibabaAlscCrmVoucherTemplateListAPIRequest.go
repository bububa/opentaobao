package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmVoucherTemplateListAPIRequest 获取优惠券模版列表 API请求
// alibaba.alsc.crm.voucher.template.list
//
// 获取优惠券模版列表
type AlibabaAlscCrmVoucherTemplateListAPIRequest struct {
	model.Params
	// 获取优惠模版规则请求参数
	_voucherTemplateOpenReq *VoucherTemplateOpenReq
}

// NewAlibabaAlscCrmVoucherTemplateListRequest 初始化AlibabaAlscCrmVoucherTemplateListAPIRequest对象
func NewAlibabaAlscCrmVoucherTemplateListRequest() *AlibabaAlscCrmVoucherTemplateListAPIRequest {
	return &AlibabaAlscCrmVoucherTemplateListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmVoucherTemplateListAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.voucher.template.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmVoucherTemplateListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetVoucherTemplateOpenReq is VoucherTemplateOpenReq Setter
// 获取优惠模版规则请求参数
func (r *AlibabaAlscCrmVoucherTemplateListAPIRequest) SetVoucherTemplateOpenReq(_voucherTemplateOpenReq *VoucherTemplateOpenReq) error {
	r._voucherTemplateOpenReq = _voucherTemplateOpenReq
	r.Set("voucher_template_open_req", _voucherTemplateOpenReq)
	return nil
}

// GetVoucherTemplateOpenReq VoucherTemplateOpenReq Getter
func (r AlibabaAlscCrmVoucherTemplateListAPIRequest) GetVoucherTemplateOpenReq() *VoucherTemplateOpenReq {
	return r._voucherTemplateOpenReq
}
