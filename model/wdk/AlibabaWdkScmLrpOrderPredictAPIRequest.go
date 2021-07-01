package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkScmLrpOrderPredictAPIRequest
单量预测 API请求
alibaba.wdk.scm.lrp.order.predict

提供基于门店和配送站的订单量预测，可用于提前安排人力资源 */
type AlibabaWdkScmLrpOrderPredictAPIRequest struct {
	model.Params
	// 单量预测查询参数
	_paramOrderPredictQuery *OrderPredictQuery
}

// New
