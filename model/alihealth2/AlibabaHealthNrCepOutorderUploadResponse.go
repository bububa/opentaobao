package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线上订单收货验收单、出入库单据生成接口 APIResponse
alibaba.health.nr.cep.outorder.upload

线上订单收货验收单、出入库单据生成接口
*/
type AlibabaHealthNrCepOutorderUploadAPIResponse struct {
    model.CommonResponse
    AlibabaHealthNrCepOutorderUploadResponse
}

type AlibabaHealthNrCepOutorderUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_health_nr_cep_outorder_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 服务出参
    
    Result   *ResponseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
