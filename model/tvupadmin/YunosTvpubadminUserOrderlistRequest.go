package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户订单列表 API请求
yunos.tvpubadmin.user.orderlist

获取用户订单列表
*/
type YunosTvpubadminUserOrderlistRequest struct {
    model.Params
    // 用户ID
    uid   int64
    // 开始时间
    createTimeStartStr   string
    // 结束时间
    createTimeEndStr   string
    // 牌照方
    license   int64
    // 页码值
    pageNo   int64
    // 每页行数
    pageSize   int64
}

// 初始化YunosTvpubadminUserOrderlistRequest对象
func NewYunosTvpubadminUserOrderlistRequest() *YunosTvpubadminUserOrderlistRequest{
    return &YunosTvpubadminUserOrderlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminUserOrderlistRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.user.orderlist"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminUserOrderlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uid Setter
// 用户ID
func (r *YunosTvpubadminUserOrderlistRequest) SetUid(uid int64) error {
    r.uid = uid
    r.Set("uid", uid)
    return nil
}

// Uid Getter
func (r YunosTvpubadminUserOrderlistRequest) GetUid() int64 {
    return r.uid
}
// CreateTimeStartStr Setter
// 开始时间
func (r *YunosTvpubadminUserOrderlistRequest) SetCreateTimeStartStr(createTimeStartStr string) error {
    r.createTimeStartStr = createTimeStartStr
    r.Set("create_time_start_str", createTimeStartStr)
    return nil
}

// CreateTimeStartStr Getter
func (r YunosTvpubadminUserOrderlistRequest) GetCreateTimeStartStr() string {
    return r.createTimeStartStr
}
// CreateTimeEndStr Setter
// 结束时间
func (r *YunosTvpubadminUserOrderlistRequest) SetCreateTimeEndStr(createTimeEndStr string) error {
    r.createTimeEndStr = createTimeEndStr
    r.Set("create_time_end_str", createTimeEndStr)
    return nil
}

// CreateTimeEndStr Getter
func (r YunosTvpubadminUserOrderlistRequest) GetCreateTimeEndStr() string {
    return r.createTimeEndStr
}
// License Setter
// 牌照方
func (r *YunosTvpubadminUserOrderlistRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

// License Getter
func (r YunosTvpubadminUserOrderlistRequest) GetLicense() int64 {
    return r.license
}
// PageNo Setter
// 页码值
func (r *YunosTvpubadminUserOrderlistRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminUserOrderlistRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页行数
func (r *YunosTvpubadminUserOrderlistRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminUserOrderlistRequest) GetPageSize() int64 {
    return r.pageSize
}
