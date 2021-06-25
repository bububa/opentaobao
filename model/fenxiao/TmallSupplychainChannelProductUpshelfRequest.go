package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品上架 APIRequest
tmall.supplychain.channel.product.upshelf

上架渠道产品
*/
type TmallSupplychainChannelProductUpshelfRequest struct {
    model.Params

    // 产品ID
    productId   int64 

}

func NewTmallSupplychainChannelProductUpshelfRequest() *TmallSupplychainChannelProductUpshelfRequest{
    return &TmallSupplychainChannelProductUpshelfRequest{
        Params: model.NewParams(),
    }
}

func (r TmallSupplychainChannelProductUpshelfRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.upshelf"
}

func (r TmallSupplychainChannelProductUpshelfRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallSupplychainChannelProductUpshelfRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TmallSupplychainChannelProductUpshelfRequest) GetProductId() int64 {
    return r.productId
}

