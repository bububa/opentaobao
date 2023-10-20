package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse 医院药品子码申请接口 API返回值
// alibaba.alihealth.drug.code.kyt.yy.applycode
//
// 根据父码及所属企业ID生成子码信息
type AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponseModel is 医院药品子码申请接口 成功返回结果
type AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_yy_applycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *AlibabaAlihealthDrugCodeKytYyApplycodeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse)
	},
}

// GetAlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse
func GetAlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse() *AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse {
	return poolAlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse.Get().(*AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse)
}

// ReleaseAlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse 将 AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse(v *AlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytYyApplycodeAPIResponse.Put(v)
}
