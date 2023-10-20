package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripaxintransfundupdateAPIRequest 修改资金单接口 API请求
// taobao.alitrip.axin.trans.fund.update
//
// 阿信供销平台-修改资金单接口
type TaobaoalitripaxintransfundupdateAPIRequest struct {
	model.Params
	// 更新资金单接口入参
	_axinFundUpdateDTO *AxinFundUpdateDto
}

// NewTaobaoalitripaxintransfundupdateRequest 初始化TaobaoalitripaxintransfundupdateAPIRequest对象
func NewTaobaoalitripaxintransfundupdateRequest() *TaobaoalitripaxintransfundupdateAPIRequest {
	return &TaobaoalitripaxintransfundupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripaxintransfundupdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.fund.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripaxintransfundupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripaxintransfundupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAxinFundUpdateDTO is AxinFundUpdateDTO Setter
// 更新资金单接口入参
func (r *TaobaoalitripaxintransfundupdateAPIRequest) SetAxinFundUpdateDTO(_axinFundUpdateDTO *AxinFundUpdateDto) error {
	r._axinFundUpdateDTO = _axinFundUpdateDTO
	r.Set("axin_fund_update_d_t_o", _axinFundUpdateDTO)
	return nil
}

// GetAxinFundUpdateDTO AxinFundUpdateDTO Getter
func (r TaobaoalitripaxintransfundupdateAPIRequest) GetAxinFundUpdateDTO() *AxinFundUpdateDto {
	return r._axinFundUpdateDTO
}
