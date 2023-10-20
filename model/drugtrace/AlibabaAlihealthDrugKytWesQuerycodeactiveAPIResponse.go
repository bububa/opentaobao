package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse 查询码是否激活 API返回值
// alibaba.alihealth.drug.kyt.wes.querycodeactive
//
// 查询码是否激活
type AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponseModel is 查询码是否激活 成功返回结果
type AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_querycodeactive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 未激活或者不存在的码
	Models []string `json:"models,omitempty" xml:"models>string,omitempty"`
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 消息提示内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功(true 成功 ,false失败)
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Models = m.Models[:0]
	m.MsgCode = ""
	m.MsgInfo = ""
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse
func GetAlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse() *AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse {
	return poolAlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse.Get().(*AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse 将 AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse(v *AlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesQuerycodeactiveAPIResponse.Put(v)
}
