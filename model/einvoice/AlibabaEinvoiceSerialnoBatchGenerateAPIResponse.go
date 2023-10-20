package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceserialnobatchgenerateAPIResponse 开票流水号批量生成接口 API返回值
// alibaba.einvoice.serialno.batch.generate
//
// 批量获取开票流水号接口。此接口1次返回1000条开票流水号，每个应用每天限流1000次调用。
// 优先使用alibaba.einvoice.serial.generate。
type AlibabaeinvoiceserialnobatchgenerateAPIResponse struct {
	model.CommonResponse
	AlibabaeinvoiceserialnobatchgenerateAPIResponseModel
}

// AlibabaeinvoiceserialnobatchgenerateAPIResponseModel is 开票流水号批量生成接口 成功返回结果
type AlibabaeinvoiceserialnobatchgenerateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_serialno_batch_generate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	SerialNoList []string `json:"serial_no_list,omitempty" xml:"serial_no_list>string,omitempty"`
}
