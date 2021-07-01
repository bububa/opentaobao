package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWttOrderContractSubscribeAPIRequest
分销商合约生产 API请求
alibaba.wtt.order.contract.subscribe

分销商合约生产 */
type AlibabaWttOrderContractSubscribeAPIRequest struct {
	model.Params
	// 分销商合约生产
	_distributionOrderModel *DistributionOrderModel
}

// New
