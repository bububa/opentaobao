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
type TaobaoPlaceStoregroupCreateRequest struct {
    model.Params
    // 库名
    name   string
    // 备注
    desc   string
}

// 初始化TaobaoPlaceStoregroupCreateRequest对象
func NewTaobaoPlaceStoregroupCreateRequest() *TaobaoPlaceStoregroupCreateRequest{
    return &TaobaoPlaceStoregroupCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoregroupCreateRequest) GetApiMethodName() string {
    return "taobao.place.storegroup.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoregroupCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 库名
func (r *TaobaoPlaceStoregroupCreateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoPlaceStoregroupCreateRequest) GetName() string {
    return r.name
}
// Desc Setter
// 备注
func (r *TaobaoPlaceStoregroupCreateRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

// Desc Getter
func (r TaobaoPlaceStoregroupCreateRequest) GetDesc() string {
    return r.desc
}
