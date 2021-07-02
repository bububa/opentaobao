package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest 供应链选品热销标准酒店覆盖情况 API请求
// taobao.xhotel.item.selection.seller.stat.hotshid
//
// 供应链选品热销标准酒店覆盖情况
type TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest struct {
	model.Params
	// 日期  默认为昨天
	_date string
	// 酒店id  默认all
	_hid string
	// vendor  默认all
	_vendor string
	// supplier  默认all
	_supplier string
	// 酒店编码
	_outHid string
}

// NewTaobaoXhotelItemSelectionSellerStatHotshidRequest 初始化TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest对象
func NewTaobaoXhotelItemSelectionSellerStatHotshidRequest() *TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest {
	return &TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.item.selection.seller.stat.hotshid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Date Setter
// 日期  默认为昨天
func (r *TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// Get Date Getter
func (r TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest) GetDate() string {
	return r._date
}

// Set is Hid Setter
// 酒店id  默认all
func (r *TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest) SetHid(_hid string) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// Get Hid Getter
func (r TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest) GetHid() string {
	return r._hid
}

// Set is Vendor Setter
// vendor  默认all
func (r *TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// Get Vendor Getter
func (r TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest) GetVendor() string {
	return r._vendor
}

// Set is Supplier Setter
// supplier  默认all
func (r *TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest) SetSupplier(_supplier string) error {
	r._supplier = _supplier
	r.Set("supplier", _supplier)
	return nil
}

// Get Supplier Getter
func (r TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest) GetSupplier() string {
	return r._supplier
}

// Set is OutHid Setter
// 酒店编码
func (r *TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// Get OutHid Getter
func (r TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest) GetOutHid() string {
	return r._outHid
}
