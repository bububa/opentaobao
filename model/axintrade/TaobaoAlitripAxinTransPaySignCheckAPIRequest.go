package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripaxintranspaysigncheckAPIRequest 阿信支付宝验签服务 API请求
// taobao.alitrip.axin.trans.pay.sign.check
//
// 阿信支付宝验签服务
type TaobaoalitripaxintranspaysigncheckAPIRequest struct {
	model.Params
	// 验签对象
	_axinPayCheckSignDto *AxinPayCheckSignDto
}

// NewTaobaoalitripaxintranspaysigncheckRequest 初始化TaobaoalitripaxintranspaysigncheckAPIRequest对象
func NewTaobaoalitripaxintranspaysigncheckRequest() *TaobaoalitripaxintranspaysigncheckAPIRequest {
	return &TaobaoalitripaxintranspaysigncheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripaxintranspaysigncheckAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.pay.sign.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripaxintranspaysigncheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripaxintranspaysigncheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAxinPayCheckSignDto is AxinPayCheckSignDto Setter
// 验签对象
func (r *TaobaoalitripaxintranspaysigncheckAPIRequest) SetAxinPayCheckSignDto(_axinPayCheckSignDto *AxinPayCheckSignDto) error {
	r._axinPayCheckSignDto = _axinPayCheckSignDto
	r.Set("axin_pay_check_sign_dto", _axinPayCheckSignDto)
	return nil
}

// GetAxinPayCheckSignDto AxinPayCheckSignDto Getter
func (r TaobaoalitripaxintranspaysigncheckAPIRequest) GetAxinPayCheckSignDto() *AxinPayCheckSignDto {
	return r._axinPayCheckSignDto
}
