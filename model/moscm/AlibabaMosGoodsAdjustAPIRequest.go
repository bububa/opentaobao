package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosGoodsAdjustAPIRequest
调整库存 API请求
alibaba.mos.goods.adjust

库存调整接口 */
type AlibabaMosGoodsAdjustAPIRequest struct {
	model.Params
	// 库存调整入参
	_paramIsvStockAdjustRequestDTO *IsvStockAdjustRequestDto
}

// New
