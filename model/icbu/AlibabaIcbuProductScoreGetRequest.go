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
type AlibabaIcbuProductScoreGetAPIRequest struct {
    model.Params
    // 混淆后的商品ID
    _productId   string
}

// 初始化AlibabaIcbuProductScoreGetAPIRequest对象
func NewAlibabaIcbuProductScoreGetRequest() *AlibabaIcbuProductScoreGetAPIRequest{
    return &AlibabaIcbuProductScoreGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductScoreGetAPIRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.score.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductScoreGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 混淆后的商品ID
func (r *AlibabaIcbuProductScoreGetAPIRequest) SetProductId(_productId string) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r AlibabaIcbuProductScoreGetAPIRequest) GetProductId() string {
    return r._productId
}
