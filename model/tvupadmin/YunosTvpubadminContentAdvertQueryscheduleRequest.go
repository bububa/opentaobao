package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
广告牌照管控查询 API请求
yunos.tvpubadmin.content.advert.queryschedule

广告牌照管控查询
*/
type YunosTvpubadminContentAdvertQueryscheduleRequest struct {
    model.Params
    // 查询范围: 1-牌照，4-uuid
    _range   int64
    // 分页，页码
    _pageNo   int64
    // 分页，单页数量
    _pageSize   int64
    // 日期
    _gmtStart   string
    // uuid
    _uuid   string
    // 牌照方，1-华数，7-CIBN
    _license   int64
    // 广告类型
    _sityType   int64
}

// 初始化YunosTvpubadminContentAdvertQueryscheduleRequest对象
func NewYunosTvpubadminContentAdvertQueryscheduleRequest() *YunosTvpubadminContentAdvertQueryscheduleRequest{
    return &YunosTvpubadminContentAdvertQueryscheduleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.advert.queryschedule"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Range Setter
// 查询范围: 1-牌照，4-uuid
func (r *YunosTvpubadminContentAdvertQueryscheduleRequest) SetRange(_range int64) error {
    r._range = _range
    r.Set("range", _range)
    return nil
}

// Range Getter
func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetRange() int64 {
    return r._range
}
// PageNo Setter
// 分页，页码
func (r *YunosTvpubadminContentAdvertQueryscheduleRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 分页，单页数量
func (r *YunosTvpubadminContentAdvertQueryscheduleRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetPageSize() int64 {
    return r._pageSize
}
// GmtStart Setter
// 日期
func (r *YunosTvpubadminContentAdvertQueryscheduleRequest) SetGmtStart(_gmtStart string) error {
    r._gmtStart = _gmtStart
    r.Set("gmt_start", _gmtStart)
    return nil
}

// GmtStart Getter
func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetGmtStart() string {
    return r._gmtStart
}
// Uuid Setter
// uuid
func (r *YunosTvpubadminContentAdvertQueryscheduleRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetUuid() string {
    return r._uuid
}
// License Setter
// 牌照方，1-华数，7-CIBN
func (r *YunosTvpubadminContentAdvertQueryscheduleRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetLicense() int64 {
    return r._license
}
// SityType Setter
// 广告类型
func (r *YunosTvpubadminContentAdvertQueryscheduleRequest) SetSityType(_sityType int64) error {
    r._sityType = _sityType
    r.Set("sity_type", _sityType)
    return nil
}

// SityType Getter
func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetSityType() int64 {
    return r._sityType
}
