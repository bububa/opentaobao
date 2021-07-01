package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户门店库创建接口 API请求
taobao.place.storegroup.create

用于商家创建线下门店库
*/
type TaobaoPlaceStoregroupCreateAPIRequest struct {
    model.Params
    // 库名
    _name   string
    // 备注
    _desc   string
}

// 初始化TaobaoPlaceStoregroupCreateAPIRequest对象
func NewTaobaoPlaceStoregroupCreateRequest() *TaobaoPlaceStoregroupCreateAPIRequest{
    return &TaobaoPlaceStoregroupCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoregroupCreateAPIRequest) GetApiMethodName() string {
    return "taobao.place.storegroup.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoregroupCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 库名
func (r *TaobaoPlaceStoregroupCreateAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoPlaceStoregroupCreateAPIRequest) GetName() string {
    return r._name
}
// Desc Setter
// 备注
func (r *TaobaoPlaceStoregroupCreateAPIRequest) SetDesc(_desc string) error {
    r._desc = _desc
    r.Set("desc", _desc)
    return nil
}

// Desc Getter
func (r TaobaoPlaceStoregroupCreateAPIRequest) GetDesc() string {
    return r._desc
}
