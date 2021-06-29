package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备列表 APIRequest
yunos.tvpubadmin.device.query

获取设备列表
*/
type YunosTvpubadminDeviceQueryRequest struct {
    model.Params

    // 终端类型
    terminalType   string 

    // 品牌ID
    brandId   int64 

    // 牌照方
    license   int64 

    // 页码值
    pageNo   int64 

    // 每页条数
    pageSize   int64 

}

func NewYunosTvpubadminDeviceQueryRequest() *YunosTvpubadminDeviceQueryRequest{
    return &YunosTvpubadminDeviceQueryRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDeviceQueryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.query"
}

func (r YunosTvpubadminDeviceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDeviceQueryRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r YunosTvpubadminDeviceQueryRequest) GetTerminalType() string {
    return r.terminalType
}

func (r *YunosTvpubadminDeviceQueryRequest) SetBrandId(brandId int64) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

func (r YunosTvpubadminDeviceQueryRequest) GetBrandId() int64 {
    return r.brandId
}

func (r *YunosTvpubadminDeviceQueryRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminDeviceQueryRequest) GetLicense() int64 {
    return r.license
}

func (r *YunosTvpubadminDeviceQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r YunosTvpubadminDeviceQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *YunosTvpubadminDeviceQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r YunosTvpubadminDeviceQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

