package wlbimports

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest 首公里揽收-快递预查询服务 API请求
// cainiao.global.im.pickup.bigbag.express.prequery
//
// 快递预查询服务
type CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest struct {
	model.Params
	// 快递预查询服务请求参数
	_expressPreQueryRequest *ExpressPreQueryRequest
}

// NewCainiaoGlobalImPickupBigbagExpressPrequeryRequest 初始化CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest对象
func NewCainiaoGlobalImPickupBigbagExpressPrequeryRequest() *CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest {
	return &CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest) Reset() {
	r._expressPreQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.bigbag.express.prequery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExpressPreQueryRequest is ExpressPreQueryRequest Setter
// 快递预查询服务请求参数
func (r *CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest) SetExpressPreQueryRequest(_expressPreQueryRequest *ExpressPreQueryRequest) error {
	r._expressPreQueryRequest = _expressPreQueryRequest
	r.Set("express_pre_query_request", _expressPreQueryRequest)
	return nil
}

// GetExpressPreQueryRequest ExpressPreQueryRequest Getter
func (r CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest) GetExpressPreQueryRequest() *ExpressPreQueryRequest {
	return r._expressPreQueryRequest
}

var poolCainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoGlobalImPickupBigbagExpressPrequeryRequest()
	},
}

// GetCainiaoGlobalImPickupBigbagExpressPrequeryRequest 从 sync.Pool 获取 CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest
func GetCainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest() *CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest {
	return poolCainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest.Get().(*CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest)
}

// ReleaseCainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest 将 CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest 放入 sync.Pool
func ReleaseCainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest(v *CainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest) {
	v.Reset()
	poolCainiaoGlobalImPickupBigbagExpressPrequeryAPIRequest.Put(v)
}
