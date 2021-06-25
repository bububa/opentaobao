package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新发商品发布页 APIRequest
taobao.banamadpc.item.render

巴拿马供应商通过此接口新发商品发布页
*/
type TaobaoBanamadpcItemRenderRequest struct {
    model.Params

    // 类目ID
    catId   int64 

}

func NewTaobaoBanamadpcItemRenderRequest() *TaobaoBanamadpcItemRenderRequest{
    return &TaobaoBanamadpcItemRenderRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBanamadpcItemRenderRequest) GetApiMethodName() string {
    return "taobao.banamadpc.item.render"
}

func (r TaobaoBanamadpcItemRenderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBanamadpcItemRenderRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

func (r TaobaoBanamadpcItemRenderRequest) GetCatId() int64 {
    return r.catId
}

