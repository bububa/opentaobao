package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse 接收盲底文件删除日志 API返回值
// alibaba.alihealth.drugcode.drugfactory.blindfiledellog
//
// 临床用药试验-接收盲底文件删除日志
type AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponseModel is 接收盲底文件删除日志 成功返回结果
type AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_blindfiledellog_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 调用结果信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 删除日志同步结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = false
}

var poolAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse)
	},
}

// GetAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse
func GetAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse() *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse {
	return poolAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse.Get().(*AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse)
}

// ReleaseAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse 将 AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse(v *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIResponse.Put(v)
}
