package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrVaequipmentListAPIResponse 获取企业冷链设备信息 API返回值
// alibaba.alihealth.drug.kyt.dr.vaequipment.list
//
// 获取企业冷链设备信息
type AlibabaAlihealthDrugKytDrVaequipmentListAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrVaequipmentListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrVaequipmentListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDrVaequipmentListAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDrVaequipmentListAPIResponseModel is 获取企业冷链设备信息 成功返回结果
type AlibabaAlihealthDrugKytDrVaequipmentListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_vaequipment_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytDrVaequipmentListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrVaequipmentListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytDrVaequipmentListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrVaequipmentListAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDrVaequipmentListAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrVaequipmentListAPIResponse
func GetAlibabaAlihealthDrugKytDrVaequipmentListAPIResponse() *AlibabaAlihealthDrugKytDrVaequipmentListAPIResponse {
	return poolAlibabaAlihealthDrugKytDrVaequipmentListAPIResponse.Get().(*AlibabaAlihealthDrugKytDrVaequipmentListAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDrVaequipmentListAPIResponse 将 AlibabaAlihealthDrugKytDrVaequipmentListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrVaequipmentListAPIResponse(v *AlibabaAlihealthDrugKytDrVaequipmentListAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrVaequipmentListAPIResponse.Put(v)
}
