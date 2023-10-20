package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripaxintransfundquerybyorderAPIRequest 通过外部订单ID查询所有资金单 API请求
// taobao.alitrip.axin.trans.fund.query.by.order
//
// 阿信供销平台-通过外部订单ID查询所有资金单
type TaobaoalitripaxintransfundquerybyorderAPIRequest struct {
	model.Params
	// 入参
	_paramAxinFundListQueryDTO *AxinFundListQueryDto
}

// NewTaobaoalitripaxintransfundquerybyorderRequest 初始化TaobaoalitripaxintransfundquerybyorderAPIRequest对象
func NewTaobaoalitripaxintransfundquerybyorderRequest() *TaobaoalitripaxintransfundquerybyorderAPIRequest {
	return &TaobaoalitripaxintransfundquerybyorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripaxintransfundquerybyorderAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.fund.query.by.order"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripaxintransfundquerybyorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripaxintransfundquerybyorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAxinFundListQueryDTO is ParamAxinFundListQueryDTO Setter
// 入参
func (r *TaobaoalitripaxintransfundquerybyorderAPIRequest) SetParamAxinFundListQueryDTO(_paramAxinFundListQueryDTO *AxinFundListQueryDto) error {
	r._paramAxinFundListQueryDTO = _paramAxinFundListQueryDTO
	r.Set("param_axin_fund_list_query_d_t_o", _paramAxinFundListQueryDTO)
	return nil
}

// GetParamAxinFundListQueryDTO ParamAxinFundListQueryDTO Getter
func (r TaobaoalitripaxintransfundquerybyorderAPIRequest) GetParamAxinFundListQueryDTO() *AxinFundListQueryDto {
	return r._paramAxinFundListQueryDTO
}
