package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosGoodsAdjustAPIRequest 调整库存 API请求
// alibaba.mos.goods.adjust
//
// 库存调整接口
type AlibabaMosGoodsAdjustAPIRequest struct {
	model.Params
	// 库存调整入参
	_paramIsvStockAdjustRequestDTO *IsvStockAdjustRequestDto
}

// NewAlibabaMosGoodsAdjustRequest 初始化AlibabaMosGoodsAdjustAPIRequest对象
func NewAlibabaMosGoodsAdjustRequest() *AlibabaMosGoodsAdjustAPIRequest {
	return &AlibabaMosGoodsAdjustAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosGoodsAdjustAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.adjust"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosGoodsAdjustAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamIsvStockAdjustRequestDTO Setter
// 库存调整入参
func (r *AlibabaMosGoodsAdjustAPIRequest) SetParamIsvStockAdjustRequestDTO(_paramIsvStockAdjustRequestDTO *IsvStockAdjustRequestDto) error {
	r._paramIsvStockAdjustRequestDTO = _paramIsvStockAdjustRequestDTO
	r.Set("param_isv_stock_adjust_request_d_t_o", _paramIsvStockAdjustRequestDTO)
	return nil
}

// Get ParamIsvStockAdjustRequestDTO Getter
func (r AlibabaMosGoodsAdjustAPIRequest) GetParamIsvStockAdjustRequestDTO() *IsvStockAdjustRequestDto {
	return r._paramIsvStockAdjustRequestDTO
}
