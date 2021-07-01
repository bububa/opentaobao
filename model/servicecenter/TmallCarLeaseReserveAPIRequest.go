package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseReserveAPIRequest
整车租车回传预约信息 API请求
tmall.car.lease.reserve

租赁公司回传预约到店信息 */
type TmallCarLeaseReserveAPIRequest struct {
	model.Params
	// 买家id
	_buyerId int64
	// 订单id
	_orderId int64
	// 文案
	_text string
	// 车架号
	_vin string
	// 1 代表 车辆到店，已预约用户到店提车   ; 2 车辆到店，未能联系到用户
	_flag int64
	// 买家昵称
	_buyerNick string
}

// New
