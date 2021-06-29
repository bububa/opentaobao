package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
房型查询接口 APIRequest
taobao.xhotel.roomtype.get

房型查询房型查询接口返回结果增加date_confirm字段
*/
type TaobaoXhotelRoomtypeGetRequest struct {
    model.Params

    // 废弃，使用商家房型ID
    rid   int64 

    // 商家房型ID
    outerId   string 

    // 系统商，一般不填写，使用须申请
    vendor   string 

}

func NewTaobaoXhotelRoomtypeGetRequest() *TaobaoXhotelRoomtypeGetRequest{
    return &TaobaoXhotelRoomtypeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelRoomtypeGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.roomtype.get"
}

func (r TaobaoXhotelRoomtypeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelRoomtypeGetRequest) SetRid(rid int64) error {
    r.rid = rid
    r.Set("rid", rid)
    return nil
}

func (r TaobaoXhotelRoomtypeGetRequest) GetRid() int64 {
    return r.rid
}

func (r *TaobaoXhotelRoomtypeGetRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoXhotelRoomtypeGetRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoXhotelRoomtypeGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelRoomtypeGetRequest) GetVendor() string {
    return r.vendor
}

