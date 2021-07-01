package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIRequest
履约结果查询 API请求
alibaba.alicom.order.preauthorize.query.fulfillment

预授权-履约结果查询 */
type AlibabaAlicomOrderPreauthorizeQueryFulfillmentAPIRequest struct {
	model.Params
	// 入参
	_preAuthorizeModel *PreAuthorizeModel
}

// New
