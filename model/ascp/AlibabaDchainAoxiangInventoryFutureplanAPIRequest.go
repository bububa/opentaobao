package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxianginventoryfutureplanAPIRequest 负卖计划 API请求
// alibaba.dchain.aoxiang.inventory.futureplan
//
// 负卖计划。底层有白名单控制，并非对所有商家开放。如果需要使用，请联系对应的小二增加白名单
type AlibabadchainaoxianginventoryfutureplanAPIRequest struct {
	model.Params
	// 负卖计划入参
	_publicFuturePlanRequest *PublicFuturePlanRequest
}

// NewAlibabadchainaoxianginventoryfutureplanRequest 初始化AlibabadchainaoxianginventoryfutureplanAPIRequest对象
func NewAlibabadchainaoxianginventoryfutureplanRequest() *AlibabadchainaoxianginventoryfutureplanAPIRequest {
	return &AlibabadchainaoxianginventoryfutureplanAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxianginventoryfutureplanAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.inventory.futureplan"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxianginventoryfutureplanAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxianginventoryfutureplanAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPublicFuturePlanRequest is PublicFuturePlanRequest Setter
// 负卖计划入参
func (r *AlibabadchainaoxianginventoryfutureplanAPIRequest) SetPublicFuturePlanRequest(_publicFuturePlanRequest *PublicFuturePlanRequest) error {
	r._publicFuturePlanRequest = _publicFuturePlanRequest
	r.Set("public_future_plan_request", _publicFuturePlanRequest)
	return nil
}

// GetPublicFuturePlanRequest PublicFuturePlanRequest Getter
func (r AlibabadchainaoxianginventoryfutureplanAPIRequest) GetPublicFuturePlanRequest() *PublicFuturePlanRequest {
	return r._publicFuturePlanRequest
}
