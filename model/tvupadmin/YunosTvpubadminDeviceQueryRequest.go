package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备列表 API请求
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

// 初始化YunosTvpubadminDeviceQueryRequest对象
func NewYunosTvpubadminDeviceQueryRequest() *YunosTvpubadminDeviceQueryRequest{
    return &YunosTvpubadminDeviceQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceQueryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.query"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TerminalType Setter
// 终端类型
func (r *YunosTvpubadminDeviceQueryRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

// TerminalType Getter
func (r YunosTvpubadminDeviceQueryRequest) GetTerminalType() string {
    return r.terminalType
}
// BrandId Setter
// 品牌ID
func (r *YunosTvpubadminDeviceQueryRequest) SetBrandId(brandId int64) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

// BrandId Getter
func (r YunosTvpubadminDeviceQueryRequest) GetBrandId() int64 {
    return r.brandId
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDeviceQueryRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceQueryRequest) GetLicense() int64 {
    return r.license
}
// PageNo Setter
// 页码值
func (r *YunosTvpubadminDeviceQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminDeviceQueryRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页条数
func (r *YunosTvpubadminDeviceQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminDeviceQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
