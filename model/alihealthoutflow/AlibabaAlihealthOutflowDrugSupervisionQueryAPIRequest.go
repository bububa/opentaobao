package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest
监管平台药品查询 API请求
alibaba.alihealth.outflow.drug.supervision.query

获取监管平台药品数据 */
type AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest struct {
	model.Params
	// 请求
	_request1 *OuterDrugVo
}

// New
