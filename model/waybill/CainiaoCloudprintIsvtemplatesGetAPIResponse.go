package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintIsvtemplatesGetAPIResponse 获取商家使用的标准模板 API返回值
// cainiao.cloudprint.isvtemplates.get
//
// 获取商家使用的标准模板
type CainiaoCloudprintIsvtemplatesGetAPIResponse struct {
	model.CommonResponse
	CainiaoCloudprintIsvtemplatesGetAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoCloudprintIsvtemplatesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCloudprintIsvtemplatesGetAPIResponseModel).Reset()
}

// CainiaoCloudprintIsvtemplatesGetAPIResponseModel is 获取商家使用的标准模板 成功返回结果
type CainiaoCloudprintIsvtemplatesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cloudprint_isvtemplates_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoCloudprintIsvtemplatesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoCloudprintIsvtemplatesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCloudprintIsvtemplatesGetAPIResponse)
	},
}

// GetCainiaoCloudprintIsvtemplatesGetAPIResponse 从 sync.Pool 获取 CainiaoCloudprintIsvtemplatesGetAPIResponse
func GetCainiaoCloudprintIsvtemplatesGetAPIResponse() *CainiaoCloudprintIsvtemplatesGetAPIResponse {
	return poolCainiaoCloudprintIsvtemplatesGetAPIResponse.Get().(*CainiaoCloudprintIsvtemplatesGetAPIResponse)
}

// ReleaseCainiaoCloudprintIsvtemplatesGetAPIResponse 将 CainiaoCloudprintIsvtemplatesGetAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCloudprintIsvtemplatesGetAPIResponse(v *CainiaoCloudprintIsvtemplatesGetAPIResponse) {
	v.Reset()
	poolCainiaoCloudprintIsvtemplatesGetAPIResponse.Put(v)
}
