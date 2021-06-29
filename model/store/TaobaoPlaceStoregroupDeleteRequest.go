package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除门店库 API请求
taobao.place.storegroup.delete

删除门店库
*/
type TaobaoPlaceStoregroupDeleteRequest struct {
    model.Params
    // 库Id
    id   int64
}

// 初始化TaobaoPlaceStoregroupDeleteRequest对象
func NewTaobaoPlaceStoregroupDeleteRequest() *TaobaoPlaceStoregroupDeleteRequest{
    return &TaobaoPlaceStoregroupDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoregroupDeleteRequest) GetApiMethodName() string {
    return "taobao.place.storegroup.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoregroupDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 库Id
func (r *TaobaoPlaceStoregroupDeleteRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TaobaoPlaceStoregroupDeleteRequest) GetId() int64 {
    return r.id
}
