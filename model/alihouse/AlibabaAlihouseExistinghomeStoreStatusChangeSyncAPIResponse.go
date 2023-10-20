package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomestorestatuschangesyncAPIResponse 门店状态变更 API返回值
// alibaba.alihouse.existinghome.store.status.change.sync
//
// 门店状态变更
type AlibabaalihouseexistinghomestorestatuschangesyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomestorestatuschangesyncAPIResponseModel
}

// AlibabaalihouseexistinghomestorestatuschangesyncAPIResponseModel is 门店状态变更 成功返回结果
type AlibabaalihouseexistinghomestorestatuschangesyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_store_status_change_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// aaa
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// aaa
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// aaa
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// aaa
	ReturnSuccess bool `json:"return_success,omitempty" xml:"return_success,omitempty"`
}
