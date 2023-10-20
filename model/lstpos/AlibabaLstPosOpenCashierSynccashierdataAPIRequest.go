package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstposopencashiersynccashierdataAPIRequest 收银快照同步接口(最多10条订单信息) API请求
// alibaba.lst.pos.open.cashier.synccashierdata
//
// 收银快照同步接口(最多10条订单信息)
type AlibabalstposopencashiersynccashierdataAPIRequest struct {
	model.Params
	// 订单对象列表
	_cashierFlowDTOList []CashierFlowDto
	// 门店对应的主账号id
	_userId int64
}

// NewAlibabalstposopencashiersynccashierdataRequest 初始化AlibabalstposopencashiersynccashierdataAPIRequest对象
func NewAlibabalstposopencashiersynccashierdataRequest() *AlibabalstposopencashiersynccashierdataAPIRequest {
	return &AlibabalstposopencashiersynccashierdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstposopencashiersynccashierdataAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.pos.open.cashier.synccashierdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstposopencashiersynccashierdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstposopencashiersynccashierdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCashierFlowDTOList is CashierFlowDTOList Setter
// 订单对象列表
func (r *AlibabalstposopencashiersynccashierdataAPIRequest) SetCashierFlowDTOList(_cashierFlowDTOList []CashierFlowDto) error {
	r._cashierFlowDTOList = _cashierFlowDTOList
	r.Set("cashier_flow_d_t_o_list", _cashierFlowDTOList)
	return nil
}

// GetCashierFlowDTOList CashierFlowDTOList Getter
func (r AlibabalstposopencashiersynccashierdataAPIRequest) GetCashierFlowDTOList() []CashierFlowDto {
	return r._cashierFlowDTOList
}

// SetUserId is UserId Setter
// 门店对应的主账号id
func (r *AlibabalstposopencashiersynccashierdataAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabalstposopencashiersynccashierdataAPIRequest) GetUserId() int64 {
	return r._userId
}
