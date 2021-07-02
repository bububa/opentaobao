package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransPayRegisterCreateAPIRequest 提交支付服务开通 API请求
// taobao.alitrip.axin.trans.pay.register.create
//
// 阿信供销平台-提交支付服务开通接口
type TaobaoAlitripAxinTransPayRegisterCreateAPIRequest struct {
	model.Params
	// 提交支付服务开通接口入参
	_createDTO *AxinPayRegisterCreateDto
}

// NewTaobaoAlitripAxinTransPayRegisterCreateRequest 初始化TaobaoAlitripAxinTransPayRegisterCreateAPIRequest对象
func NewTaobaoAlitripAxinTransPayRegisterCreateRequest() *TaobaoAlitripAxinTransPayRegisterCreateAPIRequest {
	return &TaobaoAlitripAxinTransPayRegisterCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransPayRegisterCreateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.pay.register.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransPayRegisterCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CreateDTO Setter
// 提交支付服务开通接口入参
func (r *TaobaoAlitripAxinTransPayRegisterCreateAPIRequest) SetCreateDTO(_createDTO *AxinPayRegisterCreateDto) error {
	r._createDTO = _createDTO
	r.Set("create_d_t_o", _createDTO)
	return nil
}

// Get CreateDTO Getter
func (r TaobaoAlitripAxinTransPayRegisterCreateAPIRequest) GetCreateDTO() *AxinPayRegisterCreateDto {
	return r._createDTO
}
