package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelcommoninvoiceupdateAPIRequest 常用发票信息更新接口 API请求
// taobao.xhotel.commoninvoice.update
//
// 常用发票信息更新接口(根据用户id,发票抬头和发票属性或发票id进行更新,没有则添加)
type TaobaoxhotelcommoninvoiceupdateAPIRequest struct {
	model.Params
	// 无
	_commonInvoiceInfoParam *CommonInvoiceInfo
}

// NewTaobaoxhotelcommoninvoiceupdateRequest 初始化TaobaoxhotelcommoninvoiceupdateAPIRequest对象
func NewTaobaoxhotelcommoninvoiceupdateRequest() *TaobaoxhotelcommoninvoiceupdateAPIRequest {
	return &TaobaoxhotelcommoninvoiceupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelcommoninvoiceupdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.commoninvoice.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelcommoninvoiceupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelcommoninvoiceupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommonInvoiceInfoParam is CommonInvoiceInfoParam Setter
// 无
func (r *TaobaoxhotelcommoninvoiceupdateAPIRequest) SetCommonInvoiceInfoParam(_commonInvoiceInfoParam *CommonInvoiceInfo) error {
	r._commonInvoiceInfoParam = _commonInvoiceInfoParam
	r.Set("common_invoice_info_param", _commonInvoiceInfoParam)
	return nil
}

// GetCommonInvoiceInfoParam CommonInvoiceInfoParam Getter
func (r TaobaoxhotelcommoninvoiceupdateAPIRequest) GetCommonInvoiceInfoParam() *CommonInvoiceInfo {
	return r._commonInvoiceInfoParam
}
