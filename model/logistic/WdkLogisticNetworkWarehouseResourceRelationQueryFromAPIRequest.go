package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest
中心仓查网格仓 API请求
wdk.logistic.network.warehouse.resource.relation.query.from

盒马集市，中心仓查询网格仓 */
type WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIRequest struct {
	model.Params
	// 查询参数
	_paramPageQueryWarehouseResourceRelationByFromRequest *PageQueryWarehouseResourceRelationByFromRequest
}

// New
