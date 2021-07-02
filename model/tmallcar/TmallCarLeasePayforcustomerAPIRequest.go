package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeasePayforcustomerAPIRequest 天猫开新车租后代客户还款 API请求
// tmall.car.lease.payforcustomer
//
// 天猫开新车租后代客户还款
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

// NewTmallCarLeasePayforcustomerRequest 初始化TmallCarLeasePayforcustomerAPIRequest对象
func NewTmallCarLeasePayforcustomerRequest() *TmallCarLeasePayforcustomerAPIRequest {
	return &TmallCarLeasePayforcustomerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeasePayforcustomerAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.payforcustomer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeasePayforcustomerAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 天猫开新车订单id
func (r *TmallCarLeasePayforcustomerAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r TmallCarLeasePayforcustomerAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// Set is CustIproleId Setter
// 贷款客户在网商的会员ID
func (r *TmallCarLeasePayforcustomerAPIRequest) SetCustIproleId(_custIproleId string) error {
	r._custIproleId = _custIproleId
	r.Set("cust_iprole_id", _custIproleId)
	return nil
}

// Get CustIproleId Getter
func (r TmallCarLeasePayforcustomerAPIRequest) GetCustIproleId() string {
	return r._custIproleId
}

// Set is Date Setter
// 还款日，精确到日，格式为yyyyMMdd，必须是当天
func (r *TmallCarLeasePayforcustomerAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// Get Date Getter
func (r TmallCarLeasePayforcustomerAPIRequest) GetDate() string {
	return r._date
}

// Set is LoanArNo Setter
// 贷款合约号
func (r *TmallCarLeasePayforcustomerAPIRequest) SetLoanArNo(_loanArNo string) error {
	r._loanArNo = _loanArNo
	r.Set("loan_ar_no", _loanArNo)
	return nil
}

// Get LoanArNo Getter
func (r TmallCarLeasePayforcustomerAPIRequest) GetLoanArNo() string {
	return r._loanArNo
}

// Set is PrinAmt Setter
// 还款本金金额，单位默认为元，支持小数点两位，为了便于传输用合作方将数值型转换为字符串型
func (r *TmallCarLeasePayforcustomerAPIRequest) SetPrinAmt(_prinAmt string) error {
	r._prinAmt = _prinAmt
	r.Set("prin_amt", _prinAmt)
	return nil
}

// Get PrinAmt Getter
func (r TmallCarLeasePayforcustomerAPIRequest) GetPrinAmt() string {
	return r._prinAmt
}

// Set is RequestId Setter
// 外部流水号格式：日期(8位)+序列号(8位）,序列号是数字，如00000001（必须是16位且符合该格式
func (r *TmallCarLeasePayforcustomerAPIRequest) SetRequestId(_requestId string) error {
	r._requestId = _requestId
	r.Set("request_id", _requestId)
	return nil
}

// Get RequestId Getter
func (r TmallCarLeasePayforcustomerAPIRequest) GetRequestId() string {
	return r._requestId
}
