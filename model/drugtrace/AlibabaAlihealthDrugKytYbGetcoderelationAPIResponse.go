package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytYbGetcoderelationAPIResponse 医保-查询码的所有子码 API返回值
// alibaba.alihealth.drug.kyt.yb.getcoderelation
//
// 应用于药店或医院入库环节，通过扫码获取下级码进行入库；
// 通过码查询所有子码以及包装比例
type AlibabaAlihealthDrugKytYbGetcoderelationAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytYbGetcoderelationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytYbGetcoderelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytYbGetcoderelationAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytYbGetcoderelationAPIResponseModel is 医保-查询码的所有子码 成功返回结果
type AlibabaAlihealthDrugKytYbGetcoderelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_yb_getcoderelation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytYbGetcoderelationResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytYbGetcoderelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytYbGetcoderelationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytYbGetcoderelationAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytYbGetcoderelationAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytYbGetcoderelationAPIResponse
func GetAlibabaAlihealthDrugKytYbGetcoderelationAPIResponse() *AlibabaAlihealthDrugKytYbGetcoderelationAPIResponse {
	return poolAlibabaAlihealthDrugKytYbGetcoderelationAPIResponse.Get().(*AlibabaAlihealthDrugKytYbGetcoderelationAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytYbGetcoderelationAPIResponse 将 AlibabaAlihealthDrugKytYbGetcoderelationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytYbGetcoderelationAPIResponse(v *AlibabaAlihealthDrugKytYbGetcoderelationAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytYbGetcoderelationAPIResponse.Put(v)
}
