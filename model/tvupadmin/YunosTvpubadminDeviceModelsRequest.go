package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取品牌下设备列表 APIRequest
yunos.tvpubadmin.device.models

获取品牌下设备列表
*/
type YunosTvpubadminDeviceModelsRequest struct {
    model.Params

    // 终端类型
    terminalType   string 

    // 品牌ID
    brandId   int64 

    // 牌照方
    license   int64 

}

func NewYunosTvpubadminDeviceModelsRequest() *YunosTvpubadminDeviceModelsRequest{
    return &YunosTvpubadminDeviceModelsRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceModelsRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.models"
}

func (r YunosTvpubadminDeviceModelsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceModelsRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r YunosTvpubadminDeviceModelsRequest) GetTerminalType() string {
    return r.terminalType
}

func (r *YunosTvpubadminDeviceModelsRequest) SetBrandId(brandId int64) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

func (r YunosTvpubadminDeviceModelsRequest) GetBrandId() int64 {
    return r.brandId
}

func (r *YunosTvpubadminDeviceModelsRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminDeviceModelsRequest) GetLicense() int64 {
    return r.license
}

