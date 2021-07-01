package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* WdkLogisticNetworkResourceGroupQueryAPIRequest
查询网格仓-区块-自提点关系 API请求
wdk.logistic.network.resource.group.query

查询网格仓-区块-自提点关系 */
type WdkLogisticNetworkResourceGroupQueryAPIRequest struct {
	model.Params
	// 入参
	_paramResourceGroupPageQueryRequest *ResourceGroupPageQueryRequest
}

// New
