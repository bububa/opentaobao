package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBrokerMigrateAPIResponse 融合店经纪人迁移 API返回值
// alibaba.alihouse.existinghome.broker.migrate
//
// 融合店经纪人迁移
type AlibabaAlihouseExistinghomeBrokerMigrateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeBrokerMigrateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeBrokerMigrateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeBrokerMigrateAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeBrokerMigrateAPIResponseModel is 融合店经纪人迁移 成功返回结果
type AlibabaAlihouseExistinghomeBrokerMigrateAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeBrokerMigrateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReturnCode = ""
	m.ReturnMessage = ""
	m.Data = false
	m.ReturnSuccess = false
}

var poolAlibabaAlihouseExistinghomeBrokerMigrateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeBrokerMigrateAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeBrokerMigrateAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeBrokerMigrateAPIResponse
func GetAlibabaAlihouseExistinghomeBrokerMigrateAPIResponse() *AlibabaAlihouseExistinghomeBrokerMigrateAPIResponse {
	return poolAlibabaAlihouseExistinghomeBrokerMigrateAPIResponse.Get().(*AlibabaAlihouseExistinghomeBrokerMigrateAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeBrokerMigrateAPIResponse 将 AlibabaAlihouseExistinghomeBrokerMigrateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeBrokerMigrateAPIResponse(v *AlibabaAlihouseExistinghomeBrokerMigrateAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeBrokerMigrateAPIResponse.Put(v)
}
