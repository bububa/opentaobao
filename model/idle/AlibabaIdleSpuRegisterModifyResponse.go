package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商spu挂载接口 APIResponse
alibaba.idle.spu.register.modify

闲鱼服务商通过此接口进行spu挂载，指明自己支持对该spu的服务(回收、验货等)
*/
type AlibabaIdleSpuRegisterModifyAPIResponse struct {
    model.CommonResponse
    AlibabaIdleSpuRegisterModifyResponse
}

type AlibabaIdleSpuRegisterModifyResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_spu_register_modify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaIdleSpuRegisterModifyResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
