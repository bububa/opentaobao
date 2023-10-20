package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkseriesskuaddAPIResponse 系列品商品变更-添加商品 API返回值
// alibaba.wdk.series.sku.add
//
// 系列品商品变更-添加商品
type AlibabawdkseriesskuaddAPIResponse struct {
	model.CommonResponse
	AlibabawdkseriesskuaddAPIResponseModel
}

// AlibabawdkseriesskuaddAPIResponseModel is 系列品商品变更-添加商品 成功返回结果
type AlibabawdkseriesskuaddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_series_sku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	ApiResult *AlibabawdkseriesskuaddApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
