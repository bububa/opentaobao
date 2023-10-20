package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeNodenameGetAPIResponse 根据码获取机构名称 API返回值
// alibaba.alihealth.drugcode.nodename.get
//
// 根据码获取机构名称
type AlibabaAlihealthDrugcodeNodenameGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeNodenameGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeNodenameGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugcodeNodenameGetAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugcodeNodenameGetAPIResponseModel is 根据码获取机构名称 成功返回结果
type AlibabaAlihealthDrugcodeNodenameGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_nodename_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Model []DrugScancodeFlowLogDto `json:"model,omitempty" xml:"model>drug_scancode_flow_log_dto,omitempty"`
	// 备注
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeNodenameGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = m.Model[:0]
	m.MsgInfo = ""
	m.MsgCode = ""
}

var poolAlibabaAlihealthDrugcodeNodenameGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugcodeNodenameGetAPIResponse)
	},
}

// GetAlibabaAlihealthDrugcodeNodenameGetAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugcodeNodenameGetAPIResponse
func GetAlibabaAlihealthDrugcodeNodenameGetAPIResponse() *AlibabaAlihealthDrugcodeNodenameGetAPIResponse {
	return poolAlibabaAlihealthDrugcodeNodenameGetAPIResponse.Get().(*AlibabaAlihealthDrugcodeNodenameGetAPIResponse)
}

// ReleaseAlibabaAlihealthDrugcodeNodenameGetAPIResponse 将 AlibabaAlihealthDrugcodeNodenameGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeNodenameGetAPIResponse(v *AlibabaAlihealthDrugcodeNodenameGetAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeNodenameGetAPIResponse.Put(v)
}
