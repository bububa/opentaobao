package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailCommissionOrderSyncAPIRequest
分佣数据传输 API请求
alibaba.retail.commission.order.sync

同步分佣结果 */
type AlibabaRetailCommissionOrderSyncAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *UniverseOrderVo
}

// New
