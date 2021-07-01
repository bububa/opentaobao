package xhoteloffline

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderAlipayfaceExtendAPIRequest
信用住订单延住接口 API请求
taobao.xhotel.order.alipayface.extend

信用住订单延住接口，用于将已有的信用住支付订单checkOut和订单金额等更新 */
type TaobaoXhotelOrderAlipayfaceExtendAPIRequest struct {
	model.Params
	// 延住后的离店日期(最多总共入住天数不能超过9间夜)
	_checkOut string
	// 阿里旅行订单id,和outOrderId必须至少传入一个
	_tid int64
	// 卖家系统订单号,和tid必须至少传入一个
	_outOrderId string
	// 延住房费,注意不包含原有的房费金额,单位为分
	_extendFee int64
	// 延住日期段的每日房价信息,注意不包括原有的日期房价
	_extendDailyPriceInfo string
}

// New
