package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘口令检查 API请求
alibaba.baichuan.taopassword.check

检查当前文本是否为淘口令
*/
type AlibabaBaichuanTaopasswordCheckRequest struct {
    model.Params
    // 参数DTO
    _paramDto   *ParamDTO
    // 系统自动生成
    _clientInfo   *RichClientInfo
}

// 初始化AlibabaBaichuanTaopasswordCheckRequest对象
func NewAlibabaBaichuanTaopasswordCheckRequest() *AlibabaBaichuanTaopasswordCheckRequest{
    return &AlibabaBaichuanTaopasswordCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanTaopasswordCheckRequest) GetApiMethodName() string {
    return "alibaba.baichuan.taopassword.check"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanTaopasswordCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamDto Setter
// 参数DTO
func (r *AlibabaBaichuanTaopasswordCheckRequest) SetParamDto(_paramDto *ParamDTO) error {
    r._paramDto = _paramDto
    r.Set("param_dto", _paramDto)
    return nil
}

// ParamDto Getter
func (r AlibabaBaichuanTaopasswordCheckRequest) GetParamDto() *ParamDTO {
    return r._paramDto
}
// ClientInfo Setter
// 系统自动生成
func (r *AlibabaBaichuanTaopasswordCheckRequest) SetClientInfo(_clientInfo *RichClientInfo) error {
    r._clientInfo = _clientInfo
    r.Set("client_info", _clientInfo)
    return nil
}

// ClientInfo Getter
func (r AlibabaBaichuanTaopasswordCheckRequest) GetClientInfo() *RichClientInfo {
    return r._clientInfo
}
