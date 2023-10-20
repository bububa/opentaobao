package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaxchannelskustatusupdateAPIRequest 翱象渠道商品上下架接口 API请求
// alibaba.ax.channel.sku.status.update
//
// 翱象渠道商品上下架接口
type AlibabaaxchannelskustatusupdateAPIRequest struct {
	model.Params
	// 请求入参
	_channelSkuUpdateStatusReq *ChannelSkuUpdateStatusReq
}

// NewAlibabaaxchannelskustatusupdateRequest 初始化AlibabaaxchannelskustatusupdateAPIRequest对象
func NewAlibabaaxchannelskustatusupdateRequest() *AlibabaaxchannelskustatusupdateAPIRequest {
	return &AlibabaaxchannelskustatusupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaxchannelskustatusupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ax.channel.sku.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaxchannelskustatusupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaxchannelskustatusupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelSkuUpdateStatusReq is ChannelSkuUpdateStatusReq Setter
// 请求入参
func (r *AlibabaaxchannelskustatusupdateAPIRequest) SetChannelSkuUpdateStatusReq(_channelSkuUpdateStatusReq *ChannelSkuUpdateStatusReq) error {
	r._channelSkuUpdateStatusReq = _channelSkuUpdateStatusReq
	r.Set("channel_sku_update_status_req", _channelSkuUpdateStatusReq)
	return nil
}

// GetChannelSkuUpdateStatusReq ChannelSkuUpdateStatusReq Getter
func (r AlibabaaxchannelskustatusupdateAPIRequest) GetChannelSkuUpdateStatusReq() *ChannelSkuUpdateStatusReq {
	return r._channelSkuUpdateStatusReq
}
