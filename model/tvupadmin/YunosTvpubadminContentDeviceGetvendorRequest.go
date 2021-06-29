package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备Vendor信息 API请求
yunos.tvpubadmin.content.device.getvendor

查询设备Vendor信息
*/
type YunosTvpubadminContentDeviceGetvendorRequest struct {
    model.Params
    // license
    _license   int64
    // brand_id
    _brandId   int64
}

// 初始化YunosTvpubadminContentDeviceGetvendorRequest对象
func NewYunosTvpubadminContentDeviceGetvendorRequest() *YunosTvpubadminContentDeviceGetvendorRequest{
    return &YunosTvpubadminContentDeviceGetvendorRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentDeviceGetvendorRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.device.getvendor"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentDeviceGetvendorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// License Setter
// license
func (r *YunosTvpubadminContentDeviceGetvendorRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminContentDeviceGetvendorRequest) GetLicense() int64 {
    return r._license
}
// BrandId Setter
// brand_id
func (r *YunosTvpubadminContentDeviceGetvendorRequest) SetBrandId(_brandId int64) error {
    r._brandId = _brandId
    r.Set("brand_id", _brandId)
    return nil
}

// BrandId Getter
func (r YunosTvpubadminContentDeviceGetvendorRequest) GetBrandId() int64 {
    return r._brandId
}
