package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeScanAPIResponse 查询扫码信息 API返回值
// alibaba.alihealth.drugcode.scan
//
// 查询扫码信息
type AlibabaAlihealthDrugcodeScanAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeScanAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeScanAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugcodeScanAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugcodeScanAPIResponseModel is 查询扫码信息 成功返回结果
type AlibabaAlihealthDrugcodeScanAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_scan_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeScanAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugcodeScanAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugcodeScanAPIResponse)
	},
}

// GetAlibabaAlihealthDrugcodeScanAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugcodeScanAPIResponse
func GetAlibabaAlihealthDrugcodeScanAPIResponse() *AlibabaAlihealthDrugcodeScanAPIResponse {
	return poolAlibabaAlihealthDrugcodeScanAPIResponse.Get().(*AlibabaAlihealthDrugcodeScanAPIResponse)
}

// ReleaseAlibabaAlihealthDrugcodeScanAPIResponse 将 AlibabaAlihealthDrugcodeScanAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeScanAPIResponse(v *AlibabaAlihealthDrugcodeScanAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeScanAPIResponse.Put(v)
}
