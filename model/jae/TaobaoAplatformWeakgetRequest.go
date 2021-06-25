package jae

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
活动平台弱登录接口 APIRequest
taobao.aplatform.weakget

无线活动平台的开放接口，提供商品信息等的读操作
*/
type TaobaoAplatformWeakgetRequest struct {
    model.Params

    // 客户端自带参数
    paramRichClientInfo   *RichClientInfo 

    // 业务自定义参数
    paramDto   *ParamDto 

}

func NewTaobaoAplatformWeakgetRequest() *TaobaoAplatformWeakgetRequest{
    return &TaobaoAplatformWeakgetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAplatformWeakgetRequest) GetApiMethodName() string {
    return "taobao.aplatform.weakget"
}

func (r TaobaoAplatformWeakgetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAplatformWeakgetRequest) SetParamRichClientInfo(paramRichClientInfo *RichClientInfo) error {
    r.paramRichClientInfo = paramRichClientInfo
    r.Set("param_rich_client_info", paramRichClientInfo)
    return nil
}

func (r TaobaoAplatformWeakgetRequest) GetParamRichClientInfo() *RichClientInfo {
    return r.paramRichClientInfo
}

func (r *TaobaoAplatformWeakgetRequest) SetParamDto(paramDto *ParamDto) error {
    r.paramDto = paramDto
    r.Set("param_dto", paramDto)
    return nil
}

func (r TaobaoAplatformWeakgetRequest) GetParamDto() *ParamDto {
    return r.paramDto
}

