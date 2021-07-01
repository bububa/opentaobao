package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPurSupplierAsncreateAPIRequest
asn创建 API请求
alibaba.pur.supplier.asncreate

asn创建 */
type AlibabaPurSupplierAsncreateAPIRequest struct {
	model.Params
	// asn头信息
	_asn *SupplierAsnInfoVO
}

// New
