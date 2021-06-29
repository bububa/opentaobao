package nlife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
b2c转码 APIResponse
alibaba.nlife.b2c.code.convert

将商品的URL转码，ISV将该码写入RFID
*/
type AlibabaNlifeB2cCodeConvertAPIResponse struct {
    model.CommonResponse
    AlibabaNlifeB2cCodeConvertResponse
}

type AlibabaNlifeB2cCodeConvertResponse struct {
    XMLName xml.Name `xml:"alibaba_nlife_b2c_code_convert_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // data
    
    Data   *ItemCodeConvertResponse `json:"data,omitempty" xml:"data,omitempty"`

    
}
