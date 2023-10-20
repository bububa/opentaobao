package pur

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCeresSupplierPoQuerydetailAPIResponse 采购供应商订单明细查询接口 API返回值
// alibaba.ceres.supplier.po.querydetail
//
// 采购供应商订单明细查询接口
type AlibabaCeresSupplierPoQuerydetailAPIResponse struct {
	model.CommonResponse
	AlibabaCeresSupplierPoQuerydetailAPIResponseModel
}

// AlibabaCeresSupplierPoQuerydetailAPIResponseModel is 采购供应商订单明细查询接口 成功返回结果
type AlibabaCeresSupplierPoQuerydetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ceres_supplier_po_querydetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回消息体
	Result *SupplierPoDetailDto `json:"result,omitempty" xml:"result,omitempty"`
}
