package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytGetcodebillinfoAPIResponse 根据码获取基本和单据信息 API返回值
// alibaba.alihealth.drug.kyt.getcodebillinfo
//
// 根据码信息获取基本信息和单据信息
type AlibabaAlihealthDrugKytGetcodebillinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytGetcodebillinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytGetcodebillinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytGetcodebillinfoAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytGetcodebillinfoAPIResponseModel is 根据码获取基本和单据信息 成功返回结果
type AlibabaAlihealthDrugKytGetcodebillinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_getcodebillinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytGetcodebillinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytGetcodebillinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytGetcodebillinfoAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytGetcodebillinfoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytGetcodebillinfoAPIResponse
func GetAlibabaAlihealthDrugKytGetcodebillinfoAPIResponse() *AlibabaAlihealthDrugKytGetcodebillinfoAPIResponse {
	return poolAlibabaAlihealthDrugKytGetcodebillinfoAPIResponse.Get().(*AlibabaAlihealthDrugKytGetcodebillinfoAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytGetcodebillinfoAPIResponse 将 AlibabaAlihealthDrugKytGetcodebillinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytGetcodebillinfoAPIResponse(v *AlibabaAlihealthDrugKytGetcodebillinfoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytGetcodebillinfoAPIResponse.Put(v)
}
