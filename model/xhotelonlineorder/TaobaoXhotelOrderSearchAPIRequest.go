package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderSearchAPIRequest
酒店产品库订单查询 API请求
taobao.xhotel.order.search

酒店产品库订单查询功能，查询90天内的订单 */
type TaobaoXhotelOrderSearchAPIRequest struct {
	model.Params
	// 酒店订单oids列表，多个oid用英文逗号隔开，一次不超过20个。
	_orderIds string
	// 订单创建时间查询结束时间，格式为：yyyy-MM-dd HH:mm:ss。不能早于created_start或者间隔不能大于30
	_createdEnd string
	// 订单创建时间查询起始时间，格式为：yyyy-MM-dd HH:mm:ss
	_createdStart string
	// 分页页码。取值范围，大于零的整数，默认值1，即返回第一页的数据。页面大小为20
	_pageNo int64
	// 酒店订单tids列表，多个tid用英文逗号隔开，一次不超过20个。oids和tids都传的情况下默认使用tids
	_orderTids string
	// 系统商标识
	_vendor string
	// 外部订单号out_oids列表，多个oid用英文逗号隔开，一次不超过20个。
	_outOids string
}

// NewTaobaoXhotelOrderSearchRequest 初始化TaobaoXhotelOrderSearchAPIRequest对象
func NewTaobaoXhotelOrderSearchRequest() *TaobaoXhotelOrderSearchAPIRequest {
	return &TaobaoXhotelOrderSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderSearchAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderIds Setter
// 酒店订单oids列表，多个oid用英文逗号隔开，一次不超过20个。
func (r *TaobaoXhotelOrderSearchAPIRequest) SetOrderIds(_orderIds string) error {
	r._orderIds = _orderIds
	r.Set("order_ids", _orderIds)
	return nil
}

// Get OrderIds Getter
func (r TaobaoXhotelOrderSearchAPIRequest) GetOrderIds() string {
	return r._orderIds
}

// Set is CreatedEnd Setter
// 订单创建时间查询结束时间，格式为：yyyy-MM-dd HH:mm:ss。不能早于created_start或者间隔不能大于30
func (r *TaobaoXhotelOrderSearchAPIRequest) SetCreatedEnd(_createdEnd string) error {
	r._createdEnd = _createdEnd
	r.Set("created_end", _createdEnd)
	return nil
}

// Get CreatedEnd Getter
func (r TaobaoXhotelOrderSearchAPIRequest) GetCreatedEnd() string {
	return r._createdEnd
}

// Set is CreatedStart Setter
// 订单创建时间查询起始时间，格式为：yyyy-MM-dd HH:mm:ss
func (r *TaobaoXhotelOrderSearchAPIRequest) SetCreatedStart(_createdStart string) error {
	r._createdStart = _createdStart
	r.Set("created_start", _createdStart)
	return nil
}

// Get CreatedStart Getter
func (r TaobaoXhotelOrderSearchAPIRequest) GetCreatedStart() string {
	return r._createdStart
}

// Set is PageNo Setter
// 分页页码。取值范围，大于零的整数，默认值1，即返回第一页的数据。页面大小为20
func (r *TaobaoXhotelOrderSearchAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r TaobaoXhotelOrderSearchAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is OrderTids Setter
// 酒店订单tids列表，多个tid用英文逗号隔开，一次不超过20个。oids和tids都传的情况下默认使用tids
func (r *TaobaoXhotelOrderSearchAPIRequest) SetOrderTids(_orderTids string) error {
	r._orderTids = _orderTids
	r.Set("order_tids", _orderTids)
	return nil
}

// Get OrderTids Getter
func (r TaobaoXhotelOrderSearchAPIRequest) GetOrderTids() string {
	return r._orderTids
}

// Set is Vendor Setter
// 系统商标识
func (r *TaobaoXhotelOrderSearchAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// Get Vendor Getter
func (r TaobaoXhotelOrderSearchAPIRequest) GetVendor() string {
	return r._vendor
}

// Set is OutOids Setter
// 外部订单号out_oids列表，多个oid用英文逗号隔开，一次不超过20个。
func (r *TaobaoXhotelOrderSearchAPIRequest) SetOutOids(_outOids string) error {
	r._outOids = _outOids
	r.Set("out_oids", _outOids)
	return nil
}

// Get OutOids Getter
func (r TaobaoXhotelOrderSearchAPIRequest) GetOutOids() string {
	return r._outOids
}
