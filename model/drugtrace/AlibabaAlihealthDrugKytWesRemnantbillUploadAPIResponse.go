package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse wes零头出入库单据上传 API返回值
// alibaba.alihealth.drug.kyt.wes.remnantbill.upload
//
// wes零头出入库单据上传
type AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponseModel is wes零头出入库单据上传 成功返回结果
type AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_remnantbill_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	ResponseStatus bool `json:"response_status,omitempty" xml:"response_status,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.ResponseStatus = false
}

var poolAlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse
func GetAlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse() *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse {
	return poolAlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse.Get().(*AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse 将 AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse(v *AlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesRemnantbillUploadAPIResponse.Put(v)
}
