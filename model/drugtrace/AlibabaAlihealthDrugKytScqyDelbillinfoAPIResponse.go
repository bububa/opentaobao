package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse 根据单据号删除单据 API返回值
// alibaba.alihealth.drug.kyt.scqy.delbillinfo
//
// 根据单据号删除单据
type AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponseModel is 根据单据号删除单据 成功返回结果
type AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_delbillinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugKytScqyDelbillinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse
func GetAlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse() *AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse {
	return poolAlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse.Get().(*AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse 将 AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse(v *AlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytScqyDelbillinfoAPIResponse.Put(v)
}
