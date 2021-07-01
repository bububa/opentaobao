package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlicomOrderPreauthorizeFulfillmentAPIRequest
履约结果 API请求
alibaba.alicom.order.preauthorize.fulfillment

预授权-履约结果 */
type AlibabaAlicomOrderPreauthorizeFulfillmentAPIRequest struct {
	model.Params
	// 入参
	_preAuthorizeModel *PreAuthorizeModel
}

// New
