package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDestbillListAPIResponse 直调单据查询 API返回值
// alibaba.alihealth.drug.kyt.destbill.list
//
// 为药企提供直调单据查询功能
type AlibabaAlihealthDrugKytDestbillListAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDestbillListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDestbillListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDestbillListAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDestbillListAPIResponseModel is 直调单据查询 成功返回结果
type AlibabaAlihealthDrugKytDestbillListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_destbill_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回result
	Result *AlibabaAlihealthDrugKytDestbillListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDestbillListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytDestbillListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDestbillListAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDestbillListAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDestbillListAPIResponse
func GetAlibabaAlihealthDrugKytDestbillListAPIResponse() *AlibabaAlihealthDrugKytDestbillListAPIResponse {
	return poolAlibabaAlihealthDrugKytDestbillListAPIResponse.Get().(*AlibabaAlihealthDrugKytDestbillListAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDestbillListAPIResponse 将 AlibabaAlihealthDrugKytDestbillListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDestbillListAPIResponse(v *AlibabaAlihealthDrugKytDestbillListAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDestbillListAPIResponse.Put(v)
}
