package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPurSupplierPorespcreateAPIRequest
po反馈创建 API请求
alibaba.pur.supplier.porespcreate

PO反馈接口 */
type AlibabaPurSupplierPorespcreateAPIRequest struct {
	model.Params
	// PO反馈信息
	_poResponse []SupplierPoResponseDo
}

// New
