package ship

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripShipProductSynccallAPIRequest
全量同步回调 API请求
alitrip.ship.product.synccall

全量同步接口 */
type AlitripShipProductSynccallAPIRequest struct {
	model.Params
}

// NewAlitripShipProductSynccallRequest 初始化AlitripShipProductSynccallAPIRequest对象
func NewAlitripShipProductSynccallRequest() *AlitripShipProductSynccallAPIRequest {
	return &AlitripShipProductSynccallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripShipProductSynccallAPIRequest) GetApiMethodName() string {
	return "alitrip.ship.product.synccall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripShipProductSynccallAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
