package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesellerwarehousesearchAPIResponse 查询仓库api API返回值
// alibaba.alihealth.tracecodeseller.warehouse.search
//
// 查询仓库api
type AlibabaalihealthtracecodesellerwarehousesearchAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthtracecodesellerwarehousesearchAPIResponseModel
}

// AlibabaalihealthtracecodesellerwarehousesearchAPIResponseModel is 查询仓库api 成功返回结果
type AlibabaalihealthtracecodesellerwarehousesearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_warehouse_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
