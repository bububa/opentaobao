package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkfulfillconfigreadlimitorderAPIResponse 根据仓code查询仓限单配置 API返回值
// alibaba.wdk.fulfill.config.read.limit.order
//
// 根据仓code查询仓限单配置
type AlibabawdkfulfillconfigreadlimitorderAPIResponse struct {
	model.CommonResponse
	AlibabawdkfulfillconfigreadlimitorderAPIResponseModel
}

// AlibabawdkfulfillconfigreadlimitorderAPIResponseModel is 根据仓code查询仓限单配置 成功返回结果
type AlibabawdkfulfillconfigreadlimitorderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_fulfill_config_read_limit_order_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
}
