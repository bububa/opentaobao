package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取终端类型下品牌列表 APIRequest
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

func NewYunosTvpubadminDeviceBrandsRequest() *YunosTvpubadminDeviceBrandsRequest{
    return &YunosTvpubadminDeviceBrandsRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceBrandsRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.brands"
}

func (r YunosTvpubadminDeviceBrandsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceBrandsRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r YunosTvpubadminDeviceBrandsRequest) GetTerminalType() string {
    return r.terminalType
}

func (r *YunosTvpubadminDeviceBrandsRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminDeviceBrandsRequest) GetLicense() int64 {
    return r.license
}

