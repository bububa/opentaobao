package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单视频审核获取视频列表 API请求
yunos.tvpubadmin.content.single.video.getlist

牌照方在审核单视频时获取单视频列表接口
*/
type YunosTvpubadminContentSingleVideoGetlistRequest struct {
    model.Params
    // 视频外部来源类型: 1:YOUKU, 2:MONGO_TV, 3:TAOTVMEDIA, 4:GOLIVE
    _extType   int64
    // 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
    _licenseState   int64
    // 单页数量
    _pageSize   int64
    // 查询时间范围，结束时间
    _gmtEnd   string
    // 视频id
    _extVideoStrId   string
    // 查询多个审核状态
    _licenseStateList   []int64
    // 时间类型：1-licenseSubmitTime, 2-licenseAuditTime, 3-youkuPublishTime
    _dateType   int64
    // 主分类
    _category   int64
    // 页码
    _pageNo   int64
    // 查询时间范围，开始时间
    _gmtStart   string
    // 牌照方
    _license   int64
    // 视屏名称
    _videoTitleLike   string
    // 审核优先级，紧急4，高3，中2，低1
    _priority   int64
}

// 初始化YunosTvpubadminContentSingleVideoGetlistRequest对象
func NewYunosTvpubadminContentSingleVideoGetlistRequest() *YunosTvpubadminContentSingleVideoGetlistRequest{
    return &YunosTvpubadminContentSingleVideoGetlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.single.video.getlist"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtType Setter
// 视频外部来源类型: 1:YOUKU, 2:MONGO_TV, 3:TAOTVMEDIA, 4:GOLIVE
func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetExtType(_extType int64) error {
    r._extType = _extType
    r.Set("ext_type", _extType)
    return nil
}

// ExtType Getter
func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetExtType() int64 {
    return r._extType
}
// LicenseState Setter
// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetLicenseState(_licenseState int64) error {
    r._licenseState = _licenseState
    r.Set("license_state", _licenseState)
    return nil
}

// LicenseState Getter
func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetLicenseState() int64 {
    return r._licenseState
}
// PageSize Setter
// 单页数量
func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetPageSize() int64 {
    return r._pageSize
}
// GmtEnd Setter
// 查询时间范围，结束时间
func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetGmtEnd(_gmtEnd string) error {
    r._gmtEnd = _gmtEnd
    r.Set("gmt_end", _gmtEnd)
    return nil
}

// GmtEnd Getter
func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetGmtEnd() string {
    return r._gmtEnd
}
// ExtVideoStrId Setter
// 视频id
func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetExtVideoStrId(_extVideoStrId string) error {
    r._extVideoStrId = _extVideoStrId
    r.Set("ext_video_str_id", _extVideoStrId)
    return nil
}

// ExtVideoStrId Getter
func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetExtVideoStrId() string {
    return r._extVideoStrId
}
// LicenseStateList Setter
// 查询多个审核状态
func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetLicenseStateList(_licenseStateList []int64) error {
    r._licenseStateList = _licenseStateList
    r.Set("license_state_list", _licenseStateList)
    return nil
}

// LicenseStateList Getter
func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetLicenseStateList() []int64 {
    return r._licenseStateList
}
// DateType Setter
// 时间类型：1-licenseSubmitTime, 2-licenseAuditTime, 3-youkuPublishTime
func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetDateType(_dateType int64) error {
    r._dateType = _dateType
    r.Set("date_type", _dateType)
    return nil
}

// DateType Getter
func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetDateType() int64 {
    return r._dateType
}
// Category Setter
// 主分类
func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetCategory(_category int64) error {
    r._category = _category
    r.Set("category", _category)
    return nil
}

// Category Getter
func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetCategory() int64 {
    return r._category
}
// PageNo Setter
// 页码
func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetPageNo() int64 {
    return r._pageNo
}
// GmtStart Setter
// 查询时间范围，开始时间
func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetGmtStart(_gmtStart string) error {
    r._gmtStart = _gmtStart
    r.Set("gmt_start", _gmtStart)
    return nil
}

// GmtStart Getter
func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetGmtStart() string {
    return r._gmtStart
}
// License Setter
// 牌照方
func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetLicense() int64 {
    return r._license
}
// VideoTitleLike Setter
// 视屏名称
func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetVideoTitleLike(_videoTitleLike string) error {
    r._videoTitleLike = _videoTitleLike
    r.Set("video_title_like", _videoTitleLike)
    return nil
}

// VideoTitleLike Getter
func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetVideoTitleLike() string {
    return r._videoTitleLike
}
// Priority Setter
// 审核优先级，紧急4，高3，中2，低1
func (r *YunosTvpubadminContentSingleVideoGetlistRequest) SetPriority(_priority int64) error {
    r._priority = _priority
    r.Set("priority", _priority)
    return nil
}

// Priority Getter
func (r YunosTvpubadminContentSingleVideoGetlistRequest) GetPriority() int64 {
    return r._priority
}
