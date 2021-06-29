package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单视频审核获取视频列表 APIRequest
yunos.tvpubadmin.content.single.video.getlist

牌照方在审核单视频时获取单视频列表接口
*/
type YunosTvpubadminContentSingleVideoGetlistRequest struct {
    model.Params

    // 视频外部来源类型: 1:YOUKU, 2:MONGO_TV, 3:TAOTVMEDIA, 4:GOLIVE
    extType   int64 

    // 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
    licenseState   int64 

    // 单页数量
    pageSize   int64 

    // 查询时间范围，结束时间
    gmtEnd   string 

    // 视频id
    extVideoStrId   string 

    // 查询多个审核状态
    licenseStateList   []int64 

    // 时间类型：1-licenseSubmitTime, 2-licenseAuditTime, 3-youkuPublishTime
    dateType   int64 

    // 主分类
    category   int64 

    // 页码
    pageNo   int64 

    // 查询时间范围，开始时间
    gmtStart   string 

    // 牌照方
    license   int64 

    // 视屏名称
    videoTitleLike   string 

    // 审核优先级，紧急4，高3，中2，低1
    priority   int64 

}

func NewYunosTvpubadminContentSingleVideoGetlistRequest() *YunosTvpubadminContentSingleVideoGetlistRequest{
    return &YunosTvpubadminContentSingleVideoGetlistRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.single.video.getlist"
}

func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetExtType(extType int64) error {
    r.extType = extType
    r.Set("ext_type", extType)
    return nil
}

func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetExtType() int64 {
    return r.extType
}

func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetLicenseState(licenseState int64) error {
    r.licenseState = licenseState
    r.Set("license_state", licenseState)
    return nil
}

func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetLicenseState() int64 {
    return r.licenseState
}

func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetGmtEnd(gmtEnd string) error {
    r.gmtEnd = gmtEnd
    r.Set("gmt_end", gmtEnd)
    return nil
}

func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetGmtEnd() string {
    return r.gmtEnd
}

func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetExtVideoStrId(extVideoStrId string) error {
    r.extVideoStrId = extVideoStrId
    r.Set("ext_video_str_id", extVideoStrId)
    return nil
}

func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetExtVideoStrId() string {
    return r.extVideoStrId
}

func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetLicenseStateList(licenseStateList []int64) error {
    r.licenseStateList = licenseStateList
    r.Set("license_state_list", licenseStateList)
    return nil
}

func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetLicenseStateList() []int64 {
    return r.licenseStateList
}

func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetDateType(dateType int64) error {
    r.dateType = dateType
    r.Set("date_type", dateType)
    return nil
}

func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetDateType() int64 {
    return r.dateType
}

func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetCategory(category int64) error {
    r.category = category
    r.Set("category", category)
    return nil
}

func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetCategory() int64 {
    return r.category
}

func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetGmtStart(gmtStart string) error {
    r.gmtStart = gmtStart
    r.Set("gmt_start", gmtStart)
    return nil
}

func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetGmtStart() string {
    return r.gmtStart
}

func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetLicense() int64 {
    return r.license
}

func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetVideoTitleLike(videoTitleLike string) error {
    r.videoTitleLike = videoTitleLike
    r.Set("video_title_like", videoTitleLike)
    return nil
}

func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetVideoTitleLike() string {
    return r.videoTitleLike
}

func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetPriority(priority int64) error {
    r.priority = priority
    r.Set("priority", priority)
    return nil
}

func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetPriority() int64 {
    return r.priority
}

