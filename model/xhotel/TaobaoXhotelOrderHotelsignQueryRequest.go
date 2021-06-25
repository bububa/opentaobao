package xhotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取直连酒店（客栈）签约上线进度信息 APIRequest
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

func NewTaobaoXhotelOrderHotelsignQueryRequest() *TaobaoXhotelOrderHotelsignQueryRequest{
    return &TaobaoXhotelOrderHotelsignQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelOrderHotelsignQueryRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.hotelsign.query"
}

func (r TaobaoXhotelOrderHotelsignQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelOrderHotelsignQueryRequest) SetOutUuid(outUuid string) error {
    r.outUuid = outUuid
    r.Set("out_uuid", outUuid)
    return nil
}

func (r TaobaoXhotelOrderHotelsignQueryRequest) GetOutUuid() string {
    return r.outUuid
}

func (r *TaobaoXhotelOrderHotelsignQueryRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

func (r TaobaoXhotelOrderHotelsignQueryRequest) GetHotelCode() string {
    return r.hotelCode
}

func (r *TaobaoXhotelOrderHotelsignQueryRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelOrderHotelsignQueryRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelOrderHotelsignQueryRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoXhotelOrderHotelsignQueryRequest) GetType() string {
    return r.type
}

func (r *TaobaoXhotelOrderHotelsignQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoXhotelOrderHotelsignQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

