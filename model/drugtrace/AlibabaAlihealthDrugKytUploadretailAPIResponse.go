package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytUploadretailAPIResponse 门店销售扫码回传接口 API返回值
// alibaba.alihealth.drug.kyt.uploadretail
//
// 门店在销售给顾客时，扫描追溯码的数据按照单据回传；
type AlibabaAlihealthDrugKytUploadretailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytUploadretailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytUploadretailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytUploadretailAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytUploadretailAPIResponseModel is 门店销售扫码回传接口 成功返回结果
type AlibabaAlihealthDrugKytUploadretailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_uploadretail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 上传单据文件队列表标识
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 错误码(BILL_DECODE_ERROR 单据转码失败 2.BILL_FILE_NAME_DUPLICATE_UPLOAD 文件名重复)
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功(true 成功 ,false失败)
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytUploadretailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugKytUploadretailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytUploadretailAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytUploadretailAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytUploadretailAPIResponse
func GetAlibabaAlihealthDrugKytUploadretailAPIResponse() *AlibabaAlihealthDrugKytUploadretailAPIResponse {
	return poolAlibabaAlihealthDrugKytUploadretailAPIResponse.Get().(*AlibabaAlihealthDrugKytUploadretailAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytUploadretailAPIResponse 将 AlibabaAlihealthDrugKytUploadretailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytUploadretailAPIResponse(v *AlibabaAlihealthDrugKytUploadretailAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytUploadretailAPIResponse.Put(v)
}
