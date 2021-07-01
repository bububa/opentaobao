package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseStatussynchronizeAPIRequest
天猫开新车租后状态同步 API请求
tmall.car.lease.statussynchronize

天猫开新车租后状态同步 */
type TmallCarLeaseStatussynchronizeAPIRequest struct {
	model.Params
	// 天猫开新车订单号
	_orderId int64
	// 业务类型:0.退车,1.买断,2.分期，3.续租
	_bizType int64
	// 1:过户,2:续租,3.额外费用审核，4.退车完成
	_actionType int64
	// 1:通过，-1:拒绝
	_actionValue int64
	// 拒绝原因
	_refuseReason string
}

// New
