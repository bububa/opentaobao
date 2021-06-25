package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商铺货 APIRequest
tmall.supplychain.channel.product.release

供应商渠道铺货接口
*/
type TmallSupplychainChannelProductReleaseRequest struct {
    model.Params

    // 产品数字ID
    productId   int64 

    // 渠道ID
    channelCode   int64 

}

func NewTmallSupplychainChannelProductReleaseRequest() *TmallSupplychainChannelProductReleaseRequest{
    return &TmallSupplychainChannelProductReleaseRequest{
        Params: model.NewParams(),
    }
}

func (r TmallSupplychainChannelProductReleaseRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.release"
}

func (r TmallSupplychainChannelProductReleaseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallSupplychainChannelProductReleaseRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TmallSupplychainChannelProductReleaseRequest) GetProductId() int64 {
    return r.productId
}

func (r *TmallSupplychainChannelProductReleaseRequest) SetChannelCode(channelCode int64) error {
    r.channelCode = channelCode
    r.Set("channel_code", channelCode)
    return nil
}

func (r TmallSupplychainChannelProductReleaseRequest) GetChannelCode() int64 {
    return r.channelCode
}

