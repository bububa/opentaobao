package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品下架 APIRequest
tmall.supplychain.channel.product.downshelf

产品下架
*/
type TmallSupplychainChannelProductDownshelfRequest struct {
    model.Params

    // 产品ID
    productId   int64 

}

func NewTmallSupplychainChannelProductDownshelfRequest() *TmallSupplychainChannelProductDownshelfRequest{
    return &TmallSupplychainChannelProductDownshelfRequest{
        Params: model.NewParams(),
    }
}

func (r TmallSupplychainChannelProductDownshelfRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.downshelf"
}

func (r TmallSupplychainChannelProductDownshelfRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallSupplychainChannelProductDownshelfRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TmallSupplychainChannelProductDownshelfRequest) GetProductId() int64 {
    return r.productId
}

