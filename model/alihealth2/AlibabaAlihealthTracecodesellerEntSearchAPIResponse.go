package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerEntSearchAPIResponse 查询商家信息 API返回值
// alibaba.alihealth.tracecodeseller.ent.search
//
// 查询商家信息
type AlibabaAlihealthTracecodesellerEntSearchAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTracecodesellerEntSearchAPIResponseModel
}

// AlibabaAlihealthTracecodesellerEntSearchAPIResponseModel is 查询商家信息 成功返回结果
type AlibabaAlihealthTracecodesellerEntSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_ent_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
