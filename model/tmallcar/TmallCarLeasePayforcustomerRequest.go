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
    orderId   int64
    // 贷款客户在网商的会员ID
    custIproleId   string
    // 还款日，精确到日，格式为yyyyMMdd，必须是当天
    date   string
    // 贷款合约号
    loanArNo   string
    // 还款本金金额，单位默认为元，支持小数点两位，为了便于传输用合作方将数值型转换为字符串型
    prinAmt   string
    // 外部流水号格式：日期(8位)+序列号(8位）,序列号是数字，如00000001（必须是16位且符合该格式
    requestId   string
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
func (r *TmallCarLeasePayforcustomerRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TmallCarLeasePayforcustomerRequest) GetOrderId() int64 {
    return r.orderId
}
// CustIproleId Setter
// 贷款客户在网商的会员ID
func (r *TmallCarLeasePayforcustomerRequest) SetCustIproleId(custIproleId string) error {
    r.custIproleId = custIproleId
    r.Set("cust_iprole_id", custIproleId)
    return nil
}

// CustIproleId Getter
func (r TmallCarLeasePayforcustomerRequest) GetCustIproleId() string {
    return r.custIproleId
}
// Date Setter
// 还款日，精确到日，格式为yyyyMMdd，必须是当天
func (r *TmallCarLeasePayforcustomerRequest) SetDate(date string) error {
    r.date = date
    r.Set("date", date)
    return nil
}

// Date Getter
func (r TmallCarLeasePayforcustomerRequest) GetDate() string {
    return r.date
}
// LoanArNo Setter
// 贷款合约号
func (r *TmallCarLeasePayforcustomerRequest) SetLoanArNo(loanArNo string) error {
    r.loanArNo = loanArNo
    r.Set("loan_ar_no", loanArNo)
    return nil
}

// LoanArNo Getter
func (r TmallCarLeasePayforcustomerRequest) GetLoanArNo() string {
    return r.loanArNo
}
// PrinAmt Setter
// 还款本金金额，单位默认为元，支持小数点两位，为了便于传输用合作方将数值型转换为字符串型
func (r *TmallCarLeasePayforcustomerRequest) SetPrinAmt(prinAmt string) error {
    r.prinAmt = prinAmt
    r.Set("prin_amt", prinAmt)
    return nil
}

// PrinAmt Getter
func (r TmallCarLeasePayforcustomerRequest) GetPrinAmt() string {
    return r.prinAmt
}
// RequestId Setter
// 外部流水号格式：日期(8位)+序列号(8位）,序列号是数字，如00000001（必须是16位且符合该格式
func (r *TmallCarLeasePayforcustomerRequest) SetRequestId(requestId string) error {
    r.requestId = requestId
    r.Set("request_id", requestId)
    return nil
}

// RequestId Getter
func (r TmallCarLeasePayforcustomerRequest) GetRequestId() string {
    return r.requestId
}
