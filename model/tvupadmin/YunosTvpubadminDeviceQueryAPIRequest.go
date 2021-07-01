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
type YunosTvpubadminDeviceQueryAPIRequest struct {
    model.Params
    // 终端类型
    _terminalType   string
    // 品牌ID
    _brandId   int64
    // 牌照方
    _license   int64
    // 页码值
    _pageNo   int64
    // 每页条数
    _pageSize   int64
}

// 初始化YunosTvpubadminDeviceQueryAPIRequest对象
func NewYunosTvpubadminDeviceQueryRequest() *YunosTvpubadminDeviceQueryAPIRequest{
    return &YunosTvpubadminDeviceQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceQueryAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.device.query"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TerminalType Setter
// 终端类型
func (r *YunosTvpubadminDeviceQueryAPIRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r YunosTvpubadminDeviceQueryAPIRequest) GetTerminalType() string {
    return r._terminalType
}
// BrandId Setter
// 品牌ID
func (r *YunosTvpubadminDeviceQueryAPIRequest) SetBrandId(_brandId int64) error {
    r._brandId = _brandId
    r.Set("brand_id", _brandId)
    return nil
}

// BrandId Getter
func (r YunosTvpubadminDeviceQueryAPIRequest) GetBrandId() int64 {
    return r._brandId
}
// License Setter
// 牌照方
func (r *YunosTvpubadminDeviceQueryAPIRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminDeviceQueryAPIRequest) GetLicense() int64 {
    return r._license
}
// PageNo Setter
// 页码值
func (r *YunosTvpubadminDeviceQueryAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminDeviceQueryAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页条数
func (r *YunosTvpubadminDeviceQueryAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminDeviceQueryAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
