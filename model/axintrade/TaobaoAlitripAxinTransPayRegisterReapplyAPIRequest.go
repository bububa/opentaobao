package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripaxintranspayregisterreapplyAPIRequest 阿信支付入驻重新申请 API请求
// taobao.alitrip.axin.trans.pay.register.reapply
//
// 阿信支付入驻重新申请
// 用于支付平台驳回，商户提交时的业务场景
type TaobaoalitripaxintranspayregisterreapplyAPIRequest struct {
	model.Params
	// 阿信支付入驻重新申请参数
	_axinPayRegisterCreateDTO *AxinPayRegisterCreateDto
}

// NewTaobaoalitripaxintranspayregisterreapplyRequest 初始化TaobaoalitripaxintranspayregisterreapplyAPIRequest对象
func NewTaobaoalitripaxintranspayregisterreapplyRequest() *TaobaoalitripaxintranspayregisterreapplyAPIRequest {
	return &TaobaoalitripaxintranspayregisterreapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripaxintranspayregisterreapplyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.pay.register.reapply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripaxintranspayregisterreapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripaxintranspayregisterreapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAxinPayRegisterCreateDTO is AxinPayRegisterCreateDTO Setter
// 阿信支付入驻重新申请参数
func (r *TaobaoalitripaxintranspayregisterreapplyAPIRequest) SetAxinPayRegisterCreateDTO(_axinPayRegisterCreateDTO *AxinPayRegisterCreateDto) error {
	r._axinPayRegisterCreateDTO = _axinPayRegisterCreateDTO
	r.Set("axin_pay_register_create_d_t_o", _axinPayRegisterCreateDTO)
	return nil
}

// GetAxinPayRegisterCreateDTO AxinPayRegisterCreateDTO Getter
func (r TaobaoalitripaxintranspayregisterreapplyAPIRequest) GetAxinPayRegisterCreateDTO() *AxinPayRegisterCreateDto {
	return r._axinPayRegisterCreateDTO
}
