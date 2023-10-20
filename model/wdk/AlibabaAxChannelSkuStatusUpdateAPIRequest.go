package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAxChannelSkuStatusUpdateAPIRequest 翱象渠道商品上下架接口 API请求
// alibaba.ax.channel.sku.status.update
//
// 翱象渠道商品上下架接口
type AlibabaAxChannelSkuStatusUpdateAPIRequest struct {
	model.Params
	// 请求入参
	_channelSkuUpdateStatusReq *ChannelSkuUpdateStatusReq
}

// NewAlibabaAxChannelSkuStatusUpdateRequest 初始化AlibabaAxChannelSkuStatusUpdateAPIRequest对象
func NewAlibabaAxChannelSkuStatusUpdateRequest() *AlibabaAxChannelSkuStatusUpdateAPIRequest {
	return &AlibabaAxChannelSkuStatusUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAxChannelSkuStatusUpdateAPIRequest) Reset() {
	r._channelSkuUpdateStatusReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAxChannelSkuStatusUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ax.channel.sku.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAxChannelSkuStatusUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAxChannelSkuStatusUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelSkuUpdateStatusReq is ChannelSkuUpdateStatusReq Setter
// 请求入参
func (r *AlibabaAxChannelSkuStatusUpdateAPIRequest) SetChannelSkuUpdateStatusReq(_channelSkuUpdateStatusReq *ChannelSkuUpdateStatusReq) error {
	r._channelSkuUpdateStatusReq = _channelSkuUpdateStatusReq
	r.Set("channel_sku_update_status_req", _channelSkuUpdateStatusReq)
	return nil
}

// GetChannelSkuUpdateStatusReq ChannelSkuUpdateStatusReq Getter
func (r AlibabaAxChannelSkuStatusUpdateAPIRequest) GetChannelSkuUpdateStatusReq() *ChannelSkuUpdateStatusReq {
	return r._channelSkuUpdateStatusReq
}

var poolAlibabaAxChannelSkuStatusUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAxChannelSkuStatusUpdateRequest()
	},
}

// GetAlibabaAxChannelSkuStatusUpdateRequest 从 sync.Pool 获取 AlibabaAxChannelSkuStatusUpdateAPIRequest
func GetAlibabaAxChannelSkuStatusUpdateAPIRequest() *AlibabaAxChannelSkuStatusUpdateAPIRequest {
	return poolAlibabaAxChannelSkuStatusUpdateAPIRequest.Get().(*AlibabaAxChannelSkuStatusUpdateAPIRequest)
}

// ReleaseAlibabaAxChannelSkuStatusUpdateAPIRequest 将 AlibabaAxChannelSkuStatusUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAxChannelSkuStatusUpdateAPIRequest(v *AlibabaAxChannelSkuStatusUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAxChannelSkuStatusUpdateAPIRequest.Put(v)
}
