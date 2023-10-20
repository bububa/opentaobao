package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthNrCepOutorderUploadAPIResponse 线上订单收货验收单、出入库单据生成接口 API返回值
// alibaba.health.nr.cep.outorder.upload
//
// 线上订单收货验收单、出入库单据生成接口
type AlibabaHealthNrCepOutorderUploadAPIResponse struct {
	model.CommonResponse
	AlibabaHealthNrCepOutorderUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHealthNrCepOutorderUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHealthNrCepOutorderUploadAPIResponseModel).Reset()
}

// AlibabaHealthNrCepOutorderUploadAPIResponseModel is 线上订单收货验收单、出入库单据生成接口 成功返回结果
type AlibabaHealthNrCepOutorderUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_nr_cep_outorder_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Result *ResponseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHealthNrCepOutorderUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHealthNrCepOutorderUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHealthNrCepOutorderUploadAPIResponse)
	},
}

// GetAlibabaHealthNrCepOutorderUploadAPIResponse 从 sync.Pool 获取 AlibabaHealthNrCepOutorderUploadAPIResponse
func GetAlibabaHealthNrCepOutorderUploadAPIResponse() *AlibabaHealthNrCepOutorderUploadAPIResponse {
	return poolAlibabaHealthNrCepOutorderUploadAPIResponse.Get().(*AlibabaHealthNrCepOutorderUploadAPIResponse)
}

// ReleaseAlibabaHealthNrCepOutorderUploadAPIResponse 将 AlibabaHealthNrCepOutorderUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHealthNrCepOutorderUploadAPIResponse(v *AlibabaHealthNrCepOutorderUploadAPIResponse) {
	v.Reset()
	poolAlibabaHealthNrCepOutorderUploadAPIResponse.Put(v)
}
