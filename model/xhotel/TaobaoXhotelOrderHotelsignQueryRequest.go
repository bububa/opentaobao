package xhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取直连酒店（客栈）签约上线进度信息 API请求
taobao.xhotel.order.hotelsign.query

获取直连酒店（客栈）签约上线进度信息
*/
type TaobaoXhotelOrderHotelsignQueryRequest struct {
    model.Params
    // 请求流水
    outUuid   string
    // 商家酒店编码
    hotelCode   string
    // 商家vendor
    vendor   string
    // 1
    type   string
    // 页码
    pageNo   int64
}

// 初始化TaobaoXhotelOrderHotelsignQueryRequest对象
func NewTaobaoXhotelOrderHotelsignQueryRequest() *TaobaoXhotelOrderHotelsignQueryRequest{
    return &TaobaoXhotelOrderHotelsignQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderHotelsignQueryRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.hotelsign.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderHotelsignQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutUuid Setter
// 请求流水
func (r *TaobaoXhotelOrderHotelsignQueryRequest) SetOutUuid(outUuid string) error {
    r.outUuid = outUuid
    r.Set("out_uuid", outUuid)
    return nil
}

// OutUuid Getter
func (r TaobaoXhotelOrderHotelsignQueryRequest) GetOutUuid() string {
    return r.outUuid
}
// HotelCode Setter
// 商家酒店编码
func (r *TaobaoXhotelOrderHotelsignQueryRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelOrderHotelsignQueryRequest) GetHotelCode() string {
    return r.hotelCode
}
// Vendor Setter
// 商家vendor
func (r *TaobaoXhotelOrderHotelsignQueryRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelOrderHotelsignQueryRequest) GetVendor() string {
    return r.vendor
}
// Type Setter
// 1
func (r *TaobaoXhotelOrderHotelsignQueryRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoXhotelOrderHotelsignQueryRequest) GetType() string {
    return r.type
}
// PageNo Setter
// 页码
func (r *TaobaoXhotelOrderHotelsignQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoXhotelOrderHotelsignQueryRequest) GetPageNo() int64 {
    return r.pageNo
}
