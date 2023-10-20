package pur

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabapursupplierporespcreateAPIResponse po反馈创建 API返回值
// alibaba.pur.supplier.porespcreate
//
// PO反馈接口
type AlibabapursupplierporespcreateAPIResponse struct {
	model.CommonResponse
	AlibabapursupplierporespcreateAPIResponseModel
}

// AlibabapursupplierporespcreateAPIResponseModel is po反馈创建 成功返回结果
type AlibabapursupplierporespcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_supplier_porespcreate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取url的出参
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}
