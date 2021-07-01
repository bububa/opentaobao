package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelItemSelectionSellerStatExposureAPIRequest
选品曝光趋势 API请求
taobao.xhotel.item.selection.seller.stat.exposure

用于提供给商家获取选品曝光趋势 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.item.selection.seller.stat.exposure"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Date Setter
// 日期 默认为昨天
func (r *TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// Get Date Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) GetDate() string {
	return r._date
}

// Set is Hid Setter
// hid  默认为all
func (r *TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) SetHid(_hid string) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// Get Hid Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) GetHid() string {
	return r._hid
}

// Set is Vendor Setter
// 默认为all
func (r *TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// Get Vendor Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) GetVendor() string {
	return r._vendor
}

// Set is Supplier Setter
// 默认为all
func (r *TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) SetSupplier(_supplier string) error {
	r._supplier = _supplier
	r.Set("supplier", _supplier)
	return nil
}

// Get Supplier Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) GetSupplier() string {
	return r._supplier
}

// Set is OutHid Setter
// 酒店编码
func (r *TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// Get OutHid Getter
func (r TaobaoXhotelItemSelectionSellerStatExposureAPIRequest) GetOutHid() string {
	return r._outHid
}
