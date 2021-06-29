package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
广告牌照管控查询 APIRequest
yunos.tvpubadmin.content.advert.queryschedule

广告牌照管控查询
*/
type YunosTvpubadminContentAdvertQueryscheduleRequest struct {
    model.Params

    // 查询范围: 1-牌照，4-uuid
    range   int64 

    // 分页，页码
    pageNo   int64 

    // 分页，单页数量
    pageSize   int64 

    // 日期
    gmtStart   string 

    // uuid
    uuid   string 

    // 牌照方，1-华数，7-CIBN
    license   int64 

    // 广告类型
    sityType   int64 

}

func NewYunosTvpubadminContentAdvertQueryscheduleRequest() *YunosTvpubadminContentAdvertQueryscheduleRequest{
    return &YunosTvpubadminContentAdvertQueryscheduleRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.advert.queryschedule"
}

func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentAdvertQueryscheduleRequest) SetRange(range int64) error {
    r.range = range
    r.Set("range", range)
    return nil
}

func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetRange() int64 {
    return r.range
}

func (r *YunosTvpubadminContentAdvertQueryscheduleRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *YunosTvpubadminContentAdvertQueryscheduleRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *YunosTvpubadminContentAdvertQueryscheduleRequest) SetGmtStart(gmtStart string) error {
    r.gmtStart = gmtStart
    r.Set("gmt_start", gmtStart)
    return nil
}

func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetGmtStart() string {
    return r.gmtStart
}

func (r *YunosTvpubadminContentAdvertQueryscheduleRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetUuid() string {
    return r.uuid
}

func (r *YunosTvpubadminContentAdvertQueryscheduleRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetLicense() int64 {
    return r.license
}

func (r *YunosTvpubadminContentAdvertQueryscheduleRequest) SetSityType(sityType int64) error {
    r.sityType = sityType
    r.Set("sity_type", sityType)
    return nil
}

func (r YunosTvpubadminContentAdvertQueryscheduleRequest) GetSityType() int64 {
    return r.sityType
}

