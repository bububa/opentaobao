package xhoteloffline

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderAlipayfaceCreateAPIRequest
信用住支付创建接口 API请求
taobao.xhotel.order.alipayface.create

用于创建一笔信用住支付，主要应用场景是线下信用住 */
type TaobaoXhotelOrderAlipayfaceCreateAPIRequest struct {
	model.Params
	// 总房费,单位为分
	_totalFee int64
	// 离店日期(最多允许9天)
	_checkOut string
	// 发布到阿里旅行的酒店编码
	_hotelCode string
	// 入住日期
	_checkIn string
	// 每日房价,json格式
	_dailyPriceInfo string
	// 商家系统的订单号，必须全局唯一，重复会按照相同订单处理
	_outOrderId string
	// 预定的房间数量
	_roomQuantity int64
	// 入住人信息, 注意必须有且只有一个设置为主入住人, 用于信用住结算扣款. 对于java版本的SDK可以使用setGuests(List &lt Guest &gt guests)赋值; 对于.net等其他版本SDK可以通过将List &lt Guest&gt结构数据转为json串赋值.
	_guests []Guest
	// 扫描用户支付宝得到的串号, 该字段不为空时会采用串号对应的支付宝账号进行信用住结算
	_alipayNumber string
	// 订单渠道信息,可以留空
	_channel string
	// 不清楚请留空, 用于和outHid共同定位一个酒店
	_vendor string
	// 房型名称
	_roomtypeName string
	// rateplan名称(不清楚可以留空)
	_rateplanName string
	// 是否为自助入住模式下创建订单，是：true,否：false
	_selfCheckin bool
}

// New
