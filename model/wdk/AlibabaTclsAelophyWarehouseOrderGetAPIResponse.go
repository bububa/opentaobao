package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophywarehouseordergetAPIResponse 仓作业单获取 API返回值
// alibaba.tcls.aelophy.warehouse.order.get
//
// 仓作业单获取
type AlibabatclsaelophywarehouseordergetAPIResponse struct {
	model.CommonResponse
	AlibabatclsaelophywarehouseordergetAPIResponseModel
}

// AlibabatclsaelophywarehouseordergetAPIResponseModel is 仓作业单获取 成功返回结果
type AlibabatclsaelophywarehouseordergetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_warehouse_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *TopBaseResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
