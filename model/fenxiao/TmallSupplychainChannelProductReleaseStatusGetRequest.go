package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
产品铺货状态查询 API请求
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

// 初始化TmallSupplychainChannelProductReleaseStatusGetRequest对象
func NewTmallSupplychainChannelProductReleaseStatusGetRequest() *TmallSupplychainChannelProductReleaseStatusGetRequest{
    return &TmallSupplychainChannelProductReleaseStatusGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductReleaseStatusGetRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.release.status.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductReleaseStatusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品ID
func (r *TmallSupplychainChannelProductReleaseStatusGetRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TmallSupplychainChannelProductReleaseStatusGetRequest) GetProductId() int64 {
    return r.productId
}
// ChannelCode Setter
// 渠道ID ( 台湾 : 111002 )
func (r *TmallSupplychainChannelProductReleaseStatusGetRequest) SetChannelCode(channelCode int64) error {
    r.channelCode = channelCode
    r.Set("channel_code", channelCode)
    return nil
}

// ChannelCode Getter
func (r TmallSupplychainChannelProductReleaseStatusGetRequest) GetChannelCode() int64 {
    return r.channelCode
}
