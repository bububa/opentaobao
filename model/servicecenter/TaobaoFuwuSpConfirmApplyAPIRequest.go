package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFuwuSpConfirmApplyAPIRequest 内购服务确认单申请接口 API请求
// taobao.fuwu.sp.confirm.apply
//
// isv能通过该接口发起确认申请单
type TaobaoFuwuSpConfirmApplyAPIRequest struct {
	model.Params
	// 确认单申请
	_paramIncomeConfirmDTO *IncomeConfirmDto
}

// NewTaobaoFuwuSpConfirmApplyRequest 初始化TaobaoFuwuSpConfirmApplyAPIRequest对象
func NewTaobaoFuwuSpConfirmApplyRequest() *TaobaoFuwuSpConfirmApplyAPIRequest {
	return &TaobaoFuwuSpConfirmApplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFuwuSpConfirmApplyAPIRequest) Reset() {
	r._paramIncomeConfirmDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFuwuSpConfirmApplyAPIRequest) GetApiMethodName() string {
	return "taobao.fuwu.sp.confirm.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFuwuSpConfirmApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFuwuSpConfirmApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamIncomeConfirmDTO is ParamIncomeConfirmDTO Setter
// 确认单申请
func (r *TaobaoFuwuSpConfirmApplyAPIRequest) SetParamIncomeConfirmDTO(_paramIncomeConfirmDTO *IncomeConfirmDto) error {
	r._paramIncomeConfirmDTO = _paramIncomeConfirmDTO
	r.Set("param_income_confirm_d_t_o", _paramIncomeConfirmDTO)
	return nil
}

// GetParamIncomeConfirmDTO ParamIncomeConfirmDTO Getter
func (r TaobaoFuwuSpConfirmApplyAPIRequest) GetParamIncomeConfirmDTO() *IncomeConfirmDto {
	return r._paramIncomeConfirmDTO
}

var poolTaobaoFuwuSpConfirmApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFuwuSpConfirmApplyRequest()
	},
}

// GetTaobaoFuwuSpConfirmApplyRequest 从 sync.Pool 获取 TaobaoFuwuSpConfirmApplyAPIRequest
func GetTaobaoFuwuSpConfirmApplyAPIRequest() *TaobaoFuwuSpConfirmApplyAPIRequest {
	return poolTaobaoFuwuSpConfirmApplyAPIRequest.Get().(*TaobaoFuwuSpConfirmApplyAPIRequest)
}

// ReleaseTaobaoFuwuSpConfirmApplyAPIRequest 将 TaobaoFuwuSpConfirmApplyAPIRequest 放入 sync.Pool
func ReleaseTaobaoFuwuSpConfirmApplyAPIRequest(v *TaobaoFuwuSpConfirmApplyAPIRequest) {
	v.Reset()
	poolTaobaoFuwuSpConfirmApplyAPIRequest.Put(v)
}
