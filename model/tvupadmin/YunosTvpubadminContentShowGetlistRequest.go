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
    _extType   int64
    // 时间查询范围，结束时间
    _gmtEnd   string
    // 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
    _licenseState   int64
    // 分页
    _pageSize   int64
    // 节目ID
    _showId   string
    // 视频ID
    _extVideoStrId   string
    // 时间类型：1-licenseSubmitTime, 2-licenseAuditTime, 3-youkuPublishTime
    _dateType   int64
    // 主分类
    _category   int64
    // 节目名称
    _showName   string
    // 分页，页码
    _pageNo   int64
    // 时间查询范围，开始时间
    _gmtStart   string
    // 牌照方
    _license   int64
    // 视频名称
    _videoTitleLike   string
    // 老媒资ID(program_id)
    _vmacLongId   int64
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
func (r *YunosTvpubadminContentShowGetlistRequest) SetExtType(_extType int64) error {
    r._extType = _extType
    r.Set("ext_type", _extType)
    return nil
}

// ExtType Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetExtType() int64 {
    return r._extType
}
// GmtEnd Setter
// 时间查询范围，结束时间
func (r *YunosTvpubadminContentShowGetlistRequest) SetGmtEnd(_gmtEnd string) error {
    r._gmtEnd = _gmtEnd
    r.Set("gmt_end", _gmtEnd)
    return nil
}

// GmtEnd Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetGmtEnd() string {
    return r._gmtEnd
}
// LicenseState Setter
// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
func (r *YunosTvpubadminContentShowGetlistRequest) SetLicenseState(_licenseState int64) error {
    r._licenseState = _licenseState
    r.Set("license_state", _licenseState)
    return nil
}

// LicenseState Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetLicenseState() int64 {
    return r._licenseState
}
// PageSize Setter
// 分页
func (r *YunosTvpubadminContentShowGetlistRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetPageSize() int64 {
    return r._pageSize
}
// ShowId Setter
// 节目ID
func (r *YunosTvpubadminContentShowGetlistRequest) SetShowId(_showId string) error {
    r._showId = _showId
    r.Set("show_id", _showId)
    return nil
}

// ShowId Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetShowId() string {
    return r._showId
}
// ExtVideoStrId Setter
// 视频ID
func (r *YunosTvpubadminContentShowGetlistRequest) SetExtVideoStrId(_extVideoStrId string) error {
    r._extVideoStrId = _extVideoStrId
    r.Set("ext_video_str_id", _extVideoStrId)
    return nil
}

// ExtVideoStrId Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetExtVideoStrId() string {
    return r._extVideoStrId
}
// DateType Setter
// 时间类型：1-licenseSubmitTime, 2-licenseAuditTime, 3-youkuPublishTime
func (r *YunosTvpubadminContentShowGetlistRequest) SetDateType(_dateType int64) error {
    r._dateType = _dateType
    r.Set("date_type", _dateType)
    return nil
}

// DateType Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetDateType() int64 {
    return r._dateType
}
// Category Setter
// 主分类
func (r *YunosTvpubadminContentShowGetlistRequest) SetCategory(_category int64) error {
    r._category = _category
    r.Set("category", _category)
    return nil
}

// Category Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetCategory() int64 {
    return r._category
}
// ShowName Setter
// 节目名称
func (r *YunosTvpubadminContentShowGetlistRequest) SetShowName(_showName string) error {
    r._showName = _showName
    r.Set("show_name", _showName)
    return nil
}

// ShowName Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetShowName() string {
    return r._showName
}
// PageNo Setter
// 分页，页码
func (r *YunosTvpubadminContentShowGetlistRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetPageNo() int64 {
    return r._pageNo
}
// GmtStart Setter
// 时间查询范围，开始时间
func (r *YunosTvpubadminContentShowGetlistRequest) SetGmtStart(_gmtStart string) error {
    r._gmtStart = _gmtStart
    r.Set("gmt_start", _gmtStart)
    return nil
}

// GmtStart Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetGmtStart() string {
    return r._gmtStart
}
// License Setter
// 牌照方
func (r *YunosTvpubadminContentShowGetlistRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetLicense() int64 {
    return r._license
}
// VideoTitleLike Setter
// 视频名称
func (r *YunosTvpubadminContentShowGetlistRequest) SetVideoTitleLike(_videoTitleLike string) error {
    r._videoTitleLike = _videoTitleLike
    r.Set("video_title_like", _videoTitleLike)
    return nil
}

// VideoTitleLike Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetVideoTitleLike() string {
    return r._videoTitleLike
}
// VmacLongId Setter
// 老媒资ID(program_id)
func (r *YunosTvpubadminContentShowGetlistRequest) SetVmacLongId(_vmacLongId int64) error {
    r._vmacLongId = _vmacLongId
    r.Set("vmac_long_id", _vmacLongId)
    return nil
}

// VmacLongId Getter
func (r YunosTvpubadminContentShowGetlistRequest) GetVmacLongId() int64 {
    return r._vmacLongId
}
