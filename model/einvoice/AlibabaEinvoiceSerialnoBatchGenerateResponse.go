package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
开票流水号批量生成接口 APIResponse
alibaba.einvoice.serialno.batch.generate

批量获取开票流水号接口。此接口1次返回1000条开票流水号，每个应用每天限流1000次调用。
优先使用alibaba.einvoice.serial.generate。
*/
type AlibabaEinvoiceSerialnoBatchGenerateAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceSerialnoBatchGenerateResponse
}

type AlibabaEinvoiceSerialnoBatchGenerateResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_serialno_batch_generate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    SerialNoList   []string `json:"serial_no_list,omitempty" xml:"serial_no_list>string,omitempty"`
    
    
}
