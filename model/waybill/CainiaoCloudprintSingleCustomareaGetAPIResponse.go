package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintSingleCustomareaGetAPIResponse 获取商家单一自定义区 API返回值
// cainiao.cloudprint.single.customarea.get
//
// 商家所有快递公司模板只有一个自定义区
type CainiaoCloudprintSingleCustomareaGetAPIResponse struct {
	model.CommonResponse
	CainiaoCloudprintSingleCustomareaGetAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoCloudprintSingleCustomareaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCloudprintSingleCustomareaGetAPIResponseModel).Reset()
}

// CainiaoCloudprintSingleCustomareaGetAPIResponseModel is 获取商家单一自定义区 成功返回结果
type CainiaoCloudprintSingleCustomareaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cloudprint_single_customarea_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoCloudprintSingleCustomareaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoCloudprintSingleCustomareaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCloudprintSingleCustomareaGetAPIResponse)
	},
}

// GetCainiaoCloudprintSingleCustomareaGetAPIResponse 从 sync.Pool 获取 CainiaoCloudprintSingleCustomareaGetAPIResponse
func GetCainiaoCloudprintSingleCustomareaGetAPIResponse() *CainiaoCloudprintSingleCustomareaGetAPIResponse {
	return poolCainiaoCloudprintSingleCustomareaGetAPIResponse.Get().(*CainiaoCloudprintSingleCustomareaGetAPIResponse)
}

// ReleaseCainiaoCloudprintSingleCustomareaGetAPIResponse 将 CainiaoCloudprintSingleCustomareaGetAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCloudprintSingleCustomareaGetAPIResponse(v *CainiaoCloudprintSingleCustomareaGetAPIResponse) {
	v.Reset()
	poolCainiaoCloudprintSingleCustomareaGetAPIResponse.Put(v)
}
