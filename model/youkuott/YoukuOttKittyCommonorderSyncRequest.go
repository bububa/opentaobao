package youkuott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
运营商一般订单同步 API请求
youku.ott.kitty.commonorder.sync

运营商一般订单同步
*/
type YoukuOttKittyCommonorderSyncRequest struct {
    model.Params
    // 运营商订单id,最好是16位及以上唯一ID
    orderId   string
    // 充值的商品id（此商品需要事先给到优酷，并把商品的业务逻辑确定下来，比如是连续包月还是单月/单季/单年)
    productId   string
    // 同步时间 格式yyyy-MM-dd HH:mm:ss 说明：如果是线上或线下订单此时间是用户支付成功时间，如果是退订则是退订时间
    syncTime   string
    // 运营商渠道（需要找优酷方确认）
    channelId   string
    // 运营商用户账号账号id,与盒子登录账号tuid一致
    accountId   string
    // 订单类型 1:线上支付订单(线上应用内购买), 2:线下支付订单(比如营业厅订单), 3:连续包取消续订, 4:全额退款(立即终止权益,不分产品包,不计财务), 5:续费(运营商侧发起时才使用),6:非连续包退订(按未使用天数退款)
    type   string
    // 扩展字段，根据需要，约定具体的字段，json格式
    extInfo   string
}

// 初始化YoukuOttKittyCommonorderSyncRequest对象
func NewYoukuOttKittyCommonorderSyncRequest() *YoukuOttKittyCommonorderSyncRequest{
    return &YoukuOttKittyCommonorderSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttKittyCommonorderSyncRequest) GetApiMethodName() string {
    return "youku.ott.kitty.commonorder.sync"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttKittyCommonorderSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 运营商订单id,最好是16位及以上唯一ID
func (r *YoukuOttKittyCommonorderSyncRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r YoukuOttKittyCommonorderSyncRequest) GetOrderId() string {
    return r.orderId
}
// ProductId Setter
// 充值的商品id（此商品需要事先给到优酷，并把商品的业务逻辑确定下来，比如是连续包月还是单月/单季/单年)
func (r *YoukuOttKittyCommonorderSyncRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r YoukuOttKittyCommonorderSyncRequest) GetProductId() string {
    return r.productId
}
// SyncTime Setter
// 同步时间 格式yyyy-MM-dd HH:mm:ss 说明：如果是线上或线下订单此时间是用户支付成功时间，如果是退订则是退订时间
func (r *YoukuOttKittyCommonorderSyncRequest) SetSyncTime(syncTime string) error {
    r.syncTime = syncTime
    r.Set("sync_time", syncTime)
    return nil
}

// SyncTime Getter
func (r YoukuOttKittyCommonorderSyncRequest) GetSyncTime() string {
    return r.syncTime
}
// ChannelId Setter
// 运营商渠道（需要找优酷方确认）
func (r *YoukuOttKittyCommonorderSyncRequest) SetChannelId(channelId string) error {
    r.channelId = channelId
    r.Set("channel_id", channelId)
    return nil
}

// ChannelId Getter
func (r YoukuOttKittyCommonorderSyncRequest) GetChannelId() string {
    return r.channelId
}
// AccountId Setter
// 运营商用户账号账号id,与盒子登录账号tuid一致
func (r *YoukuOttKittyCommonorderSyncRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r YoukuOttKittyCommonorderSyncRequest) GetAccountId() string {
    return r.accountId
}
// Type Setter
// 订单类型 1:线上支付订单(线上应用内购买), 2:线下支付订单(比如营业厅订单), 3:连续包取消续订, 4:全额退款(立即终止权益,不分产品包,不计财务), 5:续费(运营商侧发起时才使用),6:非连续包退订(按未使用天数退款)
func (r *YoukuOttKittyCommonorderSyncRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r YoukuOttKittyCommonorderSyncRequest) GetType() string {
    return r.type
}
// ExtInfo Setter
// 扩展字段，根据需要，约定具体的字段，json格式
func (r *YoukuOttKittyCommonorderSyncRequest) SetExtInfo(extInfo string) error {
    r.extInfo = extInfo
    r.Set("ext_info", extInfo)
    return nil
}

// ExtInfo Getter
func (r YoukuOttKittyCommonorderSyncRequest) GetExtInfo() string {
    return r.extInfo
}
