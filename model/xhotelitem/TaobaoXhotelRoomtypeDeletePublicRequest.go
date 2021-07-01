package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家删除房型数据接口 API请求
taobao.xhotel.roomtype.delete.public

房型删除TOP接口
*/
type TaobaoXhotelRoomtypeDeletePublicAPIRequest struct {
    model.Params
    // 房型rid ，传参方式：rid    或者   outer_id+vendor
    _rid   int64
    // vendor
    _vendor   string
    // 外部房型ID
    _outerRid   string
    // 具体操作人，比如酒店帐号、小二名称等
    _operator   string
}

// 初始化TaobaoXhotelRoomtypeDeletePublicAPIRequest对象
func NewTaobaoXhotelRoomtypeDeletePublicRequest() *TaobaoXhotelRoomtypeDeletePublicAPIRequest{
    return &TaobaoXhotelRoomtypeDeletePublicAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomtypeDeletePublicAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.roomtype.delete.public"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomtypeDeletePublicAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rid Setter
// 房型rid ，传参方式：rid    或者   outer_id+vendor
func (r *TaobaoXhotelRoomtypeDeletePublicAPIRequest) SetRid(_rid int64) error {
    r._rid = _rid
    r.Set("rid", _rid)
    return nil
}

// Rid Getter
func (r TaobaoXhotelRoomtypeDeletePublicAPIRequest) GetRid() int64 {
    return r._rid
}
// Vendor Setter
// vendor
func (r *TaobaoXhotelRoomtypeDeletePublicAPIRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRoomtypeDeletePublicAPIRequest) GetVendor() string {
    return r._vendor
}
// OuterRid Setter
// 外部房型ID
func (r *TaobaoXhotelRoomtypeDeletePublicAPIRequest) SetOuterRid(_outerRid string) error {
    r._outerRid = _outerRid
    r.Set("outer_rid", _outerRid)
    return nil
}

// OuterRid Getter
func (r TaobaoXhotelRoomtypeDeletePublicAPIRequest) GetOuterRid() string {
    return r._outerRid
}
// Operator Setter
// 具体操作人，比如酒店帐号、小二名称等
func (r *TaobaoXhotelRoomtypeDeletePublicAPIRequest) SetOperator(_operator string) error {
    r._operator = _operator
    r.Set("operator", _operator)
    return nil
}

// Operator Getter
func (r TaobaoXhotelRoomtypeDeletePublicAPIRequest) GetOperator() string {
    return r._operator
}
