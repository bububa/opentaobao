package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelitemselectionsellerstatsummaryAPIRequest 商家数据-选品整体概况 API请求
// taobao.xhotel.item.selection.seller.stat.summary
//
// 商家数据-选品整体概况
type TaobaoxhotelitemselectionsellerstatsummaryAPIRequest struct {
	model.Params
	// 日期  默认为昨天
	_date string
	// hid  默认为all
	_hid string
	// vendor 默认为all
	_vendor string
	// supplier 默认为all
	_supplier string
	// 外部酒店编码
	_outHid string
}

// NewTaobaoxhotelitemselectionsellerstatsummaryRequest 初始化TaobaoxhotelitemselectionsellerstatsummaryAPIRequest对象
func NewTaobaoxhotelitemselectionsellerstatsummaryRequest() *TaobaoxhotelitemselectionsellerstatsummaryAPIRequest {
	return &TaobaoxhotelitemselectionsellerstatsummaryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelitemselectionsellerstatsummaryAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.item.selection.seller.stat.summary"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelitemselectionsellerstatsummaryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelitemselectionsellerstatsummaryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDate is Date Setter
// 日期  默认为昨天
func (r *TaobaoxhotelitemselectionsellerstatsummaryAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r TaobaoxhotelitemselectionsellerstatsummaryAPIRequest) GetDate() string {
	return r._date
}

// SetHid is Hid Setter
// hid  默认为all
func (r *TaobaoxhotelitemselectionsellerstatsummaryAPIRequest) SetHid(_hid string) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoxhotelitemselectionsellerstatsummaryAPIRequest) GetHid() string {
	return r._hid
}

// SetVendor is Vendor Setter
// vendor 默认为all
func (r *TaobaoxhotelitemselectionsellerstatsummaryAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelitemselectionsellerstatsummaryAPIRequest) GetVendor() string {
	return r._vendor
}

// SetSupplier is Supplier Setter
// supplier 默认为all
func (r *TaobaoxhotelitemselectionsellerstatsummaryAPIRequest) SetSupplier(_supplier string) error {
	r._supplier = _supplier
	r.Set("supplier", _supplier)
	return nil
}

// GetSupplier Supplier Getter
func (r TaobaoxhotelitemselectionsellerstatsummaryAPIRequest) GetSupplier() string {
	return r._supplier
}

// SetOutHid is OutHid Setter
// 外部酒店编码
func (r *TaobaoxhotelitemselectionsellerstatsummaryAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoxhotelitemselectionsellerstatsummaryAPIRequest) GetOutHid() string {
	return r._outHid
}
