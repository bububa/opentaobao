package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品铺货状态查询 APIRequest
tmall.supplychain.channel.product.release.status.get

巴拿马战役渠道产品状态查询
*/
type TmallSupplychainChannelProductReleaseStatusGetRequest struct {
    model.Params

    // 产品ID
    productId   int64 

    // 渠道ID ( 台湾 : 111002 )
    channelCode   int64 

}

func NewTmallSupplychainChannelProductReleaseStatusGetRequest() *TmallSupplychainChannelProductReleaseStatusGetRequest{
    return &TmallSupplychainChannelProductReleaseStatusGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallSupplychainChannelProductReleaseStatusGetRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.release.status.get"
}

func (r TmallSupplychainChannelProductReleaseStatusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallSupplychainChannelProductReleaseStatusGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TmallSupplychainChannelProductReleaseStatusGetRequest) GetProductId() int64 {
    return r.productId
}

func (r *TmallSupplychainChannelProductReleaseStatusGetRequest) SetChannelCode(channelCode int64) error {
    r.channelCode = channelCode
    r.Set("channel_code", channelCode)
    return nil
}

func (r TmallSupplychainChannelProductReleaseStatusGetRequest) GetChannelCode() int64 {
    return r.channelCode
}

