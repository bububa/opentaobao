package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
系统升级分页查询 API请求
yunos.osupdate.osfota.query

分页查询osoupdate系统升级列表
*/
type YunosOsupdateOsfotaQueryAPIRequest struct {
    model.Params
    // 设备型号ID
    _modleId   int64
    // 页码
    _page   int64
    // 每页数量
    _pageSize   int64
}

// 初始化YunosOsupdateOsfotaQueryAPIRequest对象
func NewYunosOsupdateOsfotaQueryRequest() *YunosOsupdateOsfotaQueryAPIRequest{
    return &YunosOsupdateOsfotaQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosOsupdateOsfotaQueryAPIRequest) GetApiMethodName() string {
    return "yunos.osupdate.osfota.query"
}

// IRequest interface 方法, 获取API参数
func (r YunosOsupdateOsfotaQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ModleId Setter
// 设备型号ID
func (r *YunosOsupdateOsfotaQueryAPIRequest) SetModleId(_modleId int64) error {
    r._modleId = _modleId
    r.Set("modle_id", _modleId)
    return nil
}

// ModleId Getter
func (r YunosOsupdateOsfotaQueryAPIRequest) GetModleId() int64 {
    return r._modleId
}
// Page Setter
// 页码
func (r *YunosOsupdateOsfotaQueryAPIRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r YunosOsupdateOsfotaQueryAPIRequest) GetPage() int64 {
    return r._page
}
// PageSize Setter
// 每页数量
func (r *YunosOsupdateOsfotaQueryAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosOsupdateOsfotaQueryAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
