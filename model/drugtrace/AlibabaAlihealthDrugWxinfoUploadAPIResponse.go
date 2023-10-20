package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugWxinfoUploadAPIResponse 小程序数据回传 API返回值
// alibaba.alihealth.drug.wxinfo.upload
//
// 小程序数据回传
type AlibabaAlihealthDrugWxinfoUploadAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugWxinfoUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugWxinfoUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugWxinfoUploadAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugWxinfoUploadAPIResponseModel is 小程序数据回传 成功返回结果
type AlibabaAlihealthDrugWxinfoUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_wxinfo_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugWxinfoUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Model = false
}

var poolAlibabaAlihealthDrugWxinfoUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugWxinfoUploadAPIResponse)
	},
}

// GetAlibabaAlihealthDrugWxinfoUploadAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugWxinfoUploadAPIResponse
func GetAlibabaAlihealthDrugWxinfoUploadAPIResponse() *AlibabaAlihealthDrugWxinfoUploadAPIResponse {
	return poolAlibabaAlihealthDrugWxinfoUploadAPIResponse.Get().(*AlibabaAlihealthDrugWxinfoUploadAPIResponse)
}

// ReleaseAlibabaAlihealthDrugWxinfoUploadAPIResponse 将 AlibabaAlihealthDrugWxinfoUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugWxinfoUploadAPIResponse(v *AlibabaAlihealthDrugWxinfoUploadAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugWxinfoUploadAPIResponse.Put(v)
}
