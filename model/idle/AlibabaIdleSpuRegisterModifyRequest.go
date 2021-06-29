package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商spu挂载接口 API请求
alibaba.idle.spu.register.modify

闲鱼服务商通过此接口进行spu挂载，指明自己支持对该spu的服务(回收、验货等)
*/
type AlibabaIdleSpuRegisterModifyRequest struct {
    model.Params
    // 入参
    _idleSpuRegister4TopDto   *IdleSpuRegister4TopDTO
}

// 初始化AlibabaIdleSpuRegisterModifyRequest对象
func NewAlibabaIdleSpuRegisterModifyRequest() *AlibabaIdleSpuRegisterModifyRequest{
    return &AlibabaIdleSpuRegisterModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleSpuRegisterModifyRequest) GetApiMethodName() string {
    return "alibaba.idle.spu.register.modify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleSpuRegisterModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IdleSpuRegister4TopDto Setter
// 入参
func (r *AlibabaIdleSpuRegisterModifyRequest) SetIdleSpuRegister4TopDto(_idleSpuRegister4TopDto *IdleSpuRegister4TopDTO) error {
    r._idleSpuRegister4TopDto = _idleSpuRegister4TopDto
    r.Set("idle_spu_register4_top_dto", _idleSpuRegister4TopDto)
    return nil
}

// IdleSpuRegister4TopDto Getter
func (r AlibabaIdleSpuRegisterModifyRequest) GetIdleSpuRegister4TopDto() *IdleSpuRegister4TopDTO {
    return r._idleSpuRegister4TopDto
}
