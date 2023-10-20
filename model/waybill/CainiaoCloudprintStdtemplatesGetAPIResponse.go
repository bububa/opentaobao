package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintStdtemplatesGetAPIResponse 获取所有的菜鸟标准电子面单模板 API返回值
// cainiao.cloudprint.stdtemplates.get
//
// 获取菜鸟标准电子面单模板
type CainiaoCloudprintStdtemplatesGetAPIResponse struct {
	model.CommonResponse
	CainiaoCloudprintStdtemplatesGetAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoCloudprintStdtemplatesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCloudprintStdtemplatesGetAPIResponseModel).Reset()
}

// CainiaoCloudprintStdtemplatesGetAPIResponseModel is 获取所有的菜鸟标准电子面单模板 成功返回结果
type CainiaoCloudprintStdtemplatesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cloudprint_stdtemplates_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoCloudprintStdtemplatesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoCloudprintStdtemplatesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCloudprintStdtemplatesGetAPIResponse)
	},
}

// GetCainiaoCloudprintStdtemplatesGetAPIResponse 从 sync.Pool 获取 CainiaoCloudprintStdtemplatesGetAPIResponse
func GetCainiaoCloudprintStdtemplatesGetAPIResponse() *CainiaoCloudprintStdtemplatesGetAPIResponse {
	return poolCainiaoCloudprintStdtemplatesGetAPIResponse.Get().(*CainiaoCloudprintStdtemplatesGetAPIResponse)
}

// ReleaseCainiaoCloudprintStdtemplatesGetAPIResponse 将 CainiaoCloudprintStdtemplatesGetAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCloudprintStdtemplatesGetAPIResponse(v *CainiaoCloudprintStdtemplatesGetAPIResponse) {
	v.Reset()
	poolCainiaoCloudprintStdtemplatesGetAPIResponse.Put(v)
}
