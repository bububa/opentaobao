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
type TmallSupplychainChannelProductDownshelfRequest struct {
    model.Params
    // 产品ID
    productId   int64
}

// 初始化TmallSupplychainChannelProductDownshelfRequest对象
func NewTmallSupplychainChannelProductDownshelfRequest() *TmallSupplychainChannelProductDownshelfRequest{
    return &TmallSupplychainChannelProductDownshelfRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductDownshelfRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.downshelf"
}

// IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductDownshelfRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *TmallSupplychainChannelProductDownshelfRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TmallSupplychainChannelProductDownshelfRequest) GetProductId() int64 {
    return r.productId
}
