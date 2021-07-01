package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新发商品发布页 API请求
taobao.banamadpc.item.render

巴拿马供应商通过此接口新发商品发布页
*/
type TaobaoBanamadpcItemRenderAPIRequest struct {
    model.Params
    // 类目ID
    _catId   int64
}

// 初始化TaobaoBanamadpcItemRenderAPIRequest对象
func NewTaobaoBanamadpcItemRenderRequest() *TaobaoBanamadpcItemRenderAPIRequest{
    return &TaobaoBanamadpcItemRenderAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBanamadpcItemRenderAPIRequest) GetApiMethodName() string {
    return "taobao.banamadpc.item.render"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBanamadpcItemRenderAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 类目ID
func (r *TaobaoBanamadpcItemRenderAPIRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r TaobaoBanamadpcItemRenderAPIRequest) GetCatId() int64 {
    return r._catId
}
