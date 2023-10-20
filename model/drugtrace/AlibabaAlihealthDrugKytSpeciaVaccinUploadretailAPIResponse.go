package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse 零售单据上传接口（疫苗） API返回值
// alibaba.alihealth.drug.kyt.specia.vaccin.uploadretail
//
// 零售上传单据信息接口
type AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponseModel is 零售单据上传接口（疫苗） 成功返回结果
type AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_specia_vaccin_uploadretail_response"`
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
func (m *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.ResponseSuccess = false
}

var poolAlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse
func GetAlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse() *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse {
	return poolAlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse.Get().(*AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse 将 AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse(v *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIResponse.Put(v)
}
