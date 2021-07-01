package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentAdvertQueryscheduleAPIRequest
广告牌照管控查询 API请求
yunos.tvpubadmin.content.advert.queryschedule

广告牌照管控查询 */
type YunosTvpubadminContentAdvertQueryscheduleAPIRequest struct {
	model.Params
	// 查询范围: 1-牌照，4-uuid
	_range int64
	// 分页，页码
	_pageNo int64
	// 分页，单页数量
	_pageSize int64
	// 日期
	_gmtStart string
	// uuid
	_uuid string
	// 牌照方，1-华数，7-CIBN
	_license int64
	// 广告类型
	_sityType int64
}

// NewYunosTvpubadminContentAdvertQueryscheduleRequest 初始化YunosTvpubadminContentAdvertQueryscheduleAPIRequest对象
func NewYunosTvpubadminContentAdvertQueryscheduleRequest() *YunosTvpubadminContentAdvertQueryscheduleAPIRequest {
	return &YunosTvpubadminContentAdvertQueryscheduleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.advert.queryschedule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Range Setter
// 查询范围: 1-牌照，4-uuid
func (r *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) SetRange(_range int64) error {
	r._range = _range
	r.Set("range", _range)
	return nil
}

// Get Range Getter
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetRange() int64 {
	return r._range
}

// Set is PageNo Setter
// 分页，页码
func (r *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 分页，单页数量
func (r *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is GmtStart Setter
// 日期
func (r *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) SetGmtStart(_gmtStart string) error {
	r._gmtStart = _gmtStart
	r.Set("gmt_start", _gmtStart)
	return nil
}

// Get GmtStart Getter
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetGmtStart() string {
	return r._gmtStart
}

// Set is Uuid Setter
// uuid
func (r *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetUuid() string {
	return r._uuid
}

// Set is License Setter
// 牌照方，1-华数，7-CIBN
func (r *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// Get License Getter
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetLicense() int64 {
	return r._license
}

// Set is SityType Setter
// 广告类型
func (r *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) SetSityType(_sityType int64) error {
	r._sityType = _sityType
	r.Set("sity_type", _sityType)
	return nil
}

// Get SityType Getter
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetSityType() int64 {
	return r._sityType
}
