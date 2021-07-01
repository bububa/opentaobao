package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取终端类型下品牌列表 API请求
yunos.tvpubadmin.device.brands

获取终端类型下品牌列表
*/
type YunosTvpubadminDeviceBrandsAPIRequest struct {
    model.Params
    // 终端类型
    _terminalType   string
    // 牌照方
    _license   int64
}

// 初始化YunosTvpubadminDeviceBrandsAPIRequest对象
func NewYunosTvpubadminDeviceBrandsRequest() *YunosTvpubadminDeviceBrandsAPIRequest{
    return &YunosTvpubadminDeviceBrandsAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceBrandsAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.brands"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceBrandsAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TerminalType Setter
// 终端类型
func (r *YunosTvpubadminDeviceBrandsAPIRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r YunosTvpubadminDeviceBrandsAPIRequest) GetTerminalType() string {
    return r._terminalType
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDeviceBrandsAPIRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceBrandsAPIRequest) GetLicense() int64 {
    return r._license
}
