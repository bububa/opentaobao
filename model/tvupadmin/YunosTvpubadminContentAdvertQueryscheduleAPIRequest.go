package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentAdvertQueryscheduleAPIRequest 广告牌照管控查询 API请求
// yunos.tvpubadmin.content.advert.queryschedule
//
// 广告牌照管控查询
type YunosTvpubadminContentAdvertQueryscheduleAPIRequest struct {
	model.Params
	// 日期
	_gmtStart string
	// uuid
	_uuid string
	// 查询范围: 1-牌照，4-uuid
	_range int64
	// 分页，页码
	_pageNo int64
	// 分页，单页数量
	_pageSize int64
	// 牌照方，1-华数，7-CIBN
	_license int64
	// 广告类型
	_sityType int64
}

// NewYunosTvpubadminContentAdvertQueryscheduleRequest 初始化YunosTvpubadminContentAdvertQueryscheduleAPIRequest对象
func NewYunosTvpubadminContentAdvertQueryscheduleRequest() *YunosTvpubadminContentAdvertQueryscheduleAPIRequest {
	return &YunosTvpubadminContentAdvertQueryscheduleAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) Reset() {
	r._gmtStart = ""
	r._uuid = ""
	r._range = 0
	r._pageNo = 0
	r._pageSize = 0
	r._license = 0
	r._sityType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.advert.queryschedule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGmtStart is GmtStart Setter
// 日期
func (r *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) SetGmtStart(_gmtStart string) error {
	r._gmtStart = _gmtStart
	r.Set("gmt_start", _gmtStart)
	return nil
}

// GetGmtStart GmtStart Getter
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetGmtStart() string {
	return r._gmtStart
}

// SetUuid is Uuid Setter
// uuid
func (r *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetUuid() string {
	return r._uuid
}

// SetRange is Range Setter
// 查询范围: 1-牌照，4-uuid
func (r *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) SetRange(_range int64) error {
	r._range = _range
	r.Set("range", _range)
	return nil
}

// GetRange Range Getter
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetRange() int64 {
	return r._range
}

// SetPageNo is PageNo Setter
// 分页，页码
func (r *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页，单页数量
func (r *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetLicense is License Setter
// 牌照方，1-华数，7-CIBN
func (r *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetLicense() int64 {
	return r._license
}

// SetSityType is SityType Setter
// 广告类型
func (r *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) SetSityType(_sityType int64) error {
	r._sityType = _sityType
	r.Set("sity_type", _sityType)
	return nil
}

// GetSityType SityType Getter
func (r YunosTvpubadminContentAdvertQueryscheduleAPIRequest) GetSityType() int64 {
	return r._sityType
}

var poolYunosTvpubadminContentAdvertQueryscheduleAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentAdvertQueryscheduleRequest()
	},
}

// GetYunosTvpubadminContentAdvertQueryscheduleRequest 从 sync.Pool 获取 YunosTvpubadminContentAdvertQueryscheduleAPIRequest
func GetYunosTvpubadminContentAdvertQueryscheduleAPIRequest() *YunosTvpubadminContentAdvertQueryscheduleAPIRequest {
	return poolYunosTvpubadminContentAdvertQueryscheduleAPIRequest.Get().(*YunosTvpubadminContentAdvertQueryscheduleAPIRequest)
}

// ReleaseYunosTvpubadminContentAdvertQueryscheduleAPIRequest 将 YunosTvpubadminContentAdvertQueryscheduleAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentAdvertQueryscheduleAPIRequest(v *YunosTvpubadminContentAdvertQueryscheduleAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentAdvertQueryscheduleAPIRequest.Put(v)
}
