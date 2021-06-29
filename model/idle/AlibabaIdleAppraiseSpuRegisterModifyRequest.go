package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
验货宝服务商spu挂载 APIRequest
alibaba.idle.appraise.spu.register.modify

闲鱼接收回收商spu模板挂载信息
*/
type AlibabaIdleAppraiseSpuRegisterModifyRequest struct {
    model.Params

    // 入参
    idleAppraiseSpuRegister4TopDto   *IdleAppraiseSpuRegister4TopDTO 

}

func NewAlibabaIdleAppraiseSpuRegisterModifyRequest() *AlibabaIdleAppraiseSpuRegisterModifyRequest{
    return &AlibabaIdleAppraiseSpuRegisterModifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleAppraiseSpuRegisterModifyRequest) GetApiMethodName() string {
    return "alibaba.idle.appraise.spu.register.modify"
}

func (r AlibabaIdleAppraiseSpuRegisterModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleAppraiseSpuRegisterModifyRequest) SetIdleAppraiseSpuRegister4TopDto(idleAppraiseSpuRegister4TopDto *IdleAppraiseSpuRegister4TopDTO) error {
    r.idleAppraiseSpuRegister4TopDto = idleAppraiseSpuRegister4TopDto
    r.Set("idle_appraise_spu_register4_top_dto", idleAppraiseSpuRegister4TopDto)
    return nil
}

func (r AlibabaIdleAppraiseSpuRegisterModifyRequest) GetIdleAppraiseSpuRegister4TopDto() *IdleAppraiseSpuRegister4TopDTO {
    return r.idleAppraiseSpuRegister4TopDto
}

