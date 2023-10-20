package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkAxStoreUpdateAPIResponse 翱翔经营店更新接口 API返回值
// alibaba.wdk.ax.store.update
//
// 翱翔经营店更新接口
type AlibabaWdkAxStoreUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkAxStoreUpdateAPIResponseModel
}

// AlibabaWdkAxStoreUpdateAPIResponseModel is 翱翔经营店更新接口 成功返回结果
type AlibabaWdkAxStoreUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ax_store_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用接口返回结果
	ApiResult *AlibabaWdkAxStoreUpdateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
