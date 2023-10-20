package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceSerialnoBatchGenerateAPIResponse 开票流水号批量生成接口 API返回值
// alibaba.einvoice.serialno.batch.generate
//
// 批量获取开票流水号接口。此接口1次返回1000条开票流水号，每个应用每天限流1000次调用。
// 优先使用alibaba.einvoice.serial.generate。
type AlibabaEinvoiceSerialnoBatchGenerateAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceSerialnoBatchGenerateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceSerialnoBatchGenerateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceSerialnoBatchGenerateAPIResponseModel).Reset()
}

// AlibabaEinvoiceSerialnoBatchGenerateAPIResponseModel is 开票流水号批量生成接口 成功返回结果
type AlibabaEinvoiceSerialnoBatchGenerateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_serialno_batch_generate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	SerialNoList []string `json:"serial_no_list,omitempty" xml:"serial_no_list>string,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceSerialnoBatchGenerateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SerialNoList = m.SerialNoList[:0]
}

var poolAlibabaEinvoiceSerialnoBatchGenerateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceSerialnoBatchGenerateAPIResponse)
	},
}

// GetAlibabaEinvoiceSerialnoBatchGenerateAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceSerialnoBatchGenerateAPIResponse
func GetAlibabaEinvoiceSerialnoBatchGenerateAPIResponse() *AlibabaEinvoiceSerialnoBatchGenerateAPIResponse {
	return poolAlibabaEinvoiceSerialnoBatchGenerateAPIResponse.Get().(*AlibabaEinvoiceSerialnoBatchGenerateAPIResponse)
}

// ReleaseAlibabaEinvoiceSerialnoBatchGenerateAPIResponse 将 AlibabaEinvoiceSerialnoBatchGenerateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceSerialnoBatchGenerateAPIResponse(v *AlibabaEinvoiceSerialnoBatchGenerateAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceSerialnoBatchGenerateAPIResponse.Put(v)
}
