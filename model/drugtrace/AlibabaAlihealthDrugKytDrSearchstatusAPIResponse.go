package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrSearchstatusAPIResponse 疫苗企业上传单据后处理状态查询 API返回值
// alibaba.alihealth.drug.kyt.dr.searchstatus
//
// 单据处理状态查询
type AlibabaAlihealthDrugKytDrSearchstatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrSearchstatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrSearchstatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDrSearchstatusAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDrSearchstatusAPIResponseModel is 疫苗企业上传单据后处理状态查询 成功返回结果
type AlibabaAlihealthDrugKytDrSearchstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_searchstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaAlihealthDrugKytDrSearchstatusResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrSearchstatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytDrSearchstatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrSearchstatusAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDrSearchstatusAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrSearchstatusAPIResponse
func GetAlibabaAlihealthDrugKytDrSearchstatusAPIResponse() *AlibabaAlihealthDrugKytDrSearchstatusAPIResponse {
	return poolAlibabaAlihealthDrugKytDrSearchstatusAPIResponse.Get().(*AlibabaAlihealthDrugKytDrSearchstatusAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDrSearchstatusAPIResponse 将 AlibabaAlihealthDrugKytDrSearchstatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrSearchstatusAPIResponse(v *AlibabaAlihealthDrugKytDrSearchstatusAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrSearchstatusAPIResponse.Put(v)
}
