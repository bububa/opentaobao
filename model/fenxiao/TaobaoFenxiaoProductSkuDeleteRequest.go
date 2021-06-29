package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品SKU删除接口 API请求
taobao.fenxiao.product.sku.delete

根据sku properties删除sku数据
*/
type TaobaoFenxiaoProductSkuDeleteRequest struct {
    model.Params
    // 产品id
    productId   int64
    // sku属性
    properties   string
}

// 初始化TaobaoFenxiaoProductSkuDeleteRequest对象
func NewTaobaoFenxiaoProductSkuDeleteRequest() *TaobaoFenxiaoProductSkuDeleteRequest{
    return &TaobaoFenxiaoProductSkuDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductSkuDeleteRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.sku.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductSkuDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品id
func (r *TaobaoFenxiaoProductSkuDeleteRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TaobaoFenxiaoProductSkuDeleteRequest) GetProductId() int64 {
    return r.productId
}
// Properties Setter
// sku属性
func (r *TaobaoFenxiaoProductSkuDeleteRequest) SetProperties(properties string) error {
    r.properties = properties
    r.Set("properties", properties)
    return nil
}

// Properties Getter
func (r TaobaoFenxiaoProductSkuDeleteRequest) GetProperties() string {
    return r.properties
}
