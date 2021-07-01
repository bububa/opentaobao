package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosGoodsSetpriceAPIRequest
价格变更接口 API请求
alibaba.mos.goods.setprice

价格变更接口，供供应商修改价格时使用。 */
type AlibabaMosGoodsSetpriceAPIRequest struct {
	model.Params
	// 价格变更对象列表
	_priceDtoList []PriceDto
}

// New
