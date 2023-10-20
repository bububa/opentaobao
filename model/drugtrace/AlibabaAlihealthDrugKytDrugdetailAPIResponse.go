package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrugdetailAPIResponse 查询药品详细信息 API返回值
// alibaba.alihealth.drug.kyt.drugdetail
//
// 查询药品详细信息
type AlibabaAlihealthDrugKytDrugdetailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrugdetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrugdetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDrugdetailAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDrugdetailAPIResponseModel is 查询药品详细信息 成功返回结果
type AlibabaAlihealthDrugKytDrugdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_drugdetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytDrugdetailResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrugdetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytDrugdetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrugdetailAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDrugdetailAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrugdetailAPIResponse
func GetAlibabaAlihealthDrugKytDrugdetailAPIResponse() *AlibabaAlihealthDrugKytDrugdetailAPIResponse {
	return poolAlibabaAlihealthDrugKytDrugdetailAPIResponse.Get().(*AlibabaAlihealthDrugKytDrugdetailAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDrugdetailAPIResponse 将 AlibabaAlihealthDrugKytDrugdetailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrugdetailAPIResponse(v *AlibabaAlihealthDrugKytDrugdetailAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrugdetailAPIResponse.Put(v)
}
