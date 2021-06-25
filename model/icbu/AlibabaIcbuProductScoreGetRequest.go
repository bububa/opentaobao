package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品质量分查询 APIRequest
alibaba.icbu.product.score.get

产品质量分查询
*/
type AlibabaIcbuProductScoreGetRequest struct {
    model.Params

    // 混淆后的商品ID
    productId   string 

}

func NewAlibabaIcbuProductScoreGetRequest() *AlibabaIcbuProductScoreGetRequest{
    return &AlibabaIcbuProductScoreGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuProductScoreGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.product.score.get"
}

func (r AlibabaIcbuProductScoreGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuProductScoreGetRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r AlibabaIcbuProductScoreGetRequest) GetProductId() string {
    return r.productId
}

