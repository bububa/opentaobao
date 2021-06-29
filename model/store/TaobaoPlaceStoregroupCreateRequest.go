package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户门店库创建接口 APIRequest
taobao.place.storegroup.create

用于商家创建线下门店库
*/
type TaobaoPlaceStoregroupCreateRequest struct {
    model.Params

    // 库名
    name   string 

    // 备注
    desc   string 

}

func NewTaobaoPlaceStoregroupCreateRequest() *TaobaoPlaceStoregroupCreateRequest{
    return &TaobaoPlaceStoregroupCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStoregroupCreateRequest) GetApiMethodName() string {
    return "taobao.place.storegroup.create"
}

func (r TaobaoPlaceStoregroupCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStoregroupCreateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoPlaceStoregroupCreateRequest) GetName() string {
    return r.name
}

func (r *TaobaoPlaceStoregroupCreateRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

func (r TaobaoPlaceStoregroupCreateRequest) GetDesc() string {
    return r.desc
}

