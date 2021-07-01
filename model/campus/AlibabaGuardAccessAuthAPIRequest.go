package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴权 API请求
alibaba.guard.access.auth

刷卡鉴权
*/
type AlibabaGuardAccessAuthAPIRequest struct {
    model.Params
    // 请求
    _paramIdentifyAuthDTO   *IdentifyAuthDto
}

// 初始化AlibabaGuardAccessAuthAPIRequest对象
func NewAlibabaGuardAccessAuthRequest() *AlibabaGuardAccessAuthAPIRequest{
    return &AlibabaGuardAccessAuthAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaGuardAccessAuthAPIRequest) GetApiMethodName() string {
    return "alibaba.guard.access.auth"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaGuardAccessAuthAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamIdentifyAuthDTO Setter
// 请求
func (r *AlibabaGuardAccessAuthAPIRequest) SetParamIdentifyAuthDTO(_paramIdentifyAuthDTO *IdentifyAuthDto) error {
    r._paramIdentifyAuthDTO = _paramIdentifyAuthDTO
    r.Set("param_identify_auth_d_t_o", _paramIdentifyAuthDTO)
    return nil
}

// ParamIdentifyAuthDTO Getter
func (r AlibabaGuardAccessAuthAPIRequest) GetParamIdentifyAuthDTO() *IdentifyAuthDto {
    return r._paramIdentifyAuthDTO
}
