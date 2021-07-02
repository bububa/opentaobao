package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest 商家数据-选品整体概况 API请求
// taobao.xhotel.item.selection.seller.stat.summary
//
// 商家数据-选品整体概况
type TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest struct {
	model.Params
	// vendor 默认为all
	_vendor string
	// 日期  默认为昨天
	_date string
	// hid  默认为all
	_hid string
	// supplier 默认为all
	_supplier string
	// 外部酒店编码
	_outHid string
}

// NewTaobaoXhotelItemSelectionSellerStatSummaryRequest 初始化TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest对象
func NewTaobaoXhotelItemSelectionSellerStatSummaryRequest() *TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest {
	return &TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.item.selection.seller.stat.summary"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetVendor is Vendor Setter
// vendor 默认为all
func (r *TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest) GetVendor() string {
	return r._vendor
}

// SetDate is Date Setter
// 日期  默认为昨天
func (r *TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest) GetDate() string {
	return r._date
}

// SetHid is Hid Setter
// hid  默认为all
func (r *TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest) SetHid(_hid string) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest) GetHid() string {
	return r._hid
}

// SetSupplier is Supplier Setter
// supplier 默认为all
func (r *TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest) SetSupplier(_supplier string) error {
	r._supplier = _supplier
	r.Set("supplier", _supplier)
	return nil
}

// GetSupplier Supplier Getter
func (r TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest) GetSupplier() string {
	return r._supplier
}

// SetOutHid is OutHid Setter
// 外部酒店编码
func (r *TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoXhotelItemSelectionSellerStatSummaryAPIRequest) GetOutHid() string {
	return r._outHid
}
