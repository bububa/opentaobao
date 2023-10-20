package xhotelonlineorder

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelCommoninvoiceUpdateAPIRequest) Reset() {
	r._commonInvoiceInfoParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelCommoninvoiceUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.commoninvoice.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelCommoninvoiceUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelCommoninvoiceUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommonInvoiceInfoParam is CommonInvoiceInfoParam Setter
// 无
func (r *TaobaoXhotelCommoninvoiceUpdateAPIRequest) SetCommonInvoiceInfoParam(_commonInvoiceInfoParam *CommonInvoiceInfo) error {
	r._commonInvoiceInfoParam = _commonInvoiceInfoParam
	r.Set("common_invoice_info_param", _commonInvoiceInfoParam)
	return nil
}

// GetCommonInvoiceInfoParam CommonInvoiceInfoParam Getter
func (r TaobaoXhotelCommoninvoiceUpdateAPIRequest) GetCommonInvoiceInfoParam() *CommonInvoiceInfo {
	return r._commonInvoiceInfoParam
}

var poolTaobaoXhotelCommoninvoiceUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelCommoninvoiceUpdateRequest()
	},
}

// GetTaobaoXhotelCommoninvoiceUpdateRequest 从 sync.Pool 获取 TaobaoXhotelCommoninvoiceUpdateAPIRequest
func GetTaobaoXhotelCommoninvoiceUpdateAPIRequest() *TaobaoXhotelCommoninvoiceUpdateAPIRequest {
	return poolTaobaoXhotelCommoninvoiceUpdateAPIRequest.Get().(*TaobaoXhotelCommoninvoiceUpdateAPIRequest)
}

// ReleaseTaobaoXhotelCommoninvoiceUpdateAPIRequest 将 TaobaoXhotelCommoninvoiceUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelCommoninvoiceUpdateAPIRequest(v *TaobaoXhotelCommoninvoiceUpdateAPIRequest) {
	v.Reset()
	poolTaobaoXhotelCommoninvoiceUpdateAPIRequest.Put(v)
}
