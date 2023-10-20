package pur

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaceressupplierpoqueryAPIResponse 采购供应商订单查询接口 API返回值
// alibaba.ceres.supplier.po.query
//
// 采购供应商订单查询接口
type AlibabaceressupplierpoqueryAPIResponse struct {
	model.CommonResponse
	AlibabaceressupplierpoqueryAPIResponseModel
}

// AlibabaceressupplierpoqueryAPIResponseModel is 采购供应商订单查询接口 成功返回结果
type AlibabaceressupplierpoqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ceres_supplier_po_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回消息体
	Result *AlibabaceressupplierpoqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
