package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderAlipayfaceUpdateAPIRequest
酒店信用住订单状态更新 API请求
taobao.xhotel.order.alipayface.update

完成对信用住或者面付订单的状态的更新。包含订单状态的确认，入离店状态的更新等等。（不适用于预付订单） */
type TaobaoXhotelOrderAlipayfaceUpdateAPIRequest struct {
	model.Params
	// 淘宝订单号,必填
	_tid int64
	// 操作的类型：12.补录确认号,11.多间房确认无房，10.多间房确认有房，8.取消订单(cancel)酒店端发起取消,必须在和买家协商通过的情况下操作,否则有法务风险; 5.买家未入住(noshow),如果该单有担保，会收取买家的担保金额; 3.核实入住(checkIn); 4.核实离店(checkOut); 1.确认无房(直连卖家禁止该操作),2.确认有房(直连卖家禁止该操作)
	_optType int64
	// 无房原因分类:1.无房, 2.价格变动, 3.买家原因, 4.其它原因,opt_type=1时必填
	_reasonType int64
	// 无房原因描述:opt_type=1时必填
	_reasonText string
	// 入住房间号
	_outRoomNumber string
	// 客人实际入住日期,opt_type=3/4时必填
	_checkinDate string
	// 客人实际离店日期,opt_type=4时必填
	_checkoutDate string
	// 客人实际预定房间数
	_rooms int64
	// 外部订单号
	_outId string
	// opt_type为10,11启用，多间房订单号列表，逗号间隔
	_tids string
	// opt_type为11启用，多间房订单取消原因类型，逗号间隔
	_cancelType int64
	// opt_type为10,11,12启用，真实操作人
	_operator string
	// opt_type为12, 订单确认号
	_confirmCode string
	// 是否自助入住
	_selfCheckin bool
	// 是否把代理直签的订单同步到酒店，Y为同步，N不同步
	_syncToHotel string
}

// New
