package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAxIntegrationAccountImportAPIResponse ISV用户录入 API返回值
// alibaba.tcls.ax.integration.account.import
//
// ISV的用户录入翱象
type AlibabaTclsAxIntegrationAccountImportAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAxIntegrationAccountImportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTclsAxIntegrationAccountImportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTclsAxIntegrationAccountImportAPIResponseModel).Reset()
}

// AlibabaTclsAxIntegrationAccountImportAPIResponseModel is ISV用户录入 成功返回结果
type AlibabaTclsAxIntegrationAccountImportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_ax_integration_account_import_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	Result *AlibabaTclsAxIntegrationAccountImportApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTclsAxIntegrationAccountImportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTclsAxIntegrationAccountImportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTclsAxIntegrationAccountImportAPIResponse)
	},
}

// GetAlibabaTclsAxIntegrationAccountImportAPIResponse 从 sync.Pool 获取 AlibabaTclsAxIntegrationAccountImportAPIResponse
func GetAlibabaTclsAxIntegrationAccountImportAPIResponse() *AlibabaTclsAxIntegrationAccountImportAPIResponse {
	return poolAlibabaTclsAxIntegrationAccountImportAPIResponse.Get().(*AlibabaTclsAxIntegrationAccountImportAPIResponse)
}

// ReleaseAlibabaTclsAxIntegrationAccountImportAPIResponse 将 AlibabaTclsAxIntegrationAccountImportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTclsAxIntegrationAccountImportAPIResponse(v *AlibabaTclsAxIntegrationAccountImportAPIResponse) {
	v.Reset()
	poolAlibabaTclsAxIntegrationAccountImportAPIResponse.Put(v)
}
