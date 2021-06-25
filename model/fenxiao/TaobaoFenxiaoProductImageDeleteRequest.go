package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品图片删除 APIRequest
taobao.fenxiao.product.image.delete

产品图片删除，只删除图片信息，不真正删除图片
*/
type TaobaoFenxiaoProductImageDeleteRequest struct {
    model.Params

    // 产品ID
    productId   int64 

    // 图片位置
    position   int64 

    // properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项
    properties   string 

}

func NewTaobaoFenxiaoProductImageDeleteRequest() *TaobaoFenxiaoProductImageDeleteRequest{
    return &TaobaoFenxiaoProductImageDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoProductImageDeleteRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.image.delete"
}

func (r TaobaoFenxiaoProductImageDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoProductImageDeleteRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TaobaoFenxiaoProductImageDeleteRequest) GetProductId() int64 {
    return r.productId
}

func (r *TaobaoFenxiaoProductImageDeleteRequest) SetPosition(position int64) error {
    r.position = position
    r.Set("position", position)
    return nil
}

func (r TaobaoFenxiaoProductImageDeleteRequest) GetPosition() int64 {
    return r.position
}

func (r *TaobaoFenxiaoProductImageDeleteRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

func (r TaobaoFenxiaoProductImageDeleteRequest) GetProperties() string {
    return r.properties
}

