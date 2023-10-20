package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosgoodsadjustAPIRequest 调整库存 API请求
// alibaba.mos.goods.adjust
//
// 库存调整接口
type AlibabamosgoodsadjustAPIRequest struct {
	model.Params
	// 库存调整入参
	_paramIsvStockAdjustRequestDTO *IsvStockAdjustRequestDto
}

// NewAlibabamosgoodsadjustRequest 初始化AlibabamosgoodsadjustAPIRequest对象
func NewAlibabamosgoodsadjustRequest() *AlibabamosgoodsadjustAPIRequest {
	return &AlibabamosgoodsadjustAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosgoodsadjustAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.adjust"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosgoodsadjustAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosgoodsadjustAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamIsvStockAdjustRequestDTO is ParamIsvStockAdjustRequestDTO Setter
// 库存调整入参
func (r *AlibabamosgoodsadjustAPIRequest) SetParamIsvStockAdjustRequestDTO(_paramIsvStockAdjustRequestDTO *IsvStockAdjustRequestDto) error {
	r._paramIsvStockAdjustRequestDTO = _paramIsvStockAdjustRequestDTO
	r.Set("param_isv_stock_adjust_request_d_t_o", _paramIsvStockAdjustRequestDTO)
	return nil
}

// GetParamIsvStockAdjustRequestDTO ParamIsvStockAdjustRequestDTO Getter
func (r AlibabamosgoodsadjustAPIRequest) GetParamIsvStockAdjustRequestDTO() *IsvStockAdjustRequestDto {
	return r._paramIsvStockAdjustRequestDTO
}
