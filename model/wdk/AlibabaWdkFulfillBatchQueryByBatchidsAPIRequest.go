package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest
作业小票查询接口 API请求
alibaba.wdk.fulfill.batch.query.by.batchids

根据节点等条件查询履约单小票信息 */
type AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest struct {
	model.Params
	// 作业节点类型： WAREHOUSE：仓  DELIVERY_DOCK：配送站 SHOP：经营店
	_nodeType string
	// warehouseCode, 出库仓，由基础店仓维护，盒马全域统一,
	_nodeCode string
	// 批次号列表
	_batchIds []string
}

// New
