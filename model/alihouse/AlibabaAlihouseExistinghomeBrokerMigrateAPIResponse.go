package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomebrokermigrateAPIResponse 融合店经纪人迁移 API返回值
// alibaba.alihouse.existinghome.broker.migrate
//
// 融合店经纪人迁移
type AlibabaalihouseexistinghomebrokermigrateAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomebrokermigrateAPIResponseModel
}

// AlibabaalihouseexistinghomebrokermigrateAPIResponseModel is 融合店经纪人迁移 成功返回结果
type AlibabaalihouseexistinghomebrokermigrateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_broker_migrate_response"`
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
