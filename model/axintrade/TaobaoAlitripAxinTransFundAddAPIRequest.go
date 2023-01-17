package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransFundAddAPIRequest 创建资金单接口 API请求
// taobao.alitrip.axin.trans.fund.add
//
// 创建资金单
type TaobaoAlitripAxinTransFundAddAPIRequest struct {
	model.Params
	// 创建资金单接口入参
	_axinFundCreateDTO *AxinFundCreateDto
}

// NewTaobaoAlitripAxinTransFundAddRequest 初始化TaobaoAlitripAxinTransFundAddAPIRequest对象
func NewTaobaoAlitripAxinTransFundAddRequest() *TaobaoAlitripAxinTransFundAddAPIRequest {
	return &TaobaoAlitripAxinTransFundAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransFundAddAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.fund.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransFundAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripAxinTransFundAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAxinFundCreateDTO is AxinFundCreateDTO Setter
// 创建资金单接口入参
func (r *TaobaoAlitripAxinTransFundAddAPIRequest) SetAxinFundCreateDTO(_axinFundCreateDTO *AxinFundCreateDto) error {
	r._axinFundCreateDTO = _axinFundCreateDTO
	r.Set("axin_fund_create_d_t_o", _axinFundCreateDTO)
	return nil
}

// GetAxinFundCreateDTO AxinFundCreateDTO Getter
func (r TaobaoAlitripAxinTransFundAddAPIRequest) GetAxinFundCreateDTO() *AxinFundCreateDto {
	return r._axinFundCreateDTO
}
