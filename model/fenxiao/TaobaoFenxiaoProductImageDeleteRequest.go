package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品图片删除 API请求
taobao.fenxiao.product.image.delete

产品图片删除，只删除图片信息，不真正删除图片
*/
type TaobaoFenxiaoProductImageDeleteRequest struct {
    model.Params
    // 产品ID
    _productId   int64
    // 图片位置
    _position   int64
    // properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
    _properties   string
}

// 初始化TaobaoFenxiaoProductImageDeleteRequest对象
func NewTaobaoFenxiaoProductImageDeleteRequest() *TaobaoFenxiaoProductImageDeleteRequest{
    return &TaobaoFenxiaoProductImageDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductImageDeleteRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.image.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductImageDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoProductImageDeleteRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TaobaoFenxiaoProductImageDeleteRequest) GetProductId() int64 {
    return r._productId
}
// Position Setter
// 图片位置
func (r *TaobaoFenxiaoProductImageDeleteRequest) SetPosition(_position int64) error {
    r._position = _position
    r.Set("position", _position)
    return nil
}

// Position Getter
func (r TaobaoFenxiaoProductImageDeleteRequest) GetPosition() int64 {
    return r._position
}
// Properties Setter
// properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
func (r *TaobaoFenxiaoProductImageDeleteRequest) SetProperties(_properties string) error {
    r._properties = _properties
    r.Set("properties", _properties)
    return nil
}

// Properties Getter
func (r TaobaoFenxiaoProductImageDeleteRequest) GetProperties() string {
    return r._properties
}
