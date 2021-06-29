package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店库修改基本信息 API请求
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

// 初始化TaobaoPlaceStoregroupUpdateRequest对象
func NewTaobaoPlaceStoregroupUpdateRequest() *TaobaoPlaceStoregroupUpdateRequest{
    return &TaobaoPlaceStoregroupUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoregroupUpdateRequest) GetApiMethodName() string {
    return "taobao.place.storegroup.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoregroupUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 库id
func (r *TaobaoPlaceStoregroupUpdateRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TaobaoPlaceStoregroupUpdateRequest) GetId() int64 {
    return r.id
}
// Name Setter
// 库名称
func (r *TaobaoPlaceStoregroupUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoPlaceStoregroupUpdateRequest) GetName() string {
    return r.name
}
// Desc Setter
// 库备注
func (r *TaobaoPlaceStoregroupUpdateRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

// Desc Getter
func (r TaobaoPlaceStoregroupUpdateRequest) GetDesc() string {
    return r.desc
}
