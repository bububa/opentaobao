package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelitemselectionsellerstatexposureAPIRequest 选品曝光趋势 API请求
// taobao.xhotel.item.selection.seller.stat.exposure
//
// 用于提供给商家获取选品曝光趋势
type TaobaoxhotelitemselectionsellerstatexposureAPIRequest struct {
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

// NewTaobaoxhotelitemselectionsellerstatexposureRequest 初始化TaobaoxhotelitemselectionsellerstatexposureAPIRequest对象
func NewTaobaoxhotelitemselectionsellerstatexposureRequest() *TaobaoxhotelitemselectionsellerstatexposureAPIRequest {
	return &TaobaoxhotelitemselectionsellerstatexposureAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelitemselectionsellerstatexposureAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.item.selection.seller.stat.exposure"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelitemselectionsellerstatexposureAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelitemselectionsellerstatexposureAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDate is Date Setter
// 日期 默认为昨天
func (r *TaobaoxhotelitemselectionsellerstatexposureAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r TaobaoxhotelitemselectionsellerstatexposureAPIRequest) GetDate() string {
	return r._date
}

// SetHid is Hid Setter
// hid  默认为all
func (r *TaobaoxhotelitemselectionsellerstatexposureAPIRequest) SetHid(_hid string) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoxhotelitemselectionsellerstatexposureAPIRequest) GetHid() string {
	return r._hid
}

// SetVendor is Vendor Setter
// 默认为all
func (r *TaobaoxhotelitemselectionsellerstatexposureAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelitemselectionsellerstatexposureAPIRequest) GetVendor() string {
	return r._vendor
}

// SetSupplier is Supplier Setter
// 默认为all
func (r *TaobaoxhotelitemselectionsellerstatexposureAPIRequest) SetSupplier(_supplier string) error {
	r._supplier = _supplier
	r.Set("supplier", _supplier)
	return nil
}

// GetSupplier Supplier Getter
func (r TaobaoxhotelitemselectionsellerstatexposureAPIRequest) GetSupplier() string {
	return r._supplier
}

// SetOutHid is OutHid Setter
// 酒店编码
func (r *TaobaoxhotelitemselectionsellerstatexposureAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoxhotelitemselectionsellerstatexposureAPIRequest) GetOutHid() string {
	return r._outHid
}
