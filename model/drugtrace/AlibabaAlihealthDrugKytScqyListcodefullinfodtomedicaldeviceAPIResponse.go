package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse 医疗器械的码查询信息接口 API返回值
// alibaba.alihealth.drug.kyt.scqy.listcodefullinfodtomedicaldevice
//
// 医疗器械的码查询信息接口
type AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponseModel is 医疗器械的码查询信息接口 成功返回结果
type AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_listcodefullinfodtomedicaldevice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse
func GetAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse() *AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse {
	return poolAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse.Get().(*AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse 将 AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse(v *AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse.Put(v)
}
