package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelorderhotelsignqueryAPIRequest 获取直连酒店（客栈）签约上线进度信息 API请求
// taobao.xhotel.order.hotelsign.query
//
// 获取直连酒店（客栈）签约上线进度信息
type TaobaoxhotelorderhotelsignqueryAPIRequest struct {
	model.Params
	// 请求流水
	_outUuid string
	// 商家酒店编码
	_hotelCode string
	// 商家vendor
	_vendor string
	// 1
	_type string
	// 页码
	_pageNo int64
}

// NewTaobaoxhotelorderhotelsignqueryRequest 初始化TaobaoxhotelorderhotelsignqueryAPIRequest对象
func NewTaobaoxhotelorderhotelsignqueryRequest() *TaobaoxhotelorderhotelsignqueryAPIRequest {
	return &TaobaoxhotelorderhotelsignqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelorderhotelsignqueryAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.hotelsign.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelorderhotelsignqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelorderhotelsignqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutUuid is OutUuid Setter
// 请求流水
func (r *TaobaoxhotelorderhotelsignqueryAPIRequest) SetOutUuid(_outUuid string) error {
	r._outUuid = _outUuid
	r.Set("out_uuid", _outUuid)
	return nil
}

// GetOutUuid OutUuid Getter
func (r TaobaoxhotelorderhotelsignqueryAPIRequest) GetOutUuid() string {
	return r._outUuid
}

// SetHotelCode is HotelCode Setter
// 商家酒店编码
func (r *TaobaoxhotelorderhotelsignqueryAPIRequest) SetHotelCode(_hotelCode string) error {
	r._hotelCode = _hotelCode
	r.Set("hotel_code", _hotelCode)
	return nil
}

// GetHotelCode HotelCode Getter
func (r TaobaoxhotelorderhotelsignqueryAPIRequest) GetHotelCode() string {
	return r._hotelCode
}

// SetVendor is Vendor Setter
// 商家vendor
func (r *TaobaoxhotelorderhotelsignqueryAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelorderhotelsignqueryAPIRequest) GetVendor() string {
	return r._vendor
}

// SetType is Type Setter
// 1
func (r *TaobaoxhotelorderhotelsignqueryAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoxhotelorderhotelsignqueryAPIRequest) GetType() string {
	return r._type
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoxhotelorderhotelsignqueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoxhotelorderhotelsignqueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
