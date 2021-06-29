package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家删除房型数据接口 APIRequest
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

func NewTaobaoXhotelRoomtypeDeletePublicRequest() *TaobaoXhotelRoomtypeDeletePublicRequest{
    return &TaobaoXhotelRoomtypeDeletePublicRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelRoomtypeDeletePublicRequest) GetApiMethodName() string {
    return "taobao.xhotel.roomtype.delete.public"
}

func (r TaobaoXhotelRoomtypeDeletePublicRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelRoomtypeDeletePublicRequest) SetRid(rid int64) error {
    r.rid = rid
    r.Set("rid", rid)
    return nil
}

func (r TaobaoXhotelRoomtypeDeletePublicRequest) GetRid() int64 {
    return r.rid
}

func (r *TaobaoXhotelRoomtypeDeletePublicRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelRoomtypeDeletePublicRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelRoomtypeDeletePublicRequest) SetOuterRid(outerRid string) error {
    r.outerRid = outerRid
    r.Set("outer_rid", outerRid)
    return nil
}

func (r TaobaoXhotelRoomtypeDeletePublicRequest) GetOuterRid() string {
    return r.outerRid
}

func (r *TaobaoXhotelRoomtypeDeletePublicRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

func (r TaobaoXhotelRoomtypeDeletePublicRequest) GetOperator() string {
    return r.operator
}

