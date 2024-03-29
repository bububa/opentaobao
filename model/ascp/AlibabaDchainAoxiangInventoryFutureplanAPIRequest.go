package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangInventoryFutureplanAPIRequest 负卖计划 API请求
// alibaba.dchain.aoxiang.inventory.futureplan
//
// 负卖计划。底层有白名单控制，并非对所有商家开放。如果需要使用，请联系对应的小二增加白名单
type AlibabaDchainAoxiangInventoryFutureplanAPIRequest struct {
	model.Params
	// 负卖计划入参
	_publicFuturePlanRequest *PublicFuturePlanRequest
}

// NewAlibabaDchainAoxiangInventoryFutureplanRequest 初始化AlibabaDchainAoxiangInventoryFutureplanAPIRequest对象
func NewAlibabaDchainAoxiangInventoryFutureplanRequest() *AlibabaDchainAoxiangInventoryFutureplanAPIRequest {
	return &AlibabaDchainAoxiangInventoryFutureplanAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangInventoryFutureplanAPIRequest) Reset() {
	r._publicFuturePlanRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangInventoryFutureplanAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.inventory.futureplan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangInventoryFutureplanAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangInventoryFutureplanAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPublicFuturePlanRequest is PublicFuturePlanRequest Setter
// 负卖计划入参
func (r *AlibabaDchainAoxiangInventoryFutureplanAPIRequest) SetPublicFuturePlanRequest(_publicFuturePlanRequest *PublicFuturePlanRequest) error {
	r._publicFuturePlanRequest = _publicFuturePlanRequest
	r.Set("public_future_plan_request", _publicFuturePlanRequest)
	return nil
}

// GetPublicFuturePlanRequest PublicFuturePlanRequest Getter
func (r AlibabaDchainAoxiangInventoryFutureplanAPIRequest) GetPublicFuturePlanRequest() *PublicFuturePlanRequest {
	return r._publicFuturePlanRequest
}

var poolAlibabaDchainAoxiangInventoryFutureplanAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangInventoryFutureplanRequest()
	},
}

// GetAlibabaDchainAoxiangInventoryFutureplanRequest 从 sync.Pool 获取 AlibabaDchainAoxiangInventoryFutureplanAPIRequest
func GetAlibabaDchainAoxiangInventoryFutureplanAPIRequest() *AlibabaDchainAoxiangInventoryFutureplanAPIRequest {
	return poolAlibabaDchainAoxiangInventoryFutureplanAPIRequest.Get().(*AlibabaDchainAoxiangInventoryFutureplanAPIRequest)
}

// ReleaseAlibabaDchainAoxiangInventoryFutureplanAPIRequest 将 AlibabaDchainAoxiangInventoryFutureplanAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangInventoryFutureplanAPIRequest(v *AlibabaDchainAoxiangInventoryFutureplanAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangInventoryFutureplanAPIRequest.Put(v)
}
