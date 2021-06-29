package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品上架 API请求
tmall.supplychain.channel.product.upshelf

上架渠道产品
*/
type TmallSupplychainChannelProductUpshelfRequest struct {
    model.Params
    // 产品ID
    _productId   int64
}

// 初始化TmallSupplychainChannelProductUpshelfRequest对象
func NewTmallSupplychainChannelProductUpshelfRequest() *TmallSupplychainChannelProductUpshelfRequest{
    return &TmallSupplychainChannelProductUpshelfRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductUpshelfRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.upshelf"
}

// IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductUpshelfRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *TmallSupplychainChannelProductUpshelfRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TmallSupplychainChannelProductUpshelfRequest) GetProductId() int64 {
    return r._productId
}
