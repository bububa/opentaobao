package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrDrugrecalAPIResponse 疫苗药品召回 API返回值
// alibaba.alihealth.drug.kyt.dr.drugrecal
//
// 生产企业发布的召回信息，按照批次进行召回，收货和发货环节的单据处理中调用接口进行查询；
type AlibabaAlihealthDrugKytDrDrugrecalAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrDrugrecalAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrDrugrecalAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDrDrugrecalAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDrDrugrecalAPIResponseModel is 疫苗药品召回 成功返回结果
type AlibabaAlihealthDrugKytDrDrugrecalAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_drugrecal_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytDrDrugrecalResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrDrugrecalAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytDrDrugrecalAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrDrugrecalAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDrDrugrecalAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrDrugrecalAPIResponse
func GetAlibabaAlihealthDrugKytDrDrugrecalAPIResponse() *AlibabaAlihealthDrugKytDrDrugrecalAPIResponse {
	return poolAlibabaAlihealthDrugKytDrDrugrecalAPIResponse.Get().(*AlibabaAlihealthDrugKytDrDrugrecalAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDrDrugrecalAPIResponse 将 AlibabaAlihealthDrugKytDrDrugrecalAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrDrugrecalAPIResponse(v *AlibabaAlihealthDrugKytDrDrugrecalAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrDrugrecalAPIResponse.Put(v)
}
