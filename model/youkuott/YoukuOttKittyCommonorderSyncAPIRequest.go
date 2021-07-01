package youkuott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttKittyCommonorderSyncAPIRequest
运营商一般订单同步 API请求
youku.ott.kitty.commonorder.sync

运营商一般订单同步 */
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

// New
