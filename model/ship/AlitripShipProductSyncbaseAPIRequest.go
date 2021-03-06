package ship

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripShipProductSyncbaseAPIRequest 基础信息修改回调 API请求
// alitrip.ship.product.syncbase
//
// 基础信息修改回调
type AlitripShipProductSyncbaseAPIRequest struct {
	model.Params
}

// NewAlitripShipProductSyncbaseRequest 初始化AlitripShipProductSyncbaseAPIRequest对象
func NewAlitripShipProductSyncbaseRequest() *AlitripShipProductSyncbaseAPIRequest {
	return &AlitripShipProductSyncbaseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripShipProductSyncbaseAPIRequest) GetApiMethodName() string {
	return "alitrip.ship.product.syncbase"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripShipProductSyncbaseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
