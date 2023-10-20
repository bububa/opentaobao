package miniapp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest 写入商家配置信息 API请求
// taobao.miniapp.ext.delivery.sell.channel.config.sync
//
// 写入商家配置信息
type TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest struct {
	model.Params
	// 入参
	_sellerChannelConfig *SellerChannelConfigDto
}

// NewTaobaoMiniappExtDeliverySellChannelConfigSyncRequest 初始化TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest对象
func NewTaobaoMiniappExtDeliverySellChannelConfigSyncRequest() *TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest {
	return &TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest) Reset() {
	r._sellerChannelConfig = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.ext.delivery.sell.channel.config.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerChannelConfig is SellerChannelConfig Setter
// 入参
func (r *TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest) SetSellerChannelConfig(_sellerChannelConfig *SellerChannelConfigDto) error {
	r._sellerChannelConfig = _sellerChannelConfig
	r.Set("seller_channel_config", _sellerChannelConfig)
	return nil
}

// GetSellerChannelConfig SellerChannelConfig Getter
func (r TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest) GetSellerChannelConfig() *SellerChannelConfigDto {
	return r._sellerChannelConfig
}

var poolTaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappExtDeliverySellChannelConfigSyncRequest()
	},
}

// GetTaobaoMiniappExtDeliverySellChannelConfigSyncRequest 从 sync.Pool 获取 TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest
func GetTaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest() *TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest {
	return poolTaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest.Get().(*TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest)
}

// ReleaseTaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest 将 TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest(v *TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest) {
	v.Reset()
	poolTaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest.Put(v)
}
