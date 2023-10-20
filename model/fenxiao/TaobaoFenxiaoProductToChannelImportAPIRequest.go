package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproducttochannelimportAPIRequest 产品导入到渠道 API请求
// taobao.fenxiao.product.to.channel.import
//
// 支持供应商将已有产品导入到某个渠道销售
type TaobaofenxiaoproducttochannelimportAPIRequest struct {
	model.Params
	// 要导入的渠道[21 零售PLUS]目前仅支持此渠道
	_channel int64
	// 要导入的产品id
	_productId int64
}

// NewTaobaofenxiaoproducttochannelimportRequest 初始化TaobaofenxiaoproducttochannelimportAPIRequest对象
func NewTaobaofenxiaoproducttochannelimportRequest() *TaobaofenxiaoproducttochannelimportAPIRequest {
	return &TaobaofenxiaoproducttochannelimportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoproducttochannelimportAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.to.channel.import"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoproducttochannelimportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoproducttochannelimportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannel is Channel Setter
// 要导入的渠道[21 零售PLUS]目前仅支持此渠道
func (r *TaobaofenxiaoproducttochannelimportAPIRequest) SetChannel(_channel int64) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r TaobaofenxiaoproducttochannelimportAPIRequest) GetChannel() int64 {
	return r._channel
}

// SetProductId is ProductId Setter
// 要导入的产品id
func (r *TaobaofenxiaoproducttochannelimportAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaofenxiaoproducttochannelimportAPIRequest) GetProductId() int64 {
	return r._productId
}
