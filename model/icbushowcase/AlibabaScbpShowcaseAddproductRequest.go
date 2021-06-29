package icbushowcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量添加橱窗商品 APIRequest
alibaba.scbp.showcase.addproduct

批量添加商品到橱窗
*/
type AlibabaScbpShowcaseAddproductRequest struct {
    model.Params

    // 需要添加的产品ids
    productIdList   []int64 

}

func NewAlibabaScbpShowcaseAddproductRequest() *AlibabaScbpShowcaseAddproductRequest{
    return &AlibabaScbpShowcaseAddproductRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpShowcaseAddproductRequest) GetApiMethodName() string {
    return "alibaba.scbp.showcase.addproduct"
}

func (r AlibabaScbpShowcaseAddproductRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpShowcaseAddproductRequest) SetProductIdList(productIdList []int64) error {
    r.productIdList = productIdList
    r.Set("product_id_list", productIdList)
    return nil
}

func (r AlibabaScbpShowcaseAddproductRequest) GetProductIdList() []int64 {
    return r.productIdList
}

