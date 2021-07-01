package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest
商家基础经营设置信息同步 API请求
taobao.koubei.saas.base.operation.config.sync

ISV接入口碑SAAS后, 经营设置数据同步到口碑SAAS */
type TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest struct {
	model.Params
	// 商户ID
	_merchantId string
	// 请求ID
	_requestId string
	// 业务类型。支付方式：payment_method
	_bizType string
	// 经营设置json串
	_operationConfig string
	// 操作员ID
	_outerOperatorId string
}

// New
