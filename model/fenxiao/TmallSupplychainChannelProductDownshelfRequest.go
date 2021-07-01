package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品下架 API请求
tmall.supplychain.channel.product.downshelf

产品下架
*/
type TmallSupplychainChannelProductDownshelfAPIRequest struct {
    model.Params
    // 产品ID
    _productId   int64
}

// 初始化TmallSupplychainChannelProductDownshelfAPIRequest对象
func NewTmallSupplychainChannelProductDownshelfRequest() *TmallSupplychainChannelProductDownshelfAPIRequest{
    return &TmallSupplychainChannelProductDownshelfAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductDownshelfAPIRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.downshelf"
}

// IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductDownshelfAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *TmallSupplychainChannelProductDownshelfAPIRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TmallSupplychainChannelProductDownshelfAPIRequest) GetProductId() int64 {
    return r._productId
}
