package icbushowcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量添加橱窗商品 API请求
alibaba.scbp.showcase.addproduct

批量添加商品到橱窗
*/
type AlibabaScbpShowcaseAddproductRequest struct {
    model.Params
    // 需要添加的产品ids
    _productIdList   []int64
}

// 初始化AlibabaScbpShowcaseAddproductRequest对象
func NewAlibabaScbpShowcaseAddproductRequest() *AlibabaScbpShowcaseAddproductRequest{
    return &AlibabaScbpShowcaseAddproductRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseAddproductRequest) GetApiMethodName() string {
    return "alibaba.scbp.showcase.addproduct"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseAddproductRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductIdList Setter
// 需要添加的产品ids
func (r *AlibabaScbpShowcaseAddproductRequest) SetProductIdList(_productIdList []int64) error {
    r._productIdList = _productIdList
    r.Set("product_id_list", _productIdList)
    return nil
}

// ProductIdList Getter
func (r AlibabaScbpShowcaseAddproductRequest) GetProductIdList() []int64 {
    return r._productIdList
}
