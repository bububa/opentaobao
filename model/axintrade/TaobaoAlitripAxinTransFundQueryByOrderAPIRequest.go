package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripAxinTransFundQueryByOrderAPIRequest
通过外部订单ID查询所有资金单 API请求
taobao.alitrip.axin.trans.fund.query.by.order

阿信供销平台-通过外部订单ID查询所有资金单 */
type TaobaoAlitripAxinTransFundQueryByOrderAPIRequest struct {
	model.Params
	// 入参
	_paramAxinFundListQueryDTO *AxinFundListQueryDto
}

// NewTaobaoAlitripAxinTransFundQueryByOrderRequest 初始化TaobaoAlitripAxinTransFundQueryByOrderAPIRequest对象
func NewTaobaoAlitripAxinTransFundQueryByOrderRequest() *TaobaoAlitripAxinTransFundQueryByOrderAPIRequest {
	return &TaobaoAlitripAxinTransFundQueryByOrderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransFundQueryByOrderAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.fund.query.by.order"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransFundQueryByOrderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamAxinFundListQueryDTO Setter
// 入参
func (r *TaobaoAlitripAxinTransFundQueryByOrderAPIRequest) SetParamAxinFundListQueryDTO(_paramAxinFundListQueryDTO *AxinFundListQueryDto) error {
	r._paramAxinFundListQueryDTO = _paramAxinFundListQueryDTO
	r.Set("param_axin_fund_list_query_d_t_o", _paramAxinFundListQueryDTO)
	return nil
}

// Get ParamAxinFundListQueryDTO Getter
func (r TaobaoAlitripAxinTransFundQueryByOrderAPIRequest) GetParamAxinFundListQueryDTO() *AxinFundListQueryDto {
	return r._paramAxinFundListQueryDTO
}
