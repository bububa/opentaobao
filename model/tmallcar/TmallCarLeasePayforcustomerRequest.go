package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车租后代客户还款 API请求
tmall.car.lease.payforcustomer

天猫开新车租后代客户还款
*/
type TmallCarLeasePayforcustomerRequest struct {
    model.Params
    // 天猫开新车订单id
    _orderId   int64
    // 贷款客户在网商的会员ID
    _custIproleId   string
    // 还款日，精确到日，格式为yyyyMMdd，必须是当天
    _date   string
    // 贷款合约号
    _loanArNo   string
    // 还款本金金额，单位默认为元，支持小数点两位，为了便于传输用合作方将数值型转换为字符串型
    _prinAmt   string
    // 外部流水号格式：日期(8位)+序列号(8位）,序列号是数字，如00000001（必须是16位且符合该格式
    _requestId   string
}

// 初始化TmallCarLeasePayforcustomerRequest对象
func NewTmallCarLeasePayforcustomerRequest() *TmallCarLeasePayforcustomerRequest{
    return &TmallCarLeasePayforcustomerRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeasePayforcustomerRequest) GetApiMethodName() string {
    return "tmall.car.lease.payforcustomer"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeasePayforcustomerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 天猫开新车订单id
func (r *TmallCarLeasePayforcustomerRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TmallCarLeasePayforcustomerRequest) GetOrderId() int64 {
    return r._orderId
}
// CustIproleId Setter
// 贷款客户在网商的会员ID
func (r *TmallCarLeasePayforcustomerRequest) SetCustIproleId(_custIproleId string) error {
    r._custIproleId = _custIproleId
    r.Set("cust_iprole_id", _custIproleId)
    return nil
}

// CustIproleId Getter
func (r TmallCarLeasePayforcustomerRequest) GetCustIproleId() string {
    return r._custIproleId
}
// Date Setter
// 还款日，精确到日，格式为yyyyMMdd，必须是当天
func (r *TmallCarLeasePayforcustomerRequest) SetDate(_date string) error {
    r._date = _date
    r.Set("date", _date)
    return nil
}

// Date Getter
func (r TmallCarLeasePayforcustomerRequest) GetDate() string {
    return r._date
}
// LoanArNo Setter
// 贷款合约号
func (r *TmallCarLeasePayforcustomerRequest) SetLoanArNo(_loanArNo string) error {
    r._loanArNo = _loanArNo
    r.Set("loan_ar_no", _loanArNo)
    return nil
}

// LoanArNo Getter
func (r TmallCarLeasePayforcustomerRequest) GetLoanArNo() string {
    return r._loanArNo
}
// PrinAmt Setter
// 还款本金金额，单位默认为元，支持小数点两位，为了便于传输用合作方将数值型转换为字符串型
func (r *TmallCarLeasePayforcustomerRequest) SetPrinAmt(_prinAmt string) error {
    r._prinAmt = _prinAmt
    r.Set("prin_amt", _prinAmt)
    return nil
}

// PrinAmt Getter
func (r TmallCarLeasePayforcustomerRequest) GetPrinAmt() string {
    return r._prinAmt
}
// RequestId Setter
// 外部流水号格式：日期(8位)+序列号(8位）,序列号是数字，如00000001（必须是16位且符合该格式
func (r *TmallCarLeasePayforcustomerRequest) SetRequestId(_requestId string) error {
    r._requestId = _requestId
    r.Set("request_id", _requestId)
    return nil
}

// RequestId Getter
func (r TmallCarLeasePayforcustomerRequest) GetRequestId() string {
    return r._requestId
}
