package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商铺货 API请求
tmall.supplychain.channel.product.release

供应商渠道铺货接口
*/
type TmallSupplychainChannelProductReleaseRequest struct {
    model.Params
    // 产品数字ID
    _productId   int64
    // 渠道ID
    _channelCode   int64
}

// 初始化TmallSupplychainChannelProductReleaseRequest对象
func NewTmallSupplychainChannelProductReleaseRequest() *TmallSupplychainChannelProductReleaseRequest{
    return &TmallSupplychainChannelProductReleaseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductReleaseRequest) GetApiMethodName() string {
    return "tmall.supplychain.channel.product.release"
}

// IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductReleaseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductId Setter
// 产品数字ID
func (r *TmallSupplychainChannelProductReleaseRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r TmallSupplychainChannelProductReleaseRequest) GetProductId() int64 {
    return r._productId
}
// ChannelCode Setter
// 渠道ID
func (r *TmallSupplychainChannelProductReleaseRequest) SetChannelCode(_channelCode int64) error {
    r._channelCode = _channelCode
    r.Set("channel_code", _channelCode)
    return nil
}

// ChannelCode Getter
func (r TmallSupplychainChannelProductReleaseRequest) GetChannelCode() int64 {
    return r._channelCode
}
