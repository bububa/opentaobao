package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskubarcodequeryAPIResponse 商品条码查询接口 API返回值
// alibaba.wdk.sku.barcode.query
//
// 查询商品编码，支持一品多码
type AlibabawdkskubarcodequeryAPIResponse struct {
	model.CommonResponse
	AlibabawdkskubarcodequeryAPIResponseModel
}

// AlibabawdkskubarcodequeryAPIResponseModel is 商品条码查询接口 成功返回结果
type AlibabawdkskubarcodequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_barcode_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabawdkskubarcodequeryApiResults `json:"result,omitempty" xml:"result,omitempty"`
}
