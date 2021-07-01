package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlicomOrderPreauthorizeCreateAPIRequest
业务办理结果 API请求
alibaba.alicom.order.preauthorize.create

授授权:签约结果通知 */
type AlibabaAlicomOrderPreauthorizeCreateAPIRequest struct {
	model.Params
	// 入参
	_preAuthorizeModel *PreAuthorizeModel
}

// New
