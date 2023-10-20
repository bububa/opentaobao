package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse 传输盲底文件 API返回值
// alibaba.alihealth.drugcode.drugfactory.transferblind
//
// 临床用药试验-传输盲底文件
type AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponseModel is 传输盲底文件 成功返回结果
type AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_transferblind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 操作说明
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 调用结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = false
}

var poolAlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse)
	},
}

// GetAlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse
func GetAlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse() *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse {
	return poolAlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse.Get().(*AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse)
}

// ReleaseAlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse 将 AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse(v *AlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeDrugfactoryTransferblindAPIResponse.Put(v)
}
