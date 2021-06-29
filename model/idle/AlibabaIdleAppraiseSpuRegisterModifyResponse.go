package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
验货宝服务商spu挂载 APIResponse
alibaba.idle.appraise.spu.register.modify

闲鱼接收回收商spu模板挂载信息
*/
type AlibabaIdleAppraiseSpuRegisterModifyAPIResponse struct {
    model.CommonResponse
    AlibabaIdleAppraiseSpuRegisterModifyResponse
}

type AlibabaIdleAppraiseSpuRegisterModifyResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_appraise_spu_register_modify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaIdleAppraiseSpuRegisterModifyResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
