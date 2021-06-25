package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取门店类目信息 APIRequest
taobao.place.storecategory.get

获取门店类目信息
*/
type TaobaoPlaceStorecategoryGetRequest struct {
    model.Params

}

func NewTaobaoPlaceStorecategoryGetRequest() *TaobaoPlaceStorecategoryGetRequest{
    return &TaobaoPlaceStorecategoryGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPlaceStorecategoryGetRequest) GetApiMethodName() string {
    return "taobao.place.storecategory.get"
}

func (r TaobaoPlaceStorecategoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


