package jae

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
活动平台弱登录接口 API请求
taobao.aplatform.weakget

无线活动平台的开放接口，提供商品信息等的读操作
*/
type TaobaoAplatformWeakgetAPIRequest struct {
    model.Params
    // 客户端自带参数
    _paramRichClientInfo   *RichClientInfo
    // 业务自定义参数
    _paramDto   *ParamDTO
}

// 初始化TaobaoAplatformWeakgetAPIRequest对象
func NewTaobaoAplatformWeakgetRequest() *TaobaoAplatformWeakgetAPIRequest{
    return &TaobaoAplatformWeakgetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAplatformWeakgetAPIRequest) GetApiMethodName() string {
    return "taobao.aplatform.weakget"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAplatformWeakgetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamRichClientInfo Setter
// 客户端自带参数
func (r *TaobaoAplatformWeakgetAPIRequest) SetParamRichClientInfo(_paramRichClientInfo *RichClientInfo) error {
    r._paramRichClientInfo = _paramRichClientInfo
    r.Set("param_rich_client_info", _paramRichClientInfo)
    return nil
}

// ParamRichClientInfo Getter
func (r TaobaoAplatformWeakgetAPIRequest) GetParamRichClientInfo() *RichClientInfo {
    return r._paramRichClientInfo
}
// ParamDto Setter
// 业务自定义参数
func (r *TaobaoAplatformWeakgetAPIRequest) SetParamDto(_paramDto *ParamDTO) error {
    r._paramDto = _paramDto
    r.Set("param_dto", _paramDto)
    return nil
}

// ParamDto Getter
func (r TaobaoAplatformWeakgetAPIRequest) GetParamDto() *ParamDTO {
    return r._paramDto
}
