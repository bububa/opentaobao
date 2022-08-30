package ascp

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangInventoryFutureplanAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.inventory.futureplan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangInventoryFutureplanAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
