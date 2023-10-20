package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintClientinfoPutAPIResponse 云打印客户端监控信息收集 API返回值
// cainiao.cloudprint.clientinfo.put
//
// 云打印客户端监控信息收集
type CainiaoCloudprintClientinfoPutAPIResponse struct {
	model.CommonResponse
	CainiaoCloudprintClientinfoPutAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoCloudprintClientinfoPutAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCloudprintClientinfoPutAPIResponseModel).Reset()
}

// CainiaoCloudprintClientinfoPutAPIResponseModel is 云打印客户端监控信息收集 成功返回结果
type CainiaoCloudprintClientinfoPutAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cloudprint_clientinfo_put_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoCloudprintClientinfoPutAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoCloudprintClientinfoPutAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCloudprintClientinfoPutAPIResponse)
	},
}

// GetCainiaoCloudprintClientinfoPutAPIResponse 从 sync.Pool 获取 CainiaoCloudprintClientinfoPutAPIResponse
func GetCainiaoCloudprintClientinfoPutAPIResponse() *CainiaoCloudprintClientinfoPutAPIResponse {
	return poolCainiaoCloudprintClientinfoPutAPIResponse.Get().(*CainiaoCloudprintClientinfoPutAPIResponse)
}

// ReleaseCainiaoCloudprintClientinfoPutAPIResponse 将 CainiaoCloudprintClientinfoPutAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCloudprintClientinfoPutAPIResponse(v *CainiaoCloudprintClientinfoPutAPIResponse) {
	v.Reset()
	poolCainiaoCloudprintClientinfoPutAPIResponse.Put(v)
}
