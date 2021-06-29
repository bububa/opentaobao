package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
节目审核获取节目列表 API请求
yunos.tvpubadmin.content.show.getlist

节目审核获取节目列表
*/
type YunosTvpubadminContentShowGetlistRequest struct {
    model.Params
    // 视频外部来源类型: 1:YOUKU, 2:MONGO_TV, 3:TAOTVMEDIA, 4:GOLIVE
    extType   int64
    // 时间查询范围，结束时间
    gmtEnd   string
    // 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
    licenseState   int64
    // 分页
    pageSize   int64
    // 节目ID
    showId   string
    // 视频ID
    extVideoStrId   string
    // 时间类型：1-licenseSubmitTime, 2-licenseAuditTime, 3-youkuPublishTime
    dateType   int64
    // 主分类
    category   int64
    // 节目名称
    showName   string
    // 分页，页码
    pageNo   int64
    // 时间查询范围，开始时间
    gmtStart   string
    // 牌照方
    license   int64
    // 视频名称
    videoTitleLike   string
    // 老媒资ID(program_id)
    vmacLongId   int64
}

// 初始化YunosTvpubadminContentShowGetlistRequest对象
func NewYunosTvpubadminContentShowGetlistRequest() *YunosTvpubadminContentShowGetlistRequest{
    return &YunosTvpubadminContentShowGetlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowGetlistRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.show.getlist"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowGetlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtType Setter
// 视频外部来源类型: 1:YOUKU, 2:MONGO_TV, 3:TAOTVMEDIA, 4:GOLIVE
func (r *YunosTvpubadminContentShowGetlistRequest) SetExtType(extType int64) error {
    r.extType = extType
    r.Set("ext_type", extType)
    return nil
}

// ExtType Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetExtType() int64 {
    return r.extType
}
// GmtEnd Setter
// 时间查询范围，结束时间
func (r *YunosTvpubadminContentShowGetlistRequest) SetGmtEnd(gmtEnd string) error {
    r.gmtEnd = gmtEnd
    r.Set("gmt_end", gmtEnd)
    return nil
}

// GmtEnd Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetGmtEnd() string {
    return r.gmtEnd
}
// LicenseState Setter
// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
func (r *YunosTvpubadminContentShowGetlistRequest) SetLicenseState(licenseState int64) error {
    r.licenseState = licenseState
    r.Set("license_state", licenseState)
    return nil
}

// LicenseState Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetLicenseState() int64 {
    return r.licenseState
}
// PageSize Setter
// 分页
func (r *YunosTvpubadminContentShowGetlistRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetPageSize() int64 {
    return r.pageSize
}
// ShowId Setter
// 节目ID
func (r *YunosTvpubadminContentShowGetlistRequest) SetShowId(showId string) error {
    r.showId = showId
    r.Set("show_id", showId)
    return nil
}

// ShowId Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetShowId() string {
    return r.showId
}
// ExtVideoStrId Setter
// 视频ID
func (r *YunosTvpubadminContentShowGetlistRequest) SetExtVideoStrId(extVideoStrId string) error {
    r.extVideoStrId = extVideoStrId
    r.Set("ext_video_str_id", extVideoStrId)
    return nil
}

// ExtVideoStrId Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetExtVideoStrId() string {
    return r.extVideoStrId
}
// DateType Setter
// 时间类型：1-licenseSubmitTime, 2-licenseAuditTime, 3-youkuPublishTime
func (r *YunosTvpubadminContentShowGetlistRequest) SetDateType(dateType int64) error {
    r.dateType = dateType
    r.Set("date_type", dateType)
    return nil
}

// DateType Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetDateType() int64 {
    return r.dateType
}
// Category Setter
// 主分类
func (r *YunosTvpubadminContentShowGetlistRequest) SetCategory(category int64) error {
    r.category = category
    r.Set("category", category)
    return nil
}

// Category Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetCategory() int64 {
    return r.category
}
// ShowName Setter
// 节目名称
func (r *YunosTvpubadminContentShowGetlistRequest) SetShowName(showName string) error {
    r.showName = showName
    r.Set("show_name", showName)
    return nil
}

// ShowName Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetShowName() string {
    return r.showName
}
// PageNo Setter
// 分页，页码
func (r *YunosTvpubadminContentShowGetlistRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetPageNo() int64 {
    return r.pageNo
}
// GmtStart Setter
// 时间查询范围，开始时间
func (r *YunosTvpubadminContentShowGetlistRequest) SetGmtStart(gmtStart string) error {
    r.gmtStart = gmtStart
    r.Set("gmt_start", gmtStart)
    return nil
}

// GmtStart Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetGmtStart() string {
    return r.gmtStart
}
// License Setter
// 牌照方
func (r *YunosTvpubadminContentShowGetlistRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

// License Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetLicense() int64 {
    return r.license
}
// VideoTitleLike Setter
// 视频名称
func (r *YunosTvpubadminContentShowGetlistRequest) SetVideoTitleLike(videoTitleLike string) error {
    r.videoTitleLike = videoTitleLike
    r.Set("video_title_like", videoTitleLike)
    return nil
}

// VideoTitleLike Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetVideoTitleLike() string {
    return r.videoTitleLike
}
// VmacLongId Setter
// 老媒资ID(program_id)
func (r *YunosTvpubadminContentShowGetlistRequest) SetVmacLongId(vmacLongId int64) error {
    r.vmacLongId = vmacLongId
    r.Set("vmac_long_id", vmacLongId)
    return nil
}

// VmacLongId Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetVmacLongId() int64 {
    return r.vmacLongId
}
