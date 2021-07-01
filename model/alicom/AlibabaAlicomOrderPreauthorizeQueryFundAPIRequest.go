package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest
资金流水查询 API请求
alibaba.alicom.order.preauthorize.query.fund

预授权-资金流水查询 */
type AlibabaAlicomOrderPreauthorizeQueryFundAPIRequest struct {
	model.Params
	// 入参
	_preAuthorizeModel *PreAuthorizeModel
}

// New
