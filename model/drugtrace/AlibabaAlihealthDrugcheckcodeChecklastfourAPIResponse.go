package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse 校验追溯码的后4位是否正确 API返回值
// alibaba.alihealth.drugcheckcode.checklastfour
//
// 校验追溯码的后4位是否正确
type AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponseModel is 校验追溯码的后4位是否正确 成功返回结果
type AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcheckcode_checklastfour_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用成功
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 调用失败
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 对象
	SuccessI bool `json:"success_i,omitempty" xml:"success_i,omitempty"`
	// 是否正确
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.SuccessI = false
	m.Model = false
}

var poolAlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse)
	},
}

// GetAlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse
func GetAlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse() *AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse {
	return poolAlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse.Get().(*AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse)
}

// ReleaseAlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse 将 AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse(v *AlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugcheckcodeChecklastfourAPIResponse.Put(v)
}
