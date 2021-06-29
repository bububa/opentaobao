package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户订单列表 APIRequest
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

func NewYunosTvpubadminUserOrderlistRequest() *YunosTvpubadminUserOrderlistRequest{
    return &YunosTvpubadminUserOrderlistRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminUserOrderlistRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.user.orderlist"
}

func (r YunosTvpubadminUserOrderlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminUserOrderlistRequest) SetUid(uid int64) error {
    r.uid = uid
    r.Set("uid", uid)
    return nil
}

func (r YunosTvpubadminUserOrderlistRequest) GetUid() int64 {
    return r.uid
}

func (r *YunosTvpubadminUserOrderlistRequest) SetCreateTimeStartStr(createTimeStartStr string) error {
    r.createTimeStartStr = createTimeStartStr
    r.Set("create_time_start_str", createTimeStartStr)
    return nil
}

func (r YunosTvpubadminUserOrderlistRequest) GetCreateTimeStartStr() string {
    return r.createTimeStartStr
}

func (r *YunosTvpubadminUserOrderlistRequest) SetCreateTimeEndStr(createTimeEndStr string) error {
    r.createTimeEndStr = createTimeEndStr
    r.Set("create_time_end_str", createTimeEndStr)
    return nil
}

func (r YunosTvpubadminUserOrderlistRequest) GetCreateTimeEndStr() string {
    return r.createTimeEndStr
}

func (r *YunosTvpubadminUserOrderlistRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminUserOrderlistRequest) GetLicense() int64 {
    return r.license
}

func (r *YunosTvpubadminUserOrderlistRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r YunosTvpubadminUserOrderlistRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *YunosTvpubadminUserOrderlistRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r YunosTvpubadminUserOrderlistRequest) GetPageSize() int64 {
    return r.pageSize
}

