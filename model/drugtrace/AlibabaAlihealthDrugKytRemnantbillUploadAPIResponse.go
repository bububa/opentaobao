package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytRemnantbillUploadAPIResponse 零头出入库单据上传 API返回值
// alibaba.alihealth.drug.kyt.remnantbill.upload
//
// 零头出入库单据上传
type AlibabaAlihealthDrugKytRemnantbillUploadAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytRemnantbillUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytRemnantbillUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytRemnantbillUploadAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytRemnantbillUploadAPIResponseModel is 零头出入库单据上传 成功返回结果
type AlibabaAlihealthDrugKytRemnantbillUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_remnantbill_upload_response"`
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
func (m *AlibabaAlihealthDrugKytRemnantbillUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.ResponseStatus = false
}

var poolAlibabaAlihealthDrugKytRemnantbillUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytRemnantbillUploadAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytRemnantbillUploadAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytRemnantbillUploadAPIResponse
func GetAlibabaAlihealthDrugKytRemnantbillUploadAPIResponse() *AlibabaAlihealthDrugKytRemnantbillUploadAPIResponse {
	return poolAlibabaAlihealthDrugKytRemnantbillUploadAPIResponse.Get().(*AlibabaAlihealthDrugKytRemnantbillUploadAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytRemnantbillUploadAPIResponse 将 AlibabaAlihealthDrugKytRemnantbillUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytRemnantbillUploadAPIResponse(v *AlibabaAlihealthDrugKytRemnantbillUploadAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytRemnantbillUploadAPIResponse.Put(v)
}
