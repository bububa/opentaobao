package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoServindustryFinanceGeexOrderLoanAPIRequest 即科放款信息回调api API请求
// taobao.servindustry.finance.geex.order.loan
//
// 即科放款信息api
type TaobaoServindustryFinanceGeexOrderLoanAPIRequest struct {
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

// NewTaobaoServindustryFinanceGeexOrderLoanRequest 初始化TaobaoServindustryFinanceGeexOrderLoanAPIRequest对象
func NewTaobaoServindustryFinanceGeexOrderLoanRequest() *TaobaoServindustryFinanceGeexOrderLoanAPIRequest {
	return &TaobaoServindustryFinanceGeexOrderLoanAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoServindustryFinanceGeexOrderLoanAPIRequest) Reset() {
	r._priceUnit = ""
	r._alscOrderId = ""
	r._loanStatus = ""
	r._loanFlowId = ""
	r._updateTime = 0
	r._loanTime = 0
	r._loanPrice = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoServindustryFinanceGeexOrderLoanAPIRequest) GetApiMethodName() string {
	return "taobao.servindustry.finance.geex.order.loan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoServindustryFinanceGeexOrderLoanAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoServindustryFinanceGeexOrderLoanAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPriceUnit is PriceUnit Setter
// 金额单位
func (r *TaobaoServindustryFinanceGeexOrderLoanAPIRequest) SetPriceUnit(_priceUnit string) error {
	r._priceUnit = _priceUnit
	r.Set("price_unit", _priceUnit)
	return nil
}

// GetPriceUnit PriceUnit Getter
func (r TaobaoServindustryFinanceGeexOrderLoanAPIRequest) GetPriceUnit() string {
	return r._priceUnit
}

// SetAlscOrderId is AlscOrderId Setter
// 本地订单号
func (r *TaobaoServindustryFinanceGeexOrderLoanAPIRequest) SetAlscOrderId(_alscOrderId string) error {
	r._alscOrderId = _alscOrderId
	r.Set("alsc_order_id", _alscOrderId)
	return nil
}

// GetAlscOrderId AlscOrderId Getter
func (r TaobaoServindustryFinanceGeexOrderLoanAPIRequest) GetAlscOrderId() string {
	return r._alscOrderId
}

// SetLoanStatus is LoanStatus Setter
// 放款状态
func (r *TaobaoServindustryFinanceGeexOrderLoanAPIRequest) SetLoanStatus(_loanStatus string) error {
	r._loanStatus = _loanStatus
	r.Set("loan_status", _loanStatus)
	return nil
}

// GetLoanStatus LoanStatus Getter
func (r TaobaoServindustryFinanceGeexOrderLoanAPIRequest) GetLoanStatus() string {
	return r._loanStatus
}

// SetLoanFlowId is LoanFlowId Setter
// 放款流水
func (r *TaobaoServindustryFinanceGeexOrderLoanAPIRequest) SetLoanFlowId(_loanFlowId string) error {
	r._loanFlowId = _loanFlowId
	r.Set("loan_flow_id", _loanFlowId)
	return nil
}

// GetLoanFlowId LoanFlowId Getter
func (r TaobaoServindustryFinanceGeexOrderLoanAPIRequest) GetLoanFlowId() string {
	return r._loanFlowId
}

// SetUpdateTime is UpdateTime Setter
// 更新时间
func (r *TaobaoServindustryFinanceGeexOrderLoanAPIRequest) SetUpdateTime(_updateTime int64) error {
	r._updateTime = _updateTime
	r.Set("update_time", _updateTime)
	return nil
}

// GetUpdateTime UpdateTime Getter
func (r TaobaoServindustryFinanceGeexOrderLoanAPIRequest) GetUpdateTime() int64 {
	return r._updateTime
}

// SetLoanTime is LoanTime Setter
// 放款时间
func (r *TaobaoServindustryFinanceGeexOrderLoanAPIRequest) SetLoanTime(_loanTime int64) error {
	r._loanTime = _loanTime
	r.Set("loan_time", _loanTime)
	return nil
}

// GetLoanTime LoanTime Getter
func (r TaobaoServindustryFinanceGeexOrderLoanAPIRequest) GetLoanTime() int64 {
	return r._loanTime
}

// SetLoanPrice is LoanPrice Setter
// 放款金额
func (r *TaobaoServindustryFinanceGeexOrderLoanAPIRequest) SetLoanPrice(_loanPrice int64) error {
	r._loanPrice = _loanPrice
	r.Set("loan_price", _loanPrice)
	return nil
}

// GetLoanPrice LoanPrice Getter
func (r TaobaoServindustryFinanceGeexOrderLoanAPIRequest) GetLoanPrice() int64 {
	return r._loanPrice
}

var poolTaobaoServindustryFinanceGeexOrderLoanAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoServindustryFinanceGeexOrderLoanRequest()
	},
}

// GetTaobaoServindustryFinanceGeexOrderLoanRequest 从 sync.Pool 获取 TaobaoServindustryFinanceGeexOrderLoanAPIRequest
func GetTaobaoServindustryFinanceGeexOrderLoanAPIRequest() *TaobaoServindustryFinanceGeexOrderLoanAPIRequest {
	return poolTaobaoServindustryFinanceGeexOrderLoanAPIRequest.Get().(*TaobaoServindustryFinanceGeexOrderLoanAPIRequest)
}

// ReleaseTaobaoServindustryFinanceGeexOrderLoanAPIRequest 将 TaobaoServindustryFinanceGeexOrderLoanAPIRequest 放入 sync.Pool
func ReleaseTaobaoServindustryFinanceGeexOrderLoanAPIRequest(v *TaobaoServindustryFinanceGeexOrderLoanAPIRequest) {
	v.Reset()
	poolTaobaoServindustryFinanceGeexOrderLoanAPIRequest.Put(v)
}
