package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkseriesdefaultskuresetAPIResponse 系列品商品变更-重置默认商品 API返回值
// alibaba.wdk.series.defaultsku.reset
//
// 系列品商品变更-重置默认商品
type AlibabawdkseriesdefaultskuresetAPIResponse struct {
	model.CommonResponse
	AlibabawdkseriesdefaultskuresetAPIResponseModel
}

// AlibabawdkseriesdefaultskuresetAPIResponseModel is 系列品商品变更-重置默认商品 成功返回结果
type AlibabawdkseriesdefaultskuresetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_series_defaultsku_reset_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	ApiResult *AlibabawdkseriesdefaultskuresetApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
