package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelcommoninvoiceremoveAPIRequest 常用发票信息删除接口 API请求
// taobao.xhotel.commoninvoice.remove
//
// 常用发票信息删除接口
type TaobaoxhotelcommoninvoiceremoveAPIRequest struct {
	model.Params
	// 用户名
	_userNick string
	// 发票id
	_invoiceId int64
}

// NewTaobaoxhotelcommoninvoiceremoveRequest 初始化TaobaoxhotelcommoninvoiceremoveAPIRequest对象
func NewTaobaoxhotelcommoninvoiceremoveRequest() *TaobaoxhotelcommoninvoiceremoveAPIRequest {
	return &TaobaoxhotelcommoninvoiceremoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelcommoninvoiceremoveAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.commoninvoice.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelcommoninvoiceremoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelcommoninvoiceremoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 用户名
func (r *TaobaoxhotelcommoninvoiceremoveAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaoxhotelcommoninvoiceremoveAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetInvoiceId is InvoiceId Setter
// 发票id
func (r *TaobaoxhotelcommoninvoiceremoveAPIRequest) SetInvoiceId(_invoiceId int64) error {
	r._invoiceId = _invoiceId
	r.Set("invoice_id", _invoiceId)
	return nil
}

// GetInvoiceId InvoiceId Getter
func (r TaobaoxhotelcommoninvoiceremoveAPIRequest) GetInvoiceId() int64 {
	return r._invoiceId
}
