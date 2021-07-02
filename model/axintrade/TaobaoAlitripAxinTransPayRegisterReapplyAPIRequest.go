package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransPayRegisterReapplyAPIRequest 阿信支付入驻重新申请 API请求
// taobao.alitrip.axin.trans.pay.register.reapply
//
// 阿信支付入驻重新申请
// 用于支付平台驳回，商户提交时的业务场景
type TaobaoAlitripAxinTransPayRegisterReapplyAPIRequest struct {
	model.Params
	// 阿信支付入驻重新申请参数
	_axinPayRegisterCreateDTO *AxinPayRegisterCreateDto
}

// NewTaobaoAlitripAxinTransPayRegisterReapplyRequest 初始化TaobaoAlitripAxinTransPayRegisterReapplyAPIRequest对象
func NewTaobaoAlitripAxinTransPayRegisterReapplyRequest() *TaobaoAlitripAxinTransPayRegisterReapplyAPIRequest {
	return &TaobaoAlitripAxinTransPayRegisterReapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransPayRegisterReapplyAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.pay.register.reapply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransPayRegisterReapplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAxinPayRegisterCreateDTO is AxinPayRegisterCreateDTO Setter
// 阿信支付入驻重新申请参数
func (r *TaobaoAlitripAxinTransPayRegisterReapplyAPIRequest) SetAxinPayRegisterCreateDTO(_axinPayRegisterCreateDTO *AxinPayRegisterCreateDto) error {
	r._axinPayRegisterCreateDTO = _axinPayRegisterCreateDTO
	r.Set("axin_pay_register_create_d_t_o", _axinPayRegisterCreateDTO)
	return nil
}

// GetAxinPayRegisterCreateDTO AxinPayRegisterCreateDTO Getter
func (r TaobaoAlitripAxinTransPayRegisterReapplyAPIRequest) GetAxinPayRegisterCreateDTO() *AxinPayRegisterCreateDto {
	return r._axinPayRegisterCreateDTO
}
