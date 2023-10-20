package alsc

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmVoucherTemplateListAPIRequest) Reset() {
	r._voucherTemplateOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmVoucherTemplateListAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.voucher.template.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmVoucherTemplateListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmVoucherTemplateListAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlscCrmVoucherTemplateListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmVoucherTemplateListRequest()
	},
}

// GetAlibabaAlscCrmVoucherTemplateListRequest 从 sync.Pool 获取 AlibabaAlscCrmVoucherTemplateListAPIRequest
func GetAlibabaAlscCrmVoucherTemplateListAPIRequest() *AlibabaAlscCrmVoucherTemplateListAPIRequest {
	return poolAlibabaAlscCrmVoucherTemplateListAPIRequest.Get().(*AlibabaAlscCrmVoucherTemplateListAPIRequest)
}

// ReleaseAlibabaAlscCrmVoucherTemplateListAPIRequest 将 AlibabaAlscCrmVoucherTemplateListAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmVoucherTemplateListAPIRequest(v *AlibabaAlscCrmVoucherTemplateListAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmVoucherTemplateListAPIRequest.Put(v)
}
