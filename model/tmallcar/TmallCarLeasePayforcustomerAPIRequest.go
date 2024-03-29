package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeasePayforcustomerAPIRequest 天猫开新车租后代客户还款 API请求
// tmall.car.lease.payforcustomer
//
// 天猫开新车租后代客户还款
type TmallCarLeasePayforcustomerAPIRequest struct {
	model.Params
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
	// 天猫开新车订单id
	_orderId int64
}

// NewTmallCarLeasePayforcustomerRequest 初始化TmallCarLeasePayforcustomerAPIRequest对象
func NewTmallCarLeasePayforcustomerRequest() *TmallCarLeasePayforcustomerAPIRequest {
	return &TmallCarLeasePayforcustomerAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarLeasePayforcustomerAPIRequest) Reset() {
	r._custIproleId = ""
	r._date = ""
	r._loanArNo = ""
	r._prinAmt = ""
	r._requestId = ""
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeasePayforcustomerAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.payforcustomer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeasePayforcustomerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarLeasePayforcustomerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCustIproleId is CustIproleId Setter
// 贷款客户在网商的会员ID
func (r *TmallCarLeasePayforcustomerAPIRequest) SetCustIproleId(_custIproleId string) error {
	r._custIproleId = _custIproleId
	r.Set("cust_iprole_id", _custIproleId)
	return nil
}

// GetCustIproleId CustIproleId Getter
func (r TmallCarLeasePayforcustomerAPIRequest) GetCustIproleId() string {
	return r._custIproleId
}

// SetDate is Date Setter
// 还款日，精确到日，格式为yyyyMMdd，必须是当天
func (r *TmallCarLeasePayforcustomerAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r TmallCarLeasePayforcustomerAPIRequest) GetDate() string {
	return r._date
}

// SetLoanArNo is LoanArNo Setter
// 贷款合约号
func (r *TmallCarLeasePayforcustomerAPIRequest) SetLoanArNo(_loanArNo string) error {
	r._loanArNo = _loanArNo
	r.Set("loan_ar_no", _loanArNo)
	return nil
}

// GetLoanArNo LoanArNo Getter
func (r TmallCarLeasePayforcustomerAPIRequest) GetLoanArNo() string {
	return r._loanArNo
}

// SetPrinAmt is PrinAmt Setter
// 还款本金金额，单位默认为元，支持小数点两位，为了便于传输用合作方将数值型转换为字符串型
func (r *TmallCarLeasePayforcustomerAPIRequest) SetPrinAmt(_prinAmt string) error {
	r._prinAmt = _prinAmt
	r.Set("prin_amt", _prinAmt)
	return nil
}

// GetPrinAmt PrinAmt Getter
func (r TmallCarLeasePayforcustomerAPIRequest) GetPrinAmt() string {
	return r._prinAmt
}

// SetRequestId is RequestId Setter
// 外部流水号格式：日期(8位)+序列号(8位）,序列号是数字，如00000001（必须是16位且符合该格式
func (r *TmallCarLeasePayforcustomerAPIRequest) SetRequestId(_requestId string) error {
	r._requestId = _requestId
	r.Set("request_id", _requestId)
	return nil
}

// GetRequestId RequestId Getter
func (r TmallCarLeasePayforcustomerAPIRequest) GetRequestId() string {
	return r._requestId
}

// SetOrderId is OrderId Setter
// 天猫开新车订单id
func (r *TmallCarLeasePayforcustomerAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallCarLeasePayforcustomerAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolTmallCarLeasePayforcustomerAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarLeasePayforcustomerRequest()
	},
}

// GetTmallCarLeasePayforcustomerRequest 从 sync.Pool 获取 TmallCarLeasePayforcustomerAPIRequest
func GetTmallCarLeasePayforcustomerAPIRequest() *TmallCarLeasePayforcustomerAPIRequest {
	return poolTmallCarLeasePayforcustomerAPIRequest.Get().(*TmallCarLeasePayforcustomerAPIRequest)
}

// ReleaseTmallCarLeasePayforcustomerAPIRequest 将 TmallCarLeasePayforcustomerAPIRequest 放入 sync.Pool
func ReleaseTmallCarLeasePayforcustomerAPIRequest(v *TmallCarLeasePayforcustomerAPIRequest) {
	v.Reset()
	poolTmallCarLeasePayforcustomerAPIRequest.Put(v)
}
