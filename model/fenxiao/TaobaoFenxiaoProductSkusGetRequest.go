package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
SKU查询接口 API请求
taobao.fenxiao.product.skus.get

产品sku查询
*/
type TaobaoFenxiaoProductSkusGetRequest struct {
    model.Params
    // 产品ID
    productId   int64
}

// 初始化TaobaoFenxiaoProductSkusGetRequest对象
func NewTaobaoFenxiaoProductSkusGetRequest() *TaobaoFenxiaoProductSkusGetRequest{
    return &TaobaoFenxiaoProductSkusGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductSkusGetRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.skus.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductSkusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *TaobaoFenxiaoProductSkusGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TaobaoFenxiaoProductSkusGetRequest) GetProductId() int64 {
    return r.productId
}
