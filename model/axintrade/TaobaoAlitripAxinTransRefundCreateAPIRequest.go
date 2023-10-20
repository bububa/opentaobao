package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripaxintransrefundcreateAPIRequest 阿信创建退款单 API请求
// taobao.alitrip.axin.trans.refund.create
//
// 阿信供销平台-创建退款单服务
type TaobaoalitripaxintransrefundcreateAPIRequest struct {
	model.Params
	// 阿信创建退款单入参
	_axinRefundCreateDTO *AxinRefundCreateDto
}

// NewTaobaoalitripaxintransrefundcreateRequest 初始化TaobaoalitripaxintransrefundcreateAPIRequest对象
func NewTaobaoalitripaxintransrefundcreateRequest() *TaobaoalitripaxintransrefundcreateAPIRequest {
	return &TaobaoalitripaxintransrefundcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripaxintransrefundcreateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.refund.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripaxintransrefundcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripaxintransrefundcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAxinRefundCreateDTO is AxinRefundCreateDTO Setter
// 阿信创建退款单入参
func (r *TaobaoalitripaxintransrefundcreateAPIRequest) SetAxinRefundCreateDTO(_axinRefundCreateDTO *AxinRefundCreateDto) error {
	r._axinRefundCreateDTO = _axinRefundCreateDTO
	r.Set("axin_refund_create_d_t_o", _axinRefundCreateDTO)
	return nil
}

// GetAxinRefundCreateDTO AxinRefundCreateDTO Getter
func (r TaobaoalitripaxintransrefundcreateAPIRequest) GetAxinRefundCreateDTO() *AxinRefundCreateDto {
	return r._axinRefundCreateDTO
}
