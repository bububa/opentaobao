package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店查询接口 API请求
taobao.xhotel.get

酒店查询接口
*/
type TaobaoXhotelGetRequest struct {
    model.Params
    // 废弃，请使用outer_id
    hid   int64
    // 卖家系统中的酒店ID。
    outerId   string
    // 系统商，一般不用填写，使用须申请
    vendor   string
    // 是否需要在售状态(默认false不需要)
    needSaleInfo   bool
}

// 初始化TaobaoXhotelGetRequest对象
func NewTaobaoXhotelGetRequest() *TaobaoXhotelGetRequest{
    return &TaobaoXhotelGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Hid Setter
// 废弃，请使用outer_id
func (r *TaobaoXhotelGetRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelGetRequest) GetHid() int64 {
    return r.hid
}
// OuterId Setter
// 卖家系统中的酒店ID。
func (r *TaobaoXhotelGetRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelGetRequest) GetOuterId() string {
    return r.outerId
}
// Vendor Setter
// 系统商，一般不用填写，使用须申请
func (r *TaobaoXhotelGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelGetRequest) GetVendor() string {
    return r.vendor
}
// NeedSaleInfo Setter
// 是否需要在售状态(默认false不需要)
func (r *TaobaoXhotelGetRequest) SetNeedSaleInfo(needSaleInfo bool) error {
    r.needSaleInfo = needSaleInfo
    r.Set("need_sale_info", needSaleInfo)
    return nil
}

// NeedSaleInfo Getter
func (r TaobaoXhotelGetRequest) GetNeedSaleInfo() bool {
    return r.needSaleInfo
}
