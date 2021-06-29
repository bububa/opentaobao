package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据id查询商品 API请求
taobao.scitem.get

根据id查询商品
*/
type TaobaoScitemGetRequest struct {
    model.Params
    // 商品id
    itemId   int64
}

// 初始化TaobaoScitemGetRequest对象
func NewTaobaoScitemGetRequest() *TaobaoScitemGetRequest{
    return &TaobaoScitemGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoScitemGetRequest) GetApiMethodName() string {
    return "taobao.scitem.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoScitemGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoScitemGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoScitemGetRequest) GetItemId() int64 {
    return r.itemId
}
