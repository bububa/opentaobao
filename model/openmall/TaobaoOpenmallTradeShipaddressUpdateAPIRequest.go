package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallTradeShipaddressUpdateAPIRequest
Openmall订单收货地址修改 API请求
taobao.openmall.trade.shipaddress.update

Openmall订单收货地址修改 */
type TaobaoOpenmallTradeShipaddressUpdateAPIRequest struct {
	model.Params
	// 媒体渠道，代表分销者的身份，签约支付宝代扣的渠道商淘宝账号nick
	_distributor string
	// 收货地址。最大长度为228个字节。
	_receiverAddress string
	// 城市。最大长度为32个字节。如：杭州
	_receiverCity string
	// 区/县。最大长度为32个字节。如：西湖区
	_receiverDistrict string
	// 移动电话。最大长度为11个字节。
	_receiverMobile string
	// 收货人全名。最大长度为50个字节。
	_receiverName string
	// 固定电话。最大长度为30个字节。
	_receiverPhone string
	// 省份。最大长度为32个字节。如：浙江
	_receiverState string
	// 6位数字的邮编
	_receiverZip string
	// 淘宝订单号
	_tid int64
}

// New
