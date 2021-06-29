package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_体检套餐预定确认 APIResponse
alibaba.alihealth.examination.reserve.confirm

向体检机构确认用户购买的体检套餐信息
*/
type AlibabaAlihealthExaminationReserveConfirmAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationReserveConfirmResponse
}

type AlibabaAlihealthExaminationReserveConfirmResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果描述
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 体检机构预约唯一标识码
    
    UniqReserveCode   string `json:"uniq_reserve_code,omitempty" xml:"uniq_reserve_code,omitempty"`

    
    // 返回结果编码
    
    ResponseCode   string `json:"response_code,omitempty" xml:"response_code,omitempty"`

    
}
