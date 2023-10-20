package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallsupplychainchannelproductreleasestatusgetAPIRequest 产品铺货状态查询 API请求
// tmall.supplychain.channel.product.release.status.get
//
// 巴拿马战役渠道产品状态查询
type TmallsupplychainchannelproductreleasestatusgetAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// 渠道ID ( 台湾 : 111002 )
	_channelCode int64
}

// NewTmallsupplychainchannelproductreleasestatusgetRequest 初始化TmallsupplychainchannelproductreleasestatusgetAPIRequest对象
func NewTmallsupplychainchannelproductreleasestatusgetRequest() *TmallsupplychainchannelproductreleasestatusgetAPIRequest {
	return &TmallsupplychainchannelproductreleasestatusgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallsupplychainchannelproductreleasestatusgetAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.release.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallsupplychainchannelproductreleasestatusgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallsupplychainchannelproductreleasestatusgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TmallsupplychainchannelproductreleasestatusgetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallsupplychainchannelproductreleasestatusgetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetChannelCode is ChannelCode Setter
// 渠道ID ( 台湾 : 111002 )
func (r *TmallsupplychainchannelproductreleasestatusgetAPIRequest) SetChannelCode(_channelCode int64) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r TmallsupplychainchannelproductreleasestatusgetAPIRequest) GetChannelCode() int64 {
	return r._channelCode
}
