package alihealthmedical

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalOrderQueryAPIRequest
三方机构查询订单详情接口 API请求
alibaba.alihealth.medical.order.query

查询订单详情，包括评价 */
type AlibabaAlihealthMedicalOrderQueryAPIRequest struct {
	model.Params
	// 请求入参
	_requestInfo *OrderQueryRequestDto
}

// New
