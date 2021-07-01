package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest
按网格仓查中心仓（带缓存） API请求
wdk.logistic.network.warehouse.resource.relation.query.to.codes

盒马集市，网格仓查询中心仓 */
type WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest struct {
	model.Params
	// 入参
	_paramYxWarehouseResourceRelationQueryRequest *YxWarehouseResourceRelationQueryRequest
}

// New
