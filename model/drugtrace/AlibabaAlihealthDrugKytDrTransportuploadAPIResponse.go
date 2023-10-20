package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrTransportuploadAPIResponse 疫苗运输温度上传 API返回值
// alibaba.alihealth.drug.kyt.dr.transportupload
//
// 疫苗运输温度上传
type AlibabaAlihealthDrugKytDrTransportuploadAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrTransportuploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrTransportuploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDrTransportuploadAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDrTransportuploadAPIResponseModel is 疫苗运输温度上传 成功返回结果
type AlibabaAlihealthDrugKytDrTransportuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_transportupload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytDrTransportuploadResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrTransportuploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytDrTransportuploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrTransportuploadAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDrTransportuploadAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrTransportuploadAPIResponse
func GetAlibabaAlihealthDrugKytDrTransportuploadAPIResponse() *AlibabaAlihealthDrugKytDrTransportuploadAPIResponse {
	return poolAlibabaAlihealthDrugKytDrTransportuploadAPIResponse.Get().(*AlibabaAlihealthDrugKytDrTransportuploadAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDrTransportuploadAPIResponse 将 AlibabaAlihealthDrugKytDrTransportuploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrTransportuploadAPIResponse(v *AlibabaAlihealthDrugKytDrTransportuploadAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrTransportuploadAPIResponse.Put(v)
}
