package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴权 APIRequest
alibaba.guard.access.auth

刷卡鉴权
*/
type AlibabaGuardAccessAuthRequest struct {
    model.Params

    // 请求
    paramIdentifyAuthDTO   *IdentifyAuthDto 

}

func NewAlibabaGuardAccessAuthRequest() *AlibabaGuardAccessAuthRequest{
    return &AlibabaGuardAccessAuthRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaGuardAccessAuthRequest) GetApiMethodName() string {
    return "alibaba.guard.access.auth"
}

func (r AlibabaGuardAccessAuthRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaGuardAccessAuthRequest) SetParamIdentifyAuthDTO(paramIdentifyAuthDTO *IdentifyAuthDto) error {
    r.paramIdentifyAuthDTO = paramIdentifyAuthDTO
    r.Set("param_identify_auth_d_t_o", paramIdentifyAuthDTO)
    return nil
}

func (r AlibabaGuardAccessAuthRequest) GetParamIdentifyAuthDTO() *IdentifyAuthDto {
    return r.paramIdentifyAuthDTO
}

