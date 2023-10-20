package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelItemSelectionSellerStatExposureAPIRequest 选品曝光趋势 API请求
// taobao.xhotel.item.selection.seller.stat.exposure
//
// 用于提供给商家获取选品曝光趋势
type TaobaoXhotelItemSelectionSellerStatExposureAPIRequest struct {
	model.Params
	// 日期 默认为昨天
	_date string
	// hid  默认为all
	_hid string
	// 默认为all
	_vendor string
	// 默认为all
	_supplier string
	// 酒店编码
	_outHid string
}

// NewTaobaoXhotelItemSelectionSellerStatExposureRequest 初始化TaobaoXhotelItemSelectionSellerStatExposureAPIRequest对象
func NewTaobaoXhotelItemSelectionSellerStatExposureRequest() *TaobaoXhotelItemSelectionSellerStatExposureAPIRequest {
	return &TaobaoXhotelItemSelectionSellerStatExposureAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) Reset() {
	r._date = ""
	r._hid = ""
	r._vendor = ""
	r._supplier = ""
	r._outHid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.item.selection.seller.stat.exposure"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDate is Date Setter
// 日期 默认为昨天
func (r *TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) GetDate() string {
	return r._date
}

// SetHid is Hid Setter
// hid  默认为all
func (r *TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) SetHid(_hid string) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) GetHid() string {
	return r._hid
}

// SetVendor is Vendor Setter
// 默认为all
func (r *TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) GetVendor() string {
	return r._vendor
}

// SetSupplier is Supplier Setter
// 默认为all
func (r *TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) SetSupplier(_supplier string) error {
	r._supplier = _supplier
	r.Set("supplier", _supplier)
	return nil
}

// GetSupplier Supplier Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) GetSupplier() string {
	return r._supplier
}

// SetOutHid is OutHid Setter
// 酒店编码
func (r *TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) GetOutHid() string {
	return r._outHid
}

var poolTaobaoXhotelItemSelectionSellerStatExposureAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelItemSelectionSellerStatExposureRequest()
	},
}

// GetTaobaoXhotelItemSelectionSellerStatExposureRequest 从 sync.Pool 获取 TaobaoXhotelItemSelectionSellerStatExposureAPIRequest
func GetTaobaoXhotelItemSelectionSellerStatExposureAPIRequest() *TaobaoXhotelItemSelectionSellerStatExposureAPIRequest {
	return poolTaobaoXhotelItemSelectionSellerStatExposureAPIRequest.Get().(*TaobaoXhotelItemSelectionSellerStatExposureAPIRequest)
}

// ReleaseTaobaoXhotelItemSelectionSellerStatExposureAPIRequest 将 TaobaoXhotelItemSelectionSellerStatExposureAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelItemSelectionSellerStatExposureAPIRequest(v *TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) {
	v.Reset()
	poolTaobaoXhotelItemSelectionSellerStatExposureAPIRequest.Put(v)
}
