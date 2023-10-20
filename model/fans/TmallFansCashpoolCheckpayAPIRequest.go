package fans

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallfanscashpoolcheckpayAPIRequest 检查资金池付款状态 API请求
// tmall.fans.cashpool.checkpay
//
// 检查资金池付款状态
type TmallfanscashpoolcheckpayAPIRequest struct {
	model.Params
	// 资金池列表
	_cashPoolList []int64
}

// NewTmallfanscashpoolcheckpayRequest 初始化TmallfanscashpoolcheckpayAPIRequest对象
func NewTmallfanscashpoolcheckpayRequest() *TmallfanscashpoolcheckpayAPIRequest {
	return &TmallfanscashpoolcheckpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallfanscashpoolcheckpayAPIRequest) GetApiMethodName() string {
	return "tmall.fans.cashpool.checkpay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallfanscashpoolcheckpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallfanscashpoolcheckpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCashPoolList is CashPoolList Setter
// 资金池列表
func (r *TmallfanscashpoolcheckpayAPIRequest) SetCashPoolList(_cashPoolList []int64) error {
	r._cashPoolList = _cashPoolList
	r.Set("cash_pool_list", _cashPoolList)
	return nil
}

// GetCashPoolList CashPoolList Getter
func (r TmallfanscashpoolcheckpayAPIRequest) GetCashPoolList() []int64 {
	return r._cashPoolList
}
