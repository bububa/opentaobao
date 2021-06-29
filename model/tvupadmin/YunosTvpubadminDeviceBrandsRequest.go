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
type YunosTvpubadminDeviceBrandsRequest struct {
    model.Params
    // 终端类型
    terminalType   string
    // 牌照方
    license   int64
}

// 初始化YunosTvpubadminDeviceBrandsRequest对象
func NewYunosTvpubadminDeviceBrandsRequest() *YunosTvpubadminDeviceBrandsRequest{
    return &YunosTvpubadminDeviceBrandsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceBrandsRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.brands"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceBrandsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TerminalType Setter
// 终端类型
func (r *YunosTvpubadminDeviceBrandsRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

// TerminalType Getter
func (r YunosTvpubadminDeviceBrandsRequest) GetTerminalType() string {
    return r.terminalType
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDeviceBrandsRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceBrandsRequest) GetLicense() int64 {
    return r.license
}
