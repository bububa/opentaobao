package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkaxstorecreateAPIResponse 翱象经营店创建接口 API返回值
// alibaba.wdk.ax.store.create
//
// 翱象经营店创建
type AlibabawdkaxstorecreateAPIResponse struct {
	model.CommonResponse
	AlibabawdkaxstorecreateAPIResponseModel
}

// AlibabawdkaxstorecreateAPIResponseModel is 翱象经营店创建接口 成功返回结果
type AlibabawdkaxstorecreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ax_store_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabawdkaxstorecreateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
