package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallSupplychainChannelProductReleaseStatusGetAPIRequest 产品铺货状态查询 API请求
// tmall.supplychain.channel.product.release.status.get
//
// 巴拿马战役渠道产品状态查询
type TmallSupplychainChannelProductReleaseStatusGetAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// 渠道ID ( 台湾 : 111002 )
	_channelCode int64
}

// NewTmallSupplychainChannelProductReleaseStatusGetRequest 初始化TmallSupplychainChannelProductReleaseStatusGetAPIRequest对象
func NewTmallSupplychainChannelProductReleaseStatusGetRequest() *TmallSupplychainChannelProductReleaseStatusGetAPIRequest {
	return &TmallSupplychainChannelProductReleaseStatusGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductReleaseStatusGetAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.release.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductReleaseStatusGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TmallSupplychainChannelProductReleaseStatusGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallSupplychainChannelProductReleaseStatusGetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetChannelCode is ChannelCode Setter
// 渠道ID ( 台湾 : 111002 )
func (r *TmallSupplychainChannelProductReleaseStatusGetAPIRequest) SetChannelCode(_channelCode int64) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r TmallSupplychainChannelProductReleaseStatusGetAPIRequest) GetChannelCode() int64 {
	return r._channelCode
}
