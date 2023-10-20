package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse 根据企业id获取往来单位 API返回值
// alibaba.alihealth.drugcode.listentparbyrefentid
//
// 根据企业id获取往来单位
type AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponseModel is 根据企业id获取往来单位 成功返回结果
type AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_listentparbyrefentid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回状态值
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回状态描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回结果
	Model *PageInfoDto `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = nil
}

var poolAlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse)
	},
}

// GetAlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse
func GetAlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse() *AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse {
	return poolAlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse.Get().(*AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse)
}

// ReleaseAlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse 将 AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse(v *AlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeListentparbyrefentidAPIResponse.Put(v)
}
