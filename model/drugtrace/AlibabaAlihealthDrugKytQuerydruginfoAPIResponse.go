package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytQuerydruginfoAPIResponse 码查询药品 API返回值
// alibaba.alihealth.drug.kyt.querydruginfo
//
// 通过追溯码查询药品信息
type AlibabaAlihealthDrugKytQuerydruginfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytQuerydruginfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytQuerydruginfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytQuerydruginfoAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytQuerydruginfoAPIResponseModel is 码查询药品 成功返回结果
type AlibabaAlihealthDrugKytQuerydruginfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_querydruginfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihealthDrugKytQuerydruginfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytQuerydruginfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytQuerydruginfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytQuerydruginfoAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytQuerydruginfoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytQuerydruginfoAPIResponse
func GetAlibabaAlihealthDrugKytQuerydruginfoAPIResponse() *AlibabaAlihealthDrugKytQuerydruginfoAPIResponse {
	return poolAlibabaAlihealthDrugKytQuerydruginfoAPIResponse.Get().(*AlibabaAlihealthDrugKytQuerydruginfoAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytQuerydruginfoAPIResponse 将 AlibabaAlihealthDrugKytQuerydruginfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytQuerydruginfoAPIResponse(v *AlibabaAlihealthDrugKytQuerydruginfoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytQuerydruginfoAPIResponse.Put(v)
}
