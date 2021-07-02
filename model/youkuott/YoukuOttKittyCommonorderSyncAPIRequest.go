package youkuott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttKittyCommonorderSyncAPIRequest 运营商一般订单同步 API请求
// youku.ott.kitty.commonorder.sync
//
// 运营商一般订单同步
type YoukuOttKittyCommonorderSyncAPIRequest struct {
	model.Params
	// 运营商订单id,最好是16位及以上唯一ID
	_orderId string
	// 充值的商品id（此商品需要事先给到优酷，并把商品的业务逻辑确定下来，比如是连续包月还是单月/单季/单年)
	_productId string
	// 同步时间 格式yyyy-MM-dd HH:mm:ss 说明：如果是线上或线下订单此时间是用户支付成功时间，如果是退订则是退订时间
	_syncTime string
	// 运营商渠道（需要找优酷方确认）
	_channelId string
	// 运营商用户账号账号id,与盒子登录账号tuid一致
	_accountId string
	// 订单类型 1:线上支付订单(线上应用内购买), 2:线下支付订单(比如营业厅订单), 3:连续包取消续订, 4:全额退款(立即终止权益,不分产品包,不计财务), 5:续费(运营商侧发起时才使用),6:非连续包退订(按未使用天数退款)
	_type string
	// 扩展字段，根据需要，约定具体的字段，json格式
	_extInfo string
}

// NewYoukuOttKittyCommonorderSyncRequest 初始化YoukuOttKittyCommonorderSyncAPIRequest对象
func NewYoukuOttKittyCommonorderSyncRequest() *YoukuOttKittyCommonorderSyncAPIRequest {
	return &YoukuOttKittyCommonorderSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttKittyCommonorderSyncAPIRequest) GetApiMethodName() string {
	return "youku.ott.kitty.commonorder.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttKittyCommonorderSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 运营商订单id,最好是16位及以上唯一ID
func (r *YoukuOttKittyCommonorderSyncAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r YoukuOttKittyCommonorderSyncAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetProductId is ProductId Setter
// 充值的商品id（此商品需要事先给到优酷，并把商品的业务逻辑确定下来，比如是连续包月还是单月/单季/单年)
func (r *YoukuOttKittyCommonorderSyncAPIRequest) SetProductId(_productId string) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r YoukuOttKittyCommonorderSyncAPIRequest) GetProductId() string {
	return r._productId
}

// SetSyncTime is SyncTime Setter
// 同步时间 格式yyyy-MM-dd HH:mm:ss 说明：如果是线上或线下订单此时间是用户支付成功时间，如果是退订则是退订时间
func (r *YoukuOttKittyCommonorderSyncAPIRequest) SetSyncTime(_syncTime string) error {
	r._syncTime = _syncTime
	r.Set("sync_time", _syncTime)
	return nil
}

// GetSyncTime SyncTime Getter
func (r YoukuOttKittyCommonorderSyncAPIRequest) GetSyncTime() string {
	return r._syncTime
}

// SetChannelId is ChannelId Setter
// 运营商渠道（需要找优酷方确认）
func (r *YoukuOttKittyCommonorderSyncAPIRequest) SetChannelId(_channelId string) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r YoukuOttKittyCommonorderSyncAPIRequest) GetChannelId() string {
	return r._channelId
}

// SetAccountId is AccountId Setter
// 运营商用户账号账号id,与盒子登录账号tuid一致
func (r *YoukuOttKittyCommonorderSyncAPIRequest) SetAccountId(_accountId string) error {
	r._accountId = _accountId
	r.Set("account_id", _accountId)
	return nil
}

// GetAccountId AccountId Getter
func (r YoukuOttKittyCommonorderSyncAPIRequest) GetAccountId() string {
	return r._accountId
}

// SetType is Type Setter
// 订单类型 1:线上支付订单(线上应用内购买), 2:线下支付订单(比如营业厅订单), 3:连续包取消续订, 4:全额退款(立即终止权益,不分产品包,不计财务), 5:续费(运营商侧发起时才使用),6:非连续包退订(按未使用天数退款)
func (r *YoukuOttKittyCommonorderSyncAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r YoukuOttKittyCommonorderSyncAPIRequest) GetType() string {
	return r._type
}

// SetExtInfo is ExtInfo Setter
// 扩展字段，根据需要，约定具体的字段，json格式
func (r *YoukuOttKittyCommonorderSyncAPIRequest) SetExtInfo(_extInfo string) error {
	r._extInfo = _extInfo
	r.Set("ext_info", _extInfo)
	return nil
}

// GetExtInfo ExtInfo Getter
func (r YoukuOttKittyCommonorderSyncAPIRequest) GetExtInfo() string {
	return r._extInfo
}
