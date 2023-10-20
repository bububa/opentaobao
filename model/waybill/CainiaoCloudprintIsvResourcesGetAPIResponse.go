package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintIsvResourcesGetAPIResponse isv资源查询 API返回值
// cainiao.cloudprint.isv.resources.get
//
// isv资源查询，包括isv模板、打印项、预设的自定义区等
type CainiaoCloudprintIsvResourcesGetAPIResponse struct {
	model.CommonResponse
	CainiaoCloudprintIsvResourcesGetAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoCloudprintIsvResourcesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCloudprintIsvResourcesGetAPIResponseModel).Reset()
}

// CainiaoCloudprintIsvResourcesGetAPIResponseModel is isv资源查询 成功返回结果
type CainiaoCloudprintIsvResourcesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cloudprint_isv_resources_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoCloudprintIsvResourcesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoCloudprintIsvResourcesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCloudprintIsvResourcesGetAPIResponse)
	},
}

// GetCainiaoCloudprintIsvResourcesGetAPIResponse 从 sync.Pool 获取 CainiaoCloudprintIsvResourcesGetAPIResponse
func GetCainiaoCloudprintIsvResourcesGetAPIResponse() *CainiaoCloudprintIsvResourcesGetAPIResponse {
	return poolCainiaoCloudprintIsvResourcesGetAPIResponse.Get().(*CainiaoCloudprintIsvResourcesGetAPIResponse)
}

// ReleaseCainiaoCloudprintIsvResourcesGetAPIResponse 将 CainiaoCloudprintIsvResourcesGetAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCloudprintIsvResourcesGetAPIResponse(v *CainiaoCloudprintIsvResourcesGetAPIResponse) {
	v.Reset()
	poolCainiaoCloudprintIsvResourcesGetAPIResponse.Put(v)
}
