package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除门店库 APIRequest
taobao.place.storegroup.delete

删除门店库
*/
type TaobaoPlaceStoregroupDeleteRequest struct {
    model.Params

    // 库Id
    id   int64 

}

func NewTaobaoPlaceStoregroupDeleteRequest() *TaobaoPlaceStoregroupDeleteRequest{
    return &TaobaoPlaceStoregroupDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStoregroupDeleteRequest) GetApiMethodName() string {
    return "taobao.place.storegroup.delete"
}

func (r TaobaoPlaceStoregroupDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPlaceStoregroupDeleteRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r TaobaoPlaceStoregroupDeleteRequest) GetId() int64 {
    return r.id
}

