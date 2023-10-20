package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse wes码关联关系查询 API返回值
// alibaba.alihealth.drug.code.kyt.wes.querycoderelation
//
// 单码关联关系查询，通过一个码查询这个码下的所有子码。（只有做过入库的码，才能能进行查询）
type AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponseModel is wes码关联关系查询 成功返回结果
type AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_wes_querycoderelation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugCodeKytWesQuerycoderelationResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse)
	},
}

// GetAlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse
func GetAlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse() *AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse {
	return poolAlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse.Get().(*AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse)
}

// ReleaseAlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse 将 AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse(v *AlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytWesQuerycoderelationAPIResponse.Put(v)
}
