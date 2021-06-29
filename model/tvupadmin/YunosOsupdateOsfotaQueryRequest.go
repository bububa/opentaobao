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
type YunosOsupdateOsfotaQueryRequest struct {
    model.Params
    // 设备型号ID
    modleId   int64
    // 页码
    page   int64
    // 每页数量
    pageSize   int64
}

// 初始化YunosOsupdateOsfotaQueryRequest对象
func NewYunosOsupdateOsfotaQueryRequest() *YunosOsupdateOsfotaQueryRequest{
    return &YunosOsupdateOsfotaQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosOsupdateOsfotaQueryRequest) GetApiMethodName() string {
    return "yunos.osupdate.osfota.query"
}

// IRequest interface 方法, 获取API参数
func (r YunosOsupdateOsfotaQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ModleId Setter
// 设备型号ID
func (r *YunosOsupdateOsfotaQueryRequest) SetModleId(modleId int64) error {
    r.modleId = modleId
    r.Set("modle_id", modleId)
    return nil
}

// ModleId Getter
func (r YunosOsupdateOsfotaQueryRequest) GetModleId() int64 {
    return r.modleId
}
// Page Setter
// 页码
func (r *YunosOsupdateOsfotaQueryRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r YunosOsupdateOsfotaQueryRequest) GetPage() int64 {
    return r.page
}
// PageSize Setter
// 每页数量
func (r *YunosOsupdateOsfotaQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r YunosOsupdateOsfotaQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
