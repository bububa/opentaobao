package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentShowGetlistAPIRequest 节目审核获取节目列表 API请求
// yunos.tvpubadmin.content.show.getlist
//
// 节目审核获取节目列表
type YunosTvpubadminContentShowGetlistAPIRequest struct {
	model.Params
	// 时间查询范围，结束时间
	_gmtEnd string
	// 节目ID
	_showId string
	// 视频ID
	_extVideoStrId string
	// 节目名称
	_showName string
	// 时间查询范围，开始时间
	_gmtStart string
	// 视频名称
	_videoTitleLike string
	// 视频外部来源类型: 1:YOUKU, 2:MONGO_TV, 3:TAOTVMEDIA, 4:GOLIVE
	_extType int64
	// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
	_licenseState int64
	// 分页
	_pageSize int64
	// 时间类型：1-licenseSubmitTime, 2-licenseAuditTime, 3-youkuPublishTime
	_dateType int64
	// 主分类
	_category int64
	// 分页，页码
	_pageNo int64
	// 牌照方
	_license int64
	// 老媒资ID(program_id)
	_vmacLongId int64
}

// NewYunosTvpubadminContentShowGetlistRequest 初始化YunosTvpubadminContentShowGetlistAPIRequest对象
func NewYunosTvpubadminContentShowGetlistRequest() *YunosTvpubadminContentShowGetlistAPIRequest {
	return &YunosTvpubadminContentShowGetlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.show.getlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGmtEnd is GmtEnd Setter
// 时间查询范围，结束时间
func (r *YunosTvpubadminContentShowGetlistAPIRequest) SetGmtEnd(_gmtEnd string) error {
	r._gmtEnd = _gmtEnd
	r.Set("gmt_end", _gmtEnd)
	return nil
}

// GetGmtEnd GmtEnd Getter
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetGmtEnd() string {
	return r._gmtEnd
}

// SetShowId is ShowId Setter
// 节目ID
func (r *YunosTvpubadminContentShowGetlistAPIRequest) SetShowId(_showId string) error {
	r._showId = _showId
	r.Set("show_id", _showId)
	return nil
}

// GetShowId ShowId Getter
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetShowId() string {
	return r._showId
}

// SetExtVideoStrId is ExtVideoStrId Setter
// 视频ID
func (r *YunosTvpubadminContentShowGetlistAPIRequest) SetExtVideoStrId(_extVideoStrId string) error {
	r._extVideoStrId = _extVideoStrId
	r.Set("ext_video_str_id", _extVideoStrId)
	return nil
}

// GetExtVideoStrId ExtVideoStrId Getter
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetExtVideoStrId() string {
	return r._extVideoStrId
}

// SetShowName is ShowName Setter
// 节目名称
func (r *YunosTvpubadminContentShowGetlistAPIRequest) SetShowName(_showName string) error {
	r._showName = _showName
	r.Set("show_name", _showName)
	return nil
}

// GetShowName ShowName Getter
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetShowName() string {
	return r._showName
}

// SetGmtStart is GmtStart Setter
// 时间查询范围，开始时间
func (r *YunosTvpubadminContentShowGetlistAPIRequest) SetGmtStart(_gmtStart string) error {
	r._gmtStart = _gmtStart
	r.Set("gmt_start", _gmtStart)
	return nil
}

// GetGmtStart GmtStart Getter
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetGmtStart() string {
	return r._gmtStart
}

// SetVideoTitleLike is VideoTitleLike Setter
// 视频名称
func (r *YunosTvpubadminContentShowGetlistAPIRequest) SetVideoTitleLike(_videoTitleLike string) error {
	r._videoTitleLike = _videoTitleLike
	r.Set("video_title_like", _videoTitleLike)
	return nil
}

// GetVideoTitleLike VideoTitleLike Getter
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetVideoTitleLike() string {
	return r._videoTitleLike
}

// SetExtType is ExtType Setter
// 视频外部来源类型: 1:YOUKU, 2:MONGO_TV, 3:TAOTVMEDIA, 4:GOLIVE
func (r *YunosTvpubadminContentShowGetlistAPIRequest) SetExtType(_extType int64) error {
	r._extType = _extType
	r.Set("ext_type", _extType)
	return nil
}

// GetExtType ExtType Getter
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetExtType() int64 {
	return r._extType
}

// SetLicenseState is LicenseState Setter
// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
func (r *YunosTvpubadminContentShowGetlistAPIRequest) SetLicenseState(_licenseState int64) error {
	r._licenseState = _licenseState
	r.Set("license_state", _licenseState)
	return nil
}

// GetLicenseState LicenseState Getter
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetLicenseState() int64 {
	return r._licenseState
}

// SetPageSize is PageSize Setter
// 分页
func (r *YunosTvpubadminContentShowGetlistAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetDateType is DateType Setter
// 时间类型：1-licenseSubmitTime, 2-licenseAuditTime, 3-youkuPublishTime
func (r *YunosTvpubadminContentShowGetlistAPIRequest) SetDateType(_dateType int64) error {
	r._dateType = _dateType
	r.Set("date_type", _dateType)
	return nil
}

// GetDateType DateType Getter
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetDateType() int64 {
	return r._dateType
}

// SetCategory is Category Setter
// 主分类
func (r *YunosTvpubadminContentShowGetlistAPIRequest) SetCategory(_category int64) error {
	r._category = _category
	r.Set("category", _category)
	return nil
}

// GetCategory Category Getter
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetCategory() int64 {
	return r._category
}

// SetPageNo is PageNo Setter
// 分页，页码
func (r *YunosTvpubadminContentShowGetlistAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetLicense is License Setter
// 牌照方
func (r *YunosTvpubadminContentShowGetlistAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetLicense() int64 {
	return r._license
}

// SetVmacLongId is VmacLongId Setter
// 老媒资ID(program_id)
func (r *YunosTvpubadminContentShowGetlistAPIRequest) SetVmacLongId(_vmacLongId int64) error {
	r._vmacLongId = _vmacLongId
	r.Set("vmac_long_id", _vmacLongId)
	return nil
}

// GetVmacLongId VmacLongId Getter
func (r YunosTvpubadminContentShowGetlistAPIRequest) GetVmacLongId() int64 {
	return r._vmacLongId
}
