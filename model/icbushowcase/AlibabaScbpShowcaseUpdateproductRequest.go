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
type AlibabaScbpShowcaseUpdateproductAPIRequest struct {
    model.Params
    // 橱窗id
    _windowId   int64
    // 新的商品id
    _newProductId   int64
}

// 初始化AlibabaScbpShowcaseUpdateproductAPIRequest对象
func NewAlibabaScbpShowcaseUpdateproductRequest() *AlibabaScbpShowcaseUpdateproductAPIRequest{
    return &AlibabaScbpShowcaseUpdateproductAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseUpdateproductAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.showcase.updateproduct"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseUpdateproductAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WindowId Setter
// 橱窗id
func (r *AlibabaScbpShowcaseUpdateproductAPIRequest) SetWindowId(_windowId int64) error {
    r._windowId = _windowId
    r.Set("window_id", _windowId)
    return nil
}

// WindowId Getter
func (r AlibabaScbpShowcaseUpdateproductAPIRequest) GetWindowId() int64 {
    return r._windowId
}
// NewProductId Setter
// 新的商品id
func (r *AlibabaScbpShowcaseUpdateproductAPIRequest) SetNewProductId(_newProductId int64) error {
    r._newProductId = _newProductId
    r.Set("new_product_id", _newProductId)
    return nil
}

// NewProductId Getter
func (r AlibabaScbpShowcaseUpdateproductAPIRequest) GetNewProductId() int64 {
    return r._newProductId
}
