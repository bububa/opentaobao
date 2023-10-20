package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrAssociateequiAPIResponse 疫苗单据与设备绑定 API返回值
// alibaba.alihealth.drug.kyt.dr.associateequi
//
// 疫苗单据与设备绑定
type AlibabaAlihealthDrugKytDrAssociateequiAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrAssociateequiAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrAssociateequiAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDrAssociateequiAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDrAssociateequiAPIResponseModel is 疫苗单据与设备绑定 成功返回结果
type AlibabaAlihealthDrugKytDrAssociateequiAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_associateequi_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytDrAssociateequiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrAssociateequiAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytDrAssociateequiAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrAssociateequiAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDrAssociateequiAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrAssociateequiAPIResponse
func GetAlibabaAlihealthDrugKytDrAssociateequiAPIResponse() *AlibabaAlihealthDrugKytDrAssociateequiAPIResponse {
	return poolAlibabaAlihealthDrugKytDrAssociateequiAPIResponse.Get().(*AlibabaAlihealthDrugKytDrAssociateequiAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDrAssociateequiAPIResponse 将 AlibabaAlihealthDrugKytDrAssociateequiAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrAssociateequiAPIResponse(v *AlibabaAlihealthDrugKytDrAssociateequiAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrAssociateequiAPIResponse.Put(v)
}
