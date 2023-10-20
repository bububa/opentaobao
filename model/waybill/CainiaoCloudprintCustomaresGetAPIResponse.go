package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintCustomaresGetAPIResponse 获取商家的自定义区模板信息 API返回值
// cainiao.cloudprint.customares.get
//
// 供isv使用，获取商家的自定义区的模板信息
type CainiaoCloudprintCustomaresGetAPIResponse struct {
	model.CommonResponse
	CainiaoCloudprintCustomaresGetAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoCloudprintCustomaresGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoCloudprintCustomaresGetAPIResponseModel).Reset()
}

// CainiaoCloudprintCustomaresGetAPIResponseModel is 获取商家的自定义区模板信息 成功返回结果
type CainiaoCloudprintCustomaresGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cloudprint_customares_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoCloudprintCustomaresGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoCloudprintCustomaresGetAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoCloudprintCustomaresGetAPIResponse)
	},
}

// GetCainiaoCloudprintCustomaresGetAPIResponse 从 sync.Pool 获取 CainiaoCloudprintCustomaresGetAPIResponse
func GetCainiaoCloudprintCustomaresGetAPIResponse() *CainiaoCloudprintCustomaresGetAPIResponse {
	return poolCainiaoCloudprintCustomaresGetAPIResponse.Get().(*CainiaoCloudprintCustomaresGetAPIResponse)
}

// ReleaseCainiaoCloudprintCustomaresGetAPIResponse 将 CainiaoCloudprintCustomaresGetAPIResponse 保存到 sync.Pool
func ReleaseCainiaoCloudprintCustomaresGetAPIResponse(v *CainiaoCloudprintCustomaresGetAPIResponse) {
	v.Reset()
	poolCainiaoCloudprintCustomaresGetAPIResponse.Put(v)
}
