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
    _windowId   int64
    // 新的商品id
    _newProductId   int64
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
func (r *AlibabaScbpShowcaseUpdateproductRequest) SetWindowId(_windowId int64) error {
    r._windowId = _windowId
    r.Set("window_id", _windowId)
    return nil
}

// WindowId Getter
func (r AlibabaScbpShowcaseUpdateproductRequest) GetWindowId() int64 {
    return r._windowId
}
// NewProductId Setter
// 新的商品id
func (r *AlibabaScbpShowcaseUpdateproductRequest) SetNewProductId(_newProductId int64) error {
    r._newProductId = _newProductId
    r.Set("new_product_id", _newProductId)
    return nil
}

// NewProductId Getter
func (r AlibabaScbpShowcaseUpdateproductRequest) GetNewProductId() int64 {
    return r._newProductId
}
