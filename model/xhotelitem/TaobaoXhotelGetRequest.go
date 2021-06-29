package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店查询接口 APIRequest
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

func NewTaobaoXhotelGetRequest() *TaobaoXhotelGetRequest{
    return &TaobaoXhotelGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.get"
}

func (r TaobaoXhotelGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelGetRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

func (r TaobaoXhotelGetRequest) GetHid() int64 {
    return r.hid
}

func (r *TaobaoXhotelGetRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoXhotelGetRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoXhotelGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelGetRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelGetRequest) SetNeedSaleInfo(needSaleInfo bool) error {
    r.needSaleInfo = needSaleInfo
    r.Set("need_sale_info", needSaleInfo)
    return nil
}

func (r TaobaoXhotelGetRequest) GetNeedSaleInfo() bool {
    return r.needSaleInfo
}

