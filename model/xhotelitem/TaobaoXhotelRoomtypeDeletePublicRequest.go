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
type TaobaoXhotelRoomtypeDeletePublicRequest struct {
    model.Params
    // 房型rid ，传参方式：rid    或者   outer_id+vendor
    rid   int64
    // vendor
    vendor   string
    // 外部房型ID
    outerRid   string
    // 具体操作人，比如酒店帐号、小二名称等
    operator   string
}

// 初始化TaobaoXhotelRoomtypeDeletePublicRequest对象
func NewTaobaoXhotelRoomtypeDeletePublicRequest() *TaobaoXhotelRoomtypeDeletePublicRequest{
    return &TaobaoXhotelRoomtypeDeletePublicRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomtypeDeletePublicRequest) GetApiMethodName() string {
    return "taobao.xhotel.roomtype.delete.public"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomtypeDeletePublicRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rid Setter
// 房型rid ，传参方式：rid    或者   outer_id+vendor
func (r *TaobaoXhotelRoomtypeDeletePublicRequest) SetRid(rid int64) error {
    r.rid = rid
    r.Set("rid", rid)
    return nil
}

// Rid Getter
func (r TaobaoXhotelRoomtypeDeletePublicRequest) GetRid() int64 {
    return r.rid
}
// Vendor Setter
// vendor
func (r *TaobaoXhotelRoomtypeDeletePublicRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRoomtypeDeletePublicRequest) GetVendor() string {
    return r.vendor
}
// OuterRid Setter
// 外部房型ID
func (r *TaobaoXhotelRoomtypeDeletePublicRequest) SetOuterRid(outerRid string) error {
    r.outerRid = outerRid
    r.Set("outer_rid", outerRid)
    return nil
}

// OuterRid Getter
func (r TaobaoXhotelRoomtypeDeletePublicRequest) GetOuterRid() string {
    return r.outerRid
}
// Operator Setter
// 具体操作人，比如酒店帐号、小二名称等
func (r *TaobaoXhotelRoomtypeDeletePublicRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

// Operator Getter
func (r TaobaoXhotelRoomtypeDeletePublicRequest) GetOperator() string {
    return r.operator
}
