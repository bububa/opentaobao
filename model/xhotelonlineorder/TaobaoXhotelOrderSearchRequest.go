package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店产品库订单查询 API请求
taobao.xhotel.order.search

酒店产品库订单查询功能，查询90天内的订单
*/
type TaobaoXhotelOrderSearchRequest struct {
    model.Params
    // 酒店订单oids列表，多个oid用英文逗号隔开，一次不超过20个。
    orderIds   string
    // 订单创建时间查询结束时间，格式为：yyyy-MM-dd HH:mm:ss。不能早于created_start或者间隔不能大于30
    createdEnd   string
    // 订单创建时间查询起始时间，格式为：yyyy-MM-dd HH:mm:ss
    createdStart   string
    // 分页页码。取值范围，大于零的整数，默认值1，即返回第一页的数据。页面大小为20
    pageNo   int64
    // 酒店订单tids列表，多个tid用英文逗号隔开，一次不超过20个。oids和tids都传的情况下默认使用tids
    orderTids   string
    // 系统商标识
    vendor   string
    // 外部订单号out_oids列表，多个oid用英文逗号隔开，一次不超过20个。
    outOids   string
}

// 初始化TaobaoXhotelOrderSearchRequest对象
func NewTaobaoXhotelOrderSearchRequest() *TaobaoXhotelOrderSearchRequest{
    return &TaobaoXhotelOrderSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderSearchRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderIds Setter
// 酒店订单oids列表，多个oid用英文逗号隔开，一次不超过20个。
func (r *TaobaoXhotelOrderSearchRequest) SetOrderIds(orderIds string) error {
    r.orderIds = orderIds
    r.Set("order_ids", orderIds)
    return nil
}

// OrderIds Getter
func (r TaobaoXhotelOrderSearchRequest) GetOrderIds() string {
    return r.orderIds
}
// CreatedEnd Setter
// 订单创建时间查询结束时间，格式为：yyyy-MM-dd HH:mm:ss。不能早于created_start或者间隔不能大于30
func (r *TaobaoXhotelOrderSearchRequest) SetCreatedEnd(createdEnd string) error {
    r.createdEnd = createdEnd
    r.Set("created_end", createdEnd)
    return nil
}

// CreatedEnd Getter
func (r TaobaoXhotelOrderSearchRequest) GetCreatedEnd() string {
    return r.createdEnd
}
// CreatedStart Setter
// 订单创建时间查询起始时间，格式为：yyyy-MM-dd HH:mm:ss
func (r *TaobaoXhotelOrderSearchRequest) SetCreatedStart(createdStart string) error {
    r.createdStart = createdStart
    r.Set("created_start", createdStart)
    return nil
}

// CreatedStart Getter
func (r TaobaoXhotelOrderSearchRequest) GetCreatedStart() string {
    return r.createdStart
}
// PageNo Setter
// 分页页码。取值范围，大于零的整数，默认值1，即返回第一页的数据。页面大小为20
func (r *TaobaoXhotelOrderSearchRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoXhotelOrderSearchRequest) GetPageNo() int64 {
    return r.pageNo
}
// OrderTids Setter
// 酒店订单tids列表，多个tid用英文逗号隔开，一次不超过20个。oids和tids都传的情况下默认使用tids
func (r *TaobaoXhotelOrderSearchRequest) SetOrderTids(orderTids string) error {
    r.orderTids = orderTids
    r.Set("order_tids", orderTids)
    return nil
}

// OrderTids Getter
func (r TaobaoXhotelOrderSearchRequest) GetOrderTids() string {
    return r.orderTids
}
// Vendor Setter
// 系统商标识
func (r *TaobaoXhotelOrderSearchRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelOrderSearchRequest) GetVendor() string {
    return r.vendor
}
// OutOids Setter
// 外部订单号out_oids列表，多个oid用英文逗号隔开，一次不超过20个。
func (r *TaobaoXhotelOrderSearchRequest) SetOutOids(outOids string) error {
    r.outOids = outOids
    r.Set("out_oids", outOids)
    return nil
}

// OutOids Getter
func (r TaobaoXhotelOrderSearchRequest) GetOutOids() string {
    return r.outOids
}
