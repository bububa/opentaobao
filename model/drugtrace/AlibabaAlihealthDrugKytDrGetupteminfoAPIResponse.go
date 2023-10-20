package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrGetupteminfoAPIResponse 获取上游温度信息（疫苗） API返回值
// alibaba.alihealth.drug.kyt.dr.getupteminfo
//
// 根据追溯码及企业ID获取上游运输及存储温度信息（疫苗）
type AlibabaAlihealthDrugKytDrGetupteminfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrGetupteminfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrGetupteminfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDrGetupteminfoAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDrGetupteminfoAPIResponseModel is 获取上游温度信息（疫苗） 成功返回结果
type AlibabaAlihealthDrugKytDrGetupteminfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_getupteminfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytDrGetupteminfoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrGetupteminfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytDrGetupteminfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrGetupteminfoAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDrGetupteminfoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrGetupteminfoAPIResponse
func GetAlibabaAlihealthDrugKytDrGetupteminfoAPIResponse() *AlibabaAlihealthDrugKytDrGetupteminfoAPIResponse {
	return poolAlibabaAlihealthDrugKytDrGetupteminfoAPIResponse.Get().(*AlibabaAlihealthDrugKytDrGetupteminfoAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDrGetupteminfoAPIResponse 将 AlibabaAlihealthDrugKytDrGetupteminfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrGetupteminfoAPIResponse(v *AlibabaAlihealthDrugKytDrGetupteminfoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrGetupteminfoAPIResponse.Put(v)
}
