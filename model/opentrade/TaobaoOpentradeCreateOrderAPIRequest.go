package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeCreateOrderAPIRequest
订单创建 API请求
taobao.opentrade.create.order

交易开放创建订单 */
type TaobaoOpentradeCreateOrderAPIRequest struct {
	model.Params
	// 外部订单ID
	_outId string
	// 买家openID
	_openUserId string
	// 收货地址的收件人姓名
	_fullName string
	// 收货地址的手机号码
	_mobile string
	// 收货地址
	_address string
	// 卖家备忘
	_sellerMemo string
	// 卖家备忘
	_buyerMemo string
	// 商品信息，一次不能超过10个
	_itemInfos []ItemInfo
}

// New
