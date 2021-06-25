package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
失效指定用户的商品与后端商品的映射关系 APIRequest
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

func NewTaobaoScitemMapDeleteRequest() *TaobaoScitemMapDeleteRequest{
    return &TaobaoScitemMapDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoScitemMapDeleteRequest) GetApiMethodName() string {
    return "taobao.scitem.map.delete"
}

func (r TaobaoScitemMapDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoScitemMapDeleteRequest) SetScItemId(scItemId int64) error {
    r.scItemId = scItemId
    r.Set("sc_item_id", scItemId)
    return nil
}

func (r TaobaoScitemMapDeleteRequest) GetScItemId() int64 {
    return r.scItemId
}

func (r *TaobaoScitemMapDeleteRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

func (r TaobaoScitemMapDeleteRequest) GetUserNick() string {
    return r.userNick
}

