package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
失效指定用户的商品与后端商品的映射关系 API请求
taobao.scitem.map.delete

根据后端商品Id，失效指定用户的商品与后端商品的映射关系
*/
type TaobaoScitemMapDeleteRequest struct {
    model.Params
    // 后台商品ID
    scItemId   int64
    // 店铺用户nick。 如果该参数为空则删除后端商品与当前调用人的商品映射关系;如果不为空则删除指定用户与后端商品的映射关系
    userNick   string
}

// 初始化TaobaoScitemMapDeleteRequest对象
func NewTaobaoScitemMapDeleteRequest() *TaobaoScitemMapDeleteRequest{
    return &TaobaoScitemMapDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoScitemMapDeleteRequest) GetApiMethodName() string {
    return "taobao.scitem.map.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoScitemMapDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ScItemId Setter
// 后台商品ID
func (r *TaobaoScitemMapDeleteRequest) SetScItemId(scItemId int64) error {
    r.scItemId = scItemId
    r.Set("sc_item_id", scItemId)
    return nil
}

// ScItemId Getter
func (r TaobaoScitemMapDeleteRequest) GetScItemId() int64 {
    return r.scItemId
}
// UserNick Setter
// 店铺用户nick。 如果该参数为空则删除后端商品与当前调用人的商品映射关系;如果不为空则删除指定用户与后端商品的映射关系
func (r *TaobaoScitemMapDeleteRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

// UserNick Getter
func (r TaobaoScitemMapDeleteRequest) GetUserNick() string {
    return r.userNick
}
