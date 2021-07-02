package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelCommoninvoiceRemoveAPIRequest 常用发票信息删除接口 API请求
// taobao.xhotel.commoninvoice.remove
//
// 常用发票信息删除接口
type TaobaoXhotelCommoninvoiceRemoveAPIRequest struct {
	model.Params
	// 发票id
	_invoiceId int64
	// 用户名
	_userNick string
}

// NewTaobaoXhotelCommoninvoiceRemoveRequest 初始化TaobaoXhotelCommoninvoiceRemoveAPIRequest对象
func NewTaobaoXhotelCommoninvoiceRemoveRequest() *TaobaoXhotelCommoninvoiceRemoveAPIRequest {
	return &TaobaoXhotelCommoninvoiceRemoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelCommoninvoiceRemoveAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.commoninvoice.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelCommoninvoiceRemoveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is InvoiceId Setter
// 发票id
func (r *TaobaoXhotelCommoninvoiceRemoveAPIRequest) SetInvoiceId(_invoiceId int64) error {
	r._invoiceId = _invoiceId
	r.Set("invoice_id", _invoiceId)
	return nil
}

// Get InvoiceId Getter
func (r TaobaoXhotelCommoninvoiceRemoveAPIRequest) GetInvoiceId() int64 {
	return r._invoiceId
}

// Set is UserNick Setter
// 用户名
func (r *TaobaoXhotelCommoninvoiceRemoveAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// Get UserNick Getter
func (r TaobaoXhotelCommoninvoiceRemoveAPIRequest) GetUserNick() string {
	return r._userNick
}
