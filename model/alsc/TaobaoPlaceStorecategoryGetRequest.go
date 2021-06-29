package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取门店类目信息 API请求
taobao.place.storecategory.get

获取门店类目信息
*/
type TaobaoPlaceStorecategoryGetRequest struct {
    model.Params
}

// 初始化TaobaoPlaceStorecategoryGetRequest对象
func NewTaobaoPlaceStorecategoryGetRequest() *TaobaoPlaceStorecategoryGetRequest{
    return &TaobaoPlaceStorecategoryGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStorecategoryGetRequest) GetApiMethodName() string {
    return "taobao.place.storecategory.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStorecategoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
