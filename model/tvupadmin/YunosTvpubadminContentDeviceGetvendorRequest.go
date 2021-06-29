package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备Vendor信息 APIRequest
yunos.tvpubadmin.content.device.getvendor

查询设备Vendor信息
*/
type YunosTvpubadminContentDeviceGetvendorRequest struct {
    model.Params

    // license
    license   int64 

    // brand_id
    brandId   int64 

}

func NewYunosTvpubadminContentDeviceGetvendorRequest() *YunosTvpubadminContentDeviceGetvendorRequest{
    return &YunosTvpubadminContentDeviceGetvendorRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentDeviceGetvendorRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.device.getvendor"
}

func (r YunosTvpubadminContentDeviceGetvendorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentDeviceGetvendorRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminContentDeviceGetvendorRequest) GetLicense() int64 {
    return r.license
}

func (r *YunosTvpubadminContentDeviceGetvendorRequest) SetBrandId(brandId int64) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

func (r YunosTvpubadminContentDeviceGetvendorRequest) GetBrandId() int64 {
    return r.brandId
}

