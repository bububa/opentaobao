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
    _id   int64
    // 库名称
    _name   string
    // 库备注
    _desc   string
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
func (r *TaobaoPlaceStoregroupUpdateRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r TaobaoPlaceStoregroupUpdateRequest) GetId() int64 {
    return r._id
}
// Name Setter
// 库名称
func (r *TaobaoPlaceStoregroupUpdateRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoPlaceStoregroupUpdateRequest) GetName() string {
    return r._name
}
// Desc Setter
// 库备注
func (r *TaobaoPlaceStoregroupUpdateRequest) SetDesc(_desc string) error {
    r._desc = _desc
    r.Set("desc", _desc)
    return nil
}

// Desc Getter
func (r TaobaoPlaceStoregroupUpdateRequest) GetDesc() string {
    return r._desc
}
