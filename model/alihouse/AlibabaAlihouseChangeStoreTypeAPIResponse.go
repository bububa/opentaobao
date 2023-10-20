package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousechangestoretypeAPIResponse 融合店迁移门店 API返回值
// alibaba.alihouse.change.store.type
//
// 融合店迁移门店
type AlibabaalihousechangestoretypeAPIResponse struct {
	model.CommonResponse
	AlibabaalihousechangestoretypeAPIResponseModel
}

// AlibabaalihousechangestoretypeAPIResponseModel is 融合店迁移门店 成功返回结果
type AlibabaalihousechangestoretypeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_change_store_type_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 1
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 1
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 1
	ReturnSuccess bool `json:"return_success,omitempty" xml:"return_success,omitempty"`
}
