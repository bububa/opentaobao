package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取品牌下设备列表 API请求
yunos.tvpubadmin.device.models

获取品牌下设备列表
*/
type YunosTvpubadminDeviceModelsRequest struct {
    model.Params
    // 终端类型
    _terminalType   string
    // 品牌ID
    _brandId   int64
    // 牌照方
    _license   int64
}

// 初始化YunosTvpubadminDeviceModelsRequest对象
func NewYunosTvpubadminDeviceModelsRequest() *YunosTvpubadminDeviceModelsRequest{
    return &YunosTvpubadminDeviceModelsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceModelsRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.models"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceModelsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TerminalType Setter
// 终端类型
func (r *YunosTvpubadminDeviceModelsRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r YunosTvpubadminDeviceModelsRequest) GetTerminalType() string {
    return r._terminalType
}
// BrandId Setter
// 品牌ID
func (r *YunosTvpubadminDeviceModelsRequest) SetBrandId(_brandId int64) error {
    r._brandId = _brandId
    r.Set("brand_id", _brandId)
    return nil
}

// BrandId Getter
func (r YunosTvpubadminDeviceModelsRequest) GetBrandId() int64 {
    return r._brandId
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDeviceModelsRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceModelsRequest) GetLicense() int64 {
    return r._license
}
