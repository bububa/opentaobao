package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripaxintransfundaddAPIRequest 创建资金单接口 API请求
// taobao.alitrip.axin.trans.fund.add
//
// 创建资金单
type TaobaoalitripaxintransfundaddAPIRequest struct {
	model.Params
	// 创建资金单接口入参
	_axinFundCreateDTO *AxinFundCreateDto
}

// NewTaobaoalitripaxintransfundaddRequest 初始化TaobaoalitripaxintransfundaddAPIRequest对象
func NewTaobaoalitripaxintransfundaddRequest() *TaobaoalitripaxintransfundaddAPIRequest {
	return &TaobaoalitripaxintransfundaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripaxintransfundaddAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.fund.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripaxintransfundaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripaxintransfundaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAxinFundCreateDTO is AxinFundCreateDTO Setter
// 创建资金单接口入参
func (r *TaobaoalitripaxintransfundaddAPIRequest) SetAxinFundCreateDTO(_axinFundCreateDTO *AxinFundCreateDto) error {
	r._axinFundCreateDTO = _axinFundCreateDTO
	r.Set("axin_fund_create_d_t_o", _axinFundCreateDTO)
	return nil
}

// GetAxinFundCreateDTO AxinFundCreateDTO Getter
func (r TaobaoalitripaxintransfundaddAPIRequest) GetAxinFundCreateDTO() *AxinFundCreateDto {
	return r._axinFundCreateDTO
}
