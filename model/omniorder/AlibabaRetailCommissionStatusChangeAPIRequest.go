package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailCommissionStatusChangeAPIRequest
分佣状态变更 API请求
alibaba.retail.commission.status.change

分佣系统，分佣状态变更接口 */
type AlibabaRetailCommissionStatusChangeAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *UniverseOrderVo
}

// New
