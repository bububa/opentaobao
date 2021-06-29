package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商spu挂载接口 API返回值 
alibaba.idle.spu.register.modify

闲鱼服务商通过此接口进行spu挂载，指明自己支持对该spu的服务(回收、验货等)
*/
type AlibabaIdleSpuRegisterModifyAPIResponse struct {
    model.CommonResponse
    AlibabaIdleSpuRegisterModifyResponse
}

// 服务商spu挂载接口 成功返回结果
type AlibabaIdleSpuRegisterModifyResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_spu_register_modify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaIdleSpuRegisterModifyResult `json:"result,omitempty" xml:"result,omitempty"`
}
