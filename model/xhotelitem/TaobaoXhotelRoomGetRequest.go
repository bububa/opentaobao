package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
room实体查询 APIRequest
taobao.xhotel.room.get

此接口用于查询一个商品，根据传入的gid查询商品信息。卖家只能查询自己的商品。
*/
type TaobaoXhotelRoomGetRequest struct {
    model.Params

    // 卖家渠道 如果gid为空，那么out_rid和vendor都不能为空。 支持通过gid或者通过out_rid和vendor来获取商品
    vendor   string 

    // 外部房型id 如果gid为空，那么out_rid和vendor都不能为空 支持通过gid或者通过out_rid和vendor来获取商品
    outRid   string 

    // 废弃
    gid   int64 

}

func NewTaobaoXhotelRoomGetRequest() *TaobaoXhotelRoomGetRequest{
    return &TaobaoXhotelRoomGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelRoomGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.room.get"
}

func (r TaobaoXhotelRoomGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelRoomGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelRoomGetRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelRoomGetRequest) SetOutRid(outRid string) error {
    r.outRid = outRid
    r.Set("out_rid", outRid)
    return nil
}

func (r TaobaoXhotelRoomGetRequest) GetOutRid() string {
    return r.outRid
}

func (r *TaobaoXhotelRoomGetRequest) SetGid(gid int64) error {
    r.gid = gid
    r.Set("gid", gid)
    return nil
}

func (r TaobaoXhotelRoomGetRequest) GetGid() int64 {
    return r.gid
}

