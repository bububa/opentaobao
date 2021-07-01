package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeasePayforcustomerAPIRequest
天猫开新车租后代客户还款 API请求
tmall.car.lease.payforcustomer

天猫开新车租后代客户还款 */
type TmallCarLeasePayforcustomerAPIRequest struct {
	model.Params
	// 天猫开新车订单id
	_orderId int64
	// 贷款客户在网商的会员ID
	_custIproleId string
	// 还款日，精确到日，格式为yyyyMMdd，必须是当天
	_date string
	// 贷款合约号
	_loanArNo string
	// 还款本金金额，单位默认为元，支持小数点两位，为了便于传输用合作方将数值型转换为字符串型
	_prinAmt string
	// 外部流水号格式：日期(8位)+序列号(8位）,序列号是数字，如00000001（必须是16位且符合该格式
	_requestId string
}

// New
