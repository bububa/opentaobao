package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillCloudprintNetprintVerifycodeAPIResponse 打印验证码 API返回值
// cainiao.waybill.cloudprint.netprint.verifycode
//
// 打印获取验证码
type CainiaoWaybillCloudprintNetprintVerifycodeAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillCloudprintNetprintVerifycodeAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoWaybillCloudprintNetprintVerifycodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoWaybillCloudprintNetprintVerifycodeAPIResponseModel).Reset()
}

// CainiaoWaybillCloudprintNetprintVerifycodeAPIResponseModel is 打印验证码 成功返回结果
type CainiaoWaybillCloudprintNetprintVerifycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_cloudprint_netprint_verifycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *IsvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoWaybillCloudprintNetprintVerifycodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoWaybillCloudprintNetprintVerifycodeAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoWaybillCloudprintNetprintVerifycodeAPIResponse)
	},
}

// GetCainiaoWaybillCloudprintNetprintVerifycodeAPIResponse 从 sync.Pool 获取 CainiaoWaybillCloudprintNetprintVerifycodeAPIResponse
func GetCainiaoWaybillCloudprintNetprintVerifycodeAPIResponse() *CainiaoWaybillCloudprintNetprintVerifycodeAPIResponse {
	return poolCainiaoWaybillCloudprintNetprintVerifycodeAPIResponse.Get().(*CainiaoWaybillCloudprintNetprintVerifycodeAPIResponse)
}

// ReleaseCainiaoWaybillCloudprintNetprintVerifycodeAPIResponse 将 CainiaoWaybillCloudprintNetprintVerifycodeAPIResponse 保存到 sync.Pool
func ReleaseCainiaoWaybillCloudprintNetprintVerifycodeAPIResponse(v *CainiaoWaybillCloudprintNetprintVerifycodeAPIResponse) {
	v.Reset()
	poolCainiaoWaybillCloudprintNetprintVerifycodeAPIResponse.Put(v)
}
