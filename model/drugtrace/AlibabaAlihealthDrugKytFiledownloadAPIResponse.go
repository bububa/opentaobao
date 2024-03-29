package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytFiledownloadAPIResponse 处理失败单据下载 API返回值
// alibaba.alihealth.drug.kyt.filedownload
//
// 处理失败单据下载
type AlibabaAlihealthDrugKytFiledownloadAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytFiledownloadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytFiledownloadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytFiledownloadAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytFiledownloadAPIResponseModel is 处理失败单据下载 成功返回结果
type AlibabaAlihealthDrugKytFiledownloadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_filedownload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 文件内容
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回结果编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回结果
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytFiledownloadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugKytFiledownloadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytFiledownloadAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytFiledownloadAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytFiledownloadAPIResponse
func GetAlibabaAlihealthDrugKytFiledownloadAPIResponse() *AlibabaAlihealthDrugKytFiledownloadAPIResponse {
	return poolAlibabaAlihealthDrugKytFiledownloadAPIResponse.Get().(*AlibabaAlihealthDrugKytFiledownloadAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytFiledownloadAPIResponse 将 AlibabaAlihealthDrugKytFiledownloadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytFiledownloadAPIResponse(v *AlibabaAlihealthDrugKytFiledownloadAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytFiledownloadAPIResponse.Put(v)
}
