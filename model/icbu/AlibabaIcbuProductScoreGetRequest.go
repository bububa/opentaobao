package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品质量分查询 API请求
alibaba.icbu.product.score.get

产品质量分查询
*/
type AlibabaIcbuProductScoreGetRequest struct {
    model.Params
    // 混淆后的商品ID
    productId   string
}

// 初始化AlibabaIcbuProductScoreGetRequest对象
func NewAlibabaIcbuProductScoreGetRequest() *AlibabaIcbuProductScoreGetRequest{
    return &AlibabaIcbuProductScoreGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductScoreGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.score.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductScoreGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 混淆后的商品ID
func (r *AlibabaIcbuProductScoreGetRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r AlibabaIcbuProductScoreGetRequest) GetProductId() string {
    return r.productId
}
