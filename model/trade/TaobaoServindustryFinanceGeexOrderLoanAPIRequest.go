package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoservindustryfinancegeexorderloanAPIRequest 即科放款信息回调api API请求
// taobao.servindustry.finance.geex.order.loan
//
// 即科放款信息api
type TaobaoservindustryfinancegeexorderloanAPIRequest struct {
	model.Params
	// 金额单位
	_priceUnit string
	// 本地订单号
	_alscOrderId string
	// 放款状态
	_loanStatus string
	// 放款流水
	_loanFlowId string
	// 更新时间
	_updateTime int64
	// 放款时间
	_loanTime int64
	// 放款金额
	_loanPrice int64
}

// NewTaobaoservindustryfinancegeexorderloanRequest 初始化TaobaoservindustryfinancegeexorderloanAPIRequest对象
func NewTaobaoservindustryfinancegeexorderloanRequest() *TaobaoservindustryfinancegeexorderloanAPIRequest {
	return &TaobaoservindustryfinancegeexorderloanAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoservindustryfinancegeexorderloanAPIRequest) GetApiMethodName() string {
	return "taobao.servindustry.finance.geex.order.loan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoservindustryfinancegeexorderloanAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoservindustryfinancegeexorderloanAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPriceUnit is PriceUnit Setter
// 金额单位
func (r *TaobaoservindustryfinancegeexorderloanAPIRequest) SetPriceUnit(_priceUnit string) error {
	r._priceUnit = _priceUnit
	r.Set("price_unit", _priceUnit)
	return nil
}

// GetPriceUnit PriceUnit Getter
func (r TaobaoservindustryfinancegeexorderloanAPIRequest) GetPriceUnit() string {
	return r._priceUnit
}

// SetAlscOrderId is AlscOrderId Setter
// 本地订单号
func (r *TaobaoservindustryfinancegeexorderloanAPIRequest) SetAlscOrderId(_alscOrderId string) error {
	r._alscOrderId = _alscOrderId
	r.Set("alsc_order_id", _alscOrderId)
	return nil
}

// GetAlscOrderId AlscOrderId Getter
func (r TaobaoservindustryfinancegeexorderloanAPIRequest) GetAlscOrderId() string {
	return r._alscOrderId
}

// SetLoanStatus is LoanStatus Setter
// 放款状态
func (r *TaobaoservindustryfinancegeexorderloanAPIRequest) SetLoanStatus(_loanStatus string) error {
	r._loanStatus = _loanStatus
	r.Set("loan_status", _loanStatus)
	return nil
}

// GetLoanStatus LoanStatus Getter
func (r TaobaoservindustryfinancegeexorderloanAPIRequest) GetLoanStatus() string {
	return r._loanStatus
}

// SetLoanFlowId is LoanFlowId Setter
// 放款流水
func (r *TaobaoservindustryfinancegeexorderloanAPIRequest) SetLoanFlowId(_loanFlowId string) error {
	r._loanFlowId = _loanFlowId
	r.Set("loan_flow_id", _loanFlowId)
	return nil
}

// GetLoanFlowId LoanFlowId Getter
func (r TaobaoservindustryfinancegeexorderloanAPIRequest) GetLoanFlowId() string {
	return r._loanFlowId
}

// SetUpdateTime is UpdateTime Setter
// 更新时间
func (r *TaobaoservindustryfinancegeexorderloanAPIRequest) SetUpdateTime(_updateTime int64) error {
	r._updateTime = _updateTime
	r.Set("update_time", _updateTime)
	return nil
}

// GetUpdateTime UpdateTime Getter
func (r TaobaoservindustryfinancegeexorderloanAPIRequest) GetUpdateTime() int64 {
	return r._updateTime
}

// SetLoanTime is LoanTime Setter
// 放款时间
func (r *TaobaoservindustryfinancegeexorderloanAPIRequest) SetLoanTime(_loanTime int64) error {
	r._loanTime = _loanTime
	r.Set("loan_time", _loanTime)
	return nil
}

// GetLoanTime LoanTime Getter
func (r TaobaoservindustryfinancegeexorderloanAPIRequest) GetLoanTime() int64 {
	return r._loanTime
}

// SetLoanPrice is LoanPrice Setter
// 放款金额
func (r *TaobaoservindustryfinancegeexorderloanAPIRequest) SetLoanPrice(_loanPrice int64) error {
	r._loanPrice = _loanPrice
	r.Set("loan_price", _loanPrice)
	return nil
}

// GetLoanPrice LoanPrice Getter
func (r TaobaoservindustryfinancegeexorderloanAPIRequest) GetLoanPrice() int64 {
	return r._loanPrice
}
