package xhotelonlineorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelCommoninvoiceRemoveAPIRequest 常用发票信息删除接口 API请求
// taobao.xhotel.commoninvoice.remove
//
// 常用发票信息删除接口
type TaobaoXhotelCommoninvoiceRemoveAPIRequest struct {
	model.Params
	// 用户名
	_userNick string
	// 发票id
	_invoiceId int64
}

// NewTaobaoXhotelCommoninvoiceRemoveRequest 初始化TaobaoXhotelCommoninvoiceRemoveAPIRequest对象
func NewTaobaoXhotelCommoninvoiceRemoveRequest() *TaobaoXhotelCommoninvoiceRemoveAPIRequest {
	return &TaobaoXhotelCommoninvoiceRemoveAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelCommoninvoiceRemoveAPIRequest) Reset() {
	r._userNick = ""
	r._invoiceId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelCommoninvoiceRemoveAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.commoninvoice.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelCommoninvoiceRemoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelCommoninvoiceRemoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 用户名
func (r *TaobaoXhotelCommoninvoiceRemoveAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaoXhotelCommoninvoiceRemoveAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetInvoiceId is InvoiceId Setter
// 发票id
func (r *TaobaoXhotelCommoninvoiceRemoveAPIRequest) SetInvoiceId(_invoiceId int64) error {
	r._invoiceId = _invoiceId
	r.Set("invoice_id", _invoiceId)
	return nil
}

// GetInvoiceId InvoiceId Getter
func (r TaobaoXhotelCommoninvoiceRemoveAPIRequest) GetInvoiceId() int64 {
	return r._invoiceId
}

var poolTaobaoXhotelCommoninvoiceRemoveAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelCommoninvoiceRemoveRequest()
	},
}

// GetTaobaoXhotelCommoninvoiceRemoveRequest 从 sync.Pool 获取 TaobaoXhotelCommoninvoiceRemoveAPIRequest
func GetTaobaoXhotelCommoninvoiceRemoveAPIRequest() *TaobaoXhotelCommoninvoiceRemoveAPIRequest {
	return poolTaobaoXhotelCommoninvoiceRemoveAPIRequest.Get().(*TaobaoXhotelCommoninvoiceRemoveAPIRequest)
}

// ReleaseTaobaoXhotelCommoninvoiceRemoveAPIRequest 将 TaobaoXhotelCommoninvoiceRemoveAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelCommoninvoiceRemoveAPIRequest(v *TaobaoXhotelCommoninvoiceRemoveAPIRequest) {
	v.Reset()
	poolTaobaoXhotelCommoninvoiceRemoveAPIRequest.Put(v)
}
