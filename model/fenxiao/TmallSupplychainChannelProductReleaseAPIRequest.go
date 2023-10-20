package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallsupplychainchannelproductreleaseAPIRequest 供应商铺货 API请求
// tmall.supplychain.channel.product.release
//
// 供应商渠道铺货接口
type TmallsupplychainchannelproductreleaseAPIRequest struct {
	model.Params
	// 产品数字ID
	_productId int64
	// 渠道ID
	_channelCode int64
}

// NewTmallsupplychainchannelproductreleaseRequest 初始化TmallsupplychainchannelproductreleaseAPIRequest对象
func NewTmallsupplychainchannelproductreleaseRequest() *TmallsupplychainchannelproductreleaseAPIRequest {
	return &TmallsupplychainchannelproductreleaseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallsupplychainchannelproductreleaseAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.release"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallsupplychainchannelproductreleaseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallsupplychainchannelproductreleaseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品数字ID
func (r *TmallsupplychainchannelproductreleaseAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallsupplychainchannelproductreleaseAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetChannelCode is ChannelCode Setter
// 渠道ID
func (r *TmallsupplychainchannelproductreleaseAPIRequest) SetChannelCode(_channelCode int64) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r TmallsupplychainchannelproductreleaseAPIRequest) GetChannelCode() int64 {
	return r._channelCode
}
