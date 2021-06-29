package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
房型查询接口 API请求
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

// 初始化TaobaoXhotelRoomtypeGetRequest对象
func NewTaobaoXhotelRoomtypeGetRequest() *TaobaoXhotelRoomtypeGetRequest{
    return &TaobaoXhotelRoomtypeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomtypeGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.roomtype.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomtypeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rid Setter
// 废弃，使用商家房型ID
func (r *TaobaoXhotelRoomtypeGetRequest) SetRid(rid int64) error {
    r.rid = rid
    r.Set("rid", rid)
    return nil
}

// Rid Getter
func (r TaobaoXhotelRoomtypeGetRequest) GetRid() int64 {
    return r.rid
}
// OuterId Setter
// 商家房型ID
func (r *TaobaoXhotelRoomtypeGetRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelRoomtypeGetRequest) GetOuterId() string {
    return r.outerId
}
// Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelRoomtypeGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRoomtypeGetRequest) GetVendor() string {
    return r.vendor
}
