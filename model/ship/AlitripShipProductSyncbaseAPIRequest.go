package ship

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripShipProductSyncbaseAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripShipProductSyncbaseAPIRequest) GetApiMethodName() string {
	return "alitrip.ship.product.syncbase"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripShipProductSyncbaseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripShipProductSyncbaseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlitripShipProductSyncbaseAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripShipProductSyncbaseRequest()
	},
}

// GetAlitripShipProductSyncbaseRequest 从 sync.Pool 获取 AlitripShipProductSyncbaseAPIRequest
func GetAlitripShipProductSyncbaseAPIRequest() *AlitripShipProductSyncbaseAPIRequest {
	return poolAlitripShipProductSyncbaseAPIRequest.Get().(*AlitripShipProductSyncbaseAPIRequest)
}

// ReleaseAlitripShipProductSyncbaseAPIRequest 将 AlitripShipProductSyncbaseAPIRequest 放入 sync.Pool
func ReleaseAlitripShipProductSyncbaseAPIRequest(v *AlitripShipProductSyncbaseAPIRequest) {
	v.Reset()
	poolAlitripShipProductSyncbaseAPIRequest.Put(v)
}
