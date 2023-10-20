package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintTemplatesMigrateAPIResponse 云打印模板迁移接口 API返回值
// cainiao.cloudprint.templates.migrate
//
// 云打印模板迁移接口
type CainiaoCloudprintTemplatesMigrateAPIResponse struct {
	model.CommonResponse
	CainiaoCloudprintTemplatesMigrateAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoCloudprintTemplatesMigrateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCloudprintTemplatesMigrateAPIResponseModel).Reset()
}

// CainiaoCloudprintTemplatesMigrateAPIResponseModel is 云打印模板迁移接口 成功返回结果
type CainiaoCloudprintTemplatesMigrateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cloudprint_templates_migrate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoCloudprintTemplatesMigrateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoCloudprintTemplatesMigrateAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCloudprintTemplatesMigrateAPIResponse)
	},
}

// GetCainiaoCloudprintTemplatesMigrateAPIResponse 从 sync.Pool 获取 CainiaoCloudprintTemplatesMigrateAPIResponse
func GetCainiaoCloudprintTemplatesMigrateAPIResponse() *CainiaoCloudprintTemplatesMigrateAPIResponse {
	return poolCainiaoCloudprintTemplatesMigrateAPIResponse.Get().(*CainiaoCloudprintTemplatesMigrateAPIResponse)
}

// ReleaseCainiaoCloudprintTemplatesMigrateAPIResponse 将 CainiaoCloudprintTemplatesMigrateAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCloudprintTemplatesMigrateAPIResponse(v *CainiaoCloudprintTemplatesMigrateAPIResponse) {
	v.Reset()
	poolCainiaoCloudprintTemplatesMigrateAPIResponse.Put(v)
}
