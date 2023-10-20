package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse 零售单据上传接口 API返回值
// alibaba.alihealth.drugtrace.top.lsyd.uploadretail
//
// 快易通多融零售上传接口
type AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponseModel is 零售单据上传接口 成功返回结果
type AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_uploadretail_response"`
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
func (m *AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse
func GetAlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse() *AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse {
	return poolAlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse.Get().(*AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse 将 AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse(v *AlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydUploadretailAPIResponse.Put(v)
}
