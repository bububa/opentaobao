package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytUpstorebillfileAPIResponse 上传零售出入库单(上传文件) API返回值
// alibaba.alihealth.drug.kyt.upstorebillfile
//
// 上传零售出入库单(上传文件)
type AlibabaAlihealthDrugKytUpstorebillfileAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytUpstorebillfileAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytUpstorebillfileAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytUpstorebillfileAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytUpstorebillfileAPIResponseModel is 上传零售出入库单(上传文件) 成功返回结果
type AlibabaAlihealthDrugKytUpstorebillfileAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_upstorebillfile_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回ID
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 请求编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 请求描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回接口
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytUpstorebillfileAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugKytUpstorebillfileAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytUpstorebillfileAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytUpstorebillfileAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytUpstorebillfileAPIResponse
func GetAlibabaAlihealthDrugKytUpstorebillfileAPIResponse() *AlibabaAlihealthDrugKytUpstorebillfileAPIResponse {
	return poolAlibabaAlihealthDrugKytUpstorebillfileAPIResponse.Get().(*AlibabaAlihealthDrugKytUpstorebillfileAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytUpstorebillfileAPIResponse 将 AlibabaAlihealthDrugKytUpstorebillfileAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytUpstorebillfileAPIResponse(v *AlibabaAlihealthDrugKytUpstorebillfileAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytUpstorebillfileAPIResponse.Put(v)
}
