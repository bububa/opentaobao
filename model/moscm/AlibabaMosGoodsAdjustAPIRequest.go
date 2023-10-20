package moscm

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosGoodsAdjustAPIRequest) Reset() {
	r._paramIsvStockAdjustRequestDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosGoodsAdjustAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.adjust"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosGoodsAdjustAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosGoodsAdjustAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamIsvStockAdjustRequestDTO is ParamIsvStockAdjustRequestDTO Setter
// 库存调整入参
func (r *AlibabaMosGoodsAdjustAPIRequest) SetParamIsvStockAdjustRequestDTO(_paramIsvStockAdjustRequestDTO *IsvStockAdjustRequestDto) error {
	r._paramIsvStockAdjustRequestDTO = _paramIsvStockAdjustRequestDTO
	r.Set("param_isv_stock_adjust_request_d_t_o", _paramIsvStockAdjustRequestDTO)
	return nil
}

// GetParamIsvStockAdjustRequestDTO ParamIsvStockAdjustRequestDTO Getter
func (r AlibabaMosGoodsAdjustAPIRequest) GetParamIsvStockAdjustRequestDTO() *IsvStockAdjustRequestDto {
	return r._paramIsvStockAdjustRequestDTO
}

var poolAlibabaMosGoodsAdjustAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosGoodsAdjustRequest()
	},
}

// GetAlibabaMosGoodsAdjustRequest 从 sync.Pool 获取 AlibabaMosGoodsAdjustAPIRequest
func GetAlibabaMosGoodsAdjustAPIRequest() *AlibabaMosGoodsAdjustAPIRequest {
	return poolAlibabaMosGoodsAdjustAPIRequest.Get().(*AlibabaMosGoodsAdjustAPIRequest)
}

// ReleaseAlibabaMosGoodsAdjustAPIRequest 将 AlibabaMosGoodsAdjustAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosGoodsAdjustAPIRequest(v *AlibabaMosGoodsAdjustAPIRequest) {
	v.Reset()
	poolAlibabaMosGoodsAdjustAPIRequest.Put(v)
}
