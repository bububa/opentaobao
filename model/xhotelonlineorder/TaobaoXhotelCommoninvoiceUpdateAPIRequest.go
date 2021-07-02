package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelCommoninvoiceUpdateAPIRequest 常用发票信息更新接口 API请求
// taobao.xhotel.commoninvoice.update
//
// 常用发票信息更新接口(根据用户id,发票抬头和发票属性或发票id进行更新,没有则添加)
type TaobaoXhotelCommoninvoiceUpdateAPIRequest struct {
	model.Params
	// 无
	_commonInvoiceInfoParam *CommonInvoiceInfo
}

// NewTaobaoXhotelCommoninvoiceUpdateRequest 初始化TaobaoXhotelCommoninvoiceUpdateAPIRequest对象
func NewTaobaoXhotelCommoninvoiceUpdateRequest() *TaobaoXhotelCommoninvoiceUpdateAPIRequest {
	return &TaobaoXhotelCommoninvoiceUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelCommoninvoiceUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.commoninvoice.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelCommoninvoiceUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CommonInvoiceInfoParam Setter
// 无
func (r *TaobaoXhotelCommoninvoiceUpdateAPIRequest) SetCommonInvoiceInfoParam(_commonInvoiceInfoParam *CommonInvoiceInfo) error {
	r._commonInvoiceInfoParam = _commonInvoiceInfoParam
	r.Set("common_invoice_info_param", _commonInvoiceInfoParam)
	return nil
}

// Get CommonInvoiceInfoParam Getter
func (r TaobaoXhotelCommoninvoiceUpdateAPIRequest) GetCommonInvoiceInfoParam() *CommonInvoiceInfo {
	return r._commonInvoiceInfoParam
}
