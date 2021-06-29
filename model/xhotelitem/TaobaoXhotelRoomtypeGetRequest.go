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
    _rid   int64
    // 商家房型ID
    _outerId   string
    // 系统商，一般不填写，使用须申请
    _vendor   string
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
func (r *TaobaoXhotelRoomtypeGetRequest) SetRid(_rid int64) error {
    r._rid = _rid
    r.Set("rid", _rid)
    return nil
}

// Rid Getter
func (r TaobaoXhotelRoomtypeGetRequest) GetRid() int64 {
    return r._rid
}
// OuterId Setter
// 商家房型ID
func (r *TaobaoXhotelRoomtypeGetRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelRoomtypeGetRequest) GetOuterId() string {
    return r._outerId
}
// Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelRoomtypeGetRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRoomtypeGetRequest) GetVendor() string {
    return r._vendor
}
