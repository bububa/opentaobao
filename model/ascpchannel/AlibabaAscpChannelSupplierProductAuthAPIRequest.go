package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelsupplierproductauthAPIRequest 供应商授权渠道产品到市场或分销商 API请求
// alibaba.ascp.channel.supplier.product.auth
//
// 供应商授权渠道产品到市场或分销商
type AlibabaascpchannelsupplierproductauthAPIRequest struct {
	model.Params
	// 请求参数
	_channelProductAuthRequest *ChannelProductAuthRequest
}

// NewAlibabaascpchannelsupplierproductauthRequest 初始化AlibabaascpchannelsupplierproductauthAPIRequest对象
func NewAlibabaascpchannelsupplierproductauthRequest() *AlibabaascpchannelsupplierproductauthAPIRequest {
	return &AlibabaascpchannelsupplierproductauthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpchannelsupplierproductauthAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.supplier.product.auth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpchannelsupplierproductauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpchannelsupplierproductauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelProductAuthRequest is ChannelProductAuthRequest Setter
// 请求参数
func (r *AlibabaascpchannelsupplierproductauthAPIRequest) SetChannelProductAuthRequest(_channelProductAuthRequest *ChannelProductAuthRequest) error {
	r._channelProductAuthRequest = _channelProductAuthRequest
	r.Set("channel_product_auth_request", _channelProductAuthRequest)
	return nil
}

// GetChannelProductAuthRequest ChannelProductAuthRequest Getter
func (r AlibabaascpchannelsupplierproductauthAPIRequest) GetChannelProductAuthRequest() *ChannelProductAuthRequest {
	return r._channelProductAuthRequest
}
