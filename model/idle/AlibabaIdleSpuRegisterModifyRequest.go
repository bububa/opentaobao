package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商spu挂载接口 APIRequest
alibaba.idle.spu.register.modify

闲鱼服务商通过此接口进行spu挂载，指明自己支持对该spu的服务(回收、验货等)
*/
type AlibabaIdleSpuRegisterModifyRequest struct {
    model.Params

    // 入参
    idleSpuRegister4TopDto   *IdleSpuRegister4TopDto 

}

func NewAlibabaIdleSpuRegisterModifyRequest() *AlibabaIdleSpuRegisterModifyRequest{
    return &AlibabaIdleSpuRegisterModifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleSpuRegisterModifyRequest) GetApiMethodName() string {
    return "alibaba.idle.spu.register.modify"
}

func (r AlibabaIdleSpuRegisterModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleSpuRegisterModifyRequest) SetIdleSpuRegister4TopDto(idleSpuRegister4TopDto *IdleSpuRegister4TopDto) error {
    r.idleSpuRegister4TopDto = idleSpuRegister4TopDto
    r.Set("idle_spu_register4_top_dto", idleSpuRegister4TopDto)
    return nil
}

func (r AlibabaIdleSpuRegisterModifyRequest) GetIdleSpuRegister4TopDto() *IdleSpuRegister4TopDto {
    return r.idleSpuRegister4TopDto
}

