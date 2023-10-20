package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintMystdtemplatesGetAPIResponse 获取用户使用的菜鸟电子面单模板信息 API返回值
// cainiao.cloudprint.mystdtemplates.get
//
// 获取用户使用的菜鸟电子面单
type CainiaoCloudprintMystdtemplatesGetAPIResponse struct {
	model.CommonResponse
	CainiaoCloudprintMystdtemplatesGetAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoCloudprintMystdtemplatesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCloudprintMystdtemplatesGetAPIResponseModel).Reset()
}

// CainiaoCloudprintMystdtemplatesGetAPIResponseModel is 获取用户使用的菜鸟电子面单模板信息 成功返回结果
type CainiaoCloudprintMystdtemplatesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cloudprint_mystdtemplates_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoCloudprintMystdtemplatesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoCloudprintMystdtemplatesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCloudprintMystdtemplatesGetAPIResponse)
	},
}

// GetCainiaoCloudprintMystdtemplatesGetAPIResponse 从 sync.Pool 获取 CainiaoCloudprintMystdtemplatesGetAPIResponse
func GetCainiaoCloudprintMystdtemplatesGetAPIResponse() *CainiaoCloudprintMystdtemplatesGetAPIResponse {
	return poolCainiaoCloudprintMystdtemplatesGetAPIResponse.Get().(*CainiaoCloudprintMystdtemplatesGetAPIResponse)
}

// ReleaseCainiaoCloudprintMystdtemplatesGetAPIResponse 将 CainiaoCloudprintMystdtemplatesGetAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCloudprintMystdtemplatesGetAPIResponse(v *CainiaoCloudprintMystdtemplatesGetAPIResponse) {
	v.Reset()
	poolCainiaoCloudprintMystdtemplatesGetAPIResponse.Put(v)
}
