package icbushowcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
替换橱窗商品 API请求
alibaba.scbp.showcase.updateproduct

替换橱窗商品
*/
type AlibabaScbpShowcaseUpdateproductRequest struct {
    model.Params
    // 橱窗id
    windowId   int64
    // 新的商品id
    newProductId   int64
}

// 初始化AlibabaScbpShowcaseUpdateproductRequest对象
func NewAlibabaScbpShowcaseUpdateproductRequest() *AlibabaScbpShowcaseUpdateproductRequest{
    return &AlibabaScbpShowcaseUpdateproductRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseUpdateproductRequest) GetApiMethodName() string {
    return "alibaba.scbp.showcase.updateproduct"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseUpdateproductRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WindowId Setter
// 橱窗id
func (r *AlibabaScbpShowcaseUpdateproductRequest) SetWindowId(windowId int64) error {
    r.windowId = windowId
    r.Set("window_id", windowId)
    return nil
}

// WindowId Getter
func (r AlibabaScbpShowcaseUpdateproductRequest) GetWindowId() int64 {
    return r.windowId
}
// NewProductId Setter
// 新的商品id
func (r *AlibabaScbpShowcaseUpdateproductRequest) SetNewProductId(newProductId int64) error {
    r.newProductId = newProductId
    r.Set("new_product_id", newProductId)
    return nil
}

// NewProductId Getter
func (r AlibabaScbpShowcaseUpdateproductRequest) GetNewProductId() int64 {
    return r.newProductId
}
