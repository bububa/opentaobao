package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDruploadretailAPIResponse 快易通多融零售上传接口 API返回值
// alibaba.alihealth.drug.kyt.druploadretail
//
// 快易通多融零售上传接口
type AlibabaAlihealthDrugKytDruploadretailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDruploadretailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDruploadretailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDruploadretailAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDruploadretailAPIResponseModel is 快易通多融零售上传接口 成功返回结果
type AlibabaAlihealthDrugKytDruploadretailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_druploadretail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 上传单据文件队列表标识
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 错误码(BILL_DECODE_ERROR 单据转码失败 2.BILL_FILE_NAME_DUPLICATE_UPLOAD 文件名重复)
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 操作是否成功(true 成功 ,false失败)
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDruploadretailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugKytDruploadretailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDruploadretailAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDruploadretailAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDruploadretailAPIResponse
func GetAlibabaAlihealthDrugKytDruploadretailAPIResponse() *AlibabaAlihealthDrugKytDruploadretailAPIResponse {
	return poolAlibabaAlihealthDrugKytDruploadretailAPIResponse.Get().(*AlibabaAlihealthDrugKytDruploadretailAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDruploadretailAPIResponse 将 AlibabaAlihealthDrugKytDruploadretailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDruploadretailAPIResponse(v *AlibabaAlihealthDrugKytDruploadretailAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDruploadretailAPIResponse.Put(v)
}
