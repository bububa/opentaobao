package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店库修改基本信息 APIRequest
taobao.place.storegroup.update

门店库修改基本信息
*/
type TaobaoPlaceStoregroupUpdateRequest struct {
    model.Params

    // 库id
    id   int64 

    // 库名称
    name   string 

    // 库备注
    desc   string 

}

func NewTaobaoPlaceStoregroupUpdateRequest() *TaobaoPlaceStoregroupUpdateRequest{
    return &TaobaoPlaceStoregroupUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStoregroupUpdateRequest) GetApiMethodName() string {
    return "taobao.place.storegroup.update"
}

func (r TaobaoPlaceStoregroupUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStoregroupUpdateRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r TaobaoPlaceStoregroupUpdateRequest) GetId() int64 {
    return r.id
}

func (r *TaobaoPlaceStoregroupUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoPlaceStoregroupUpdateRequest) GetName() string {
    return r.name
}

func (r *TaobaoPlaceStoregroupUpdateRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

func (r TaobaoPlaceStoregroupUpdateRequest) GetDesc() string {
    return r.desc
}

