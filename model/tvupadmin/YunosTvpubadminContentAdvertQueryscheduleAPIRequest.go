package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentadvertqueryscheduleAPIRequest 广告牌照管控查询 API请求
// yunos.tvpubadmin.content.advert.queryschedule
//
// 广告牌照管控查询
type YunostvpubadmincontentadvertqueryscheduleAPIRequest struct {
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

// NewYunostvpubadmincontentadvertqueryscheduleRequest 初始化YunostvpubadmincontentadvertqueryscheduleAPIRequest对象
func NewYunostvpubadmincontentadvertqueryscheduleRequest() *YunostvpubadmincontentadvertqueryscheduleAPIRequest {
	return &YunostvpubadmincontentadvertqueryscheduleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentadvertqueryscheduleAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.advert.queryschedule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentadvertqueryscheduleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentadvertqueryscheduleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGmtStart is GmtStart Setter
// 日期
func (r *YunostvpubadmincontentadvertqueryscheduleAPIRequest) SetGmtStart(_gmtStart string) error {
	r._gmtStart = _gmtStart
	r.Set("gmt_start", _gmtStart)
	return nil
}

// GetGmtStart GmtStart Getter
func (r YunostvpubadmincontentadvertqueryscheduleAPIRequest) GetGmtStart() string {
	return r._gmtStart
}

// SetUuid is Uuid Setter
// uuid
func (r *YunostvpubadmincontentadvertqueryscheduleAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r YunostvpubadmincontentadvertqueryscheduleAPIRequest) GetUuid() string {
	return r._uuid
}

// SetRange is Range Setter
// 查询范围: 1-牌照，4-uuid
func (r *YunostvpubadmincontentadvertqueryscheduleAPIRequest) SetRange(_range int64) error {
	r._range = _range
	r.Set("range", _range)
	return nil
}

// GetRange Range Getter
func (r YunostvpubadmincontentadvertqueryscheduleAPIRequest) GetRange() int64 {
	return r._range
}

// SetPageNo is PageNo Setter
// 分页，页码
func (r *YunostvpubadmincontentadvertqueryscheduleAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunostvpubadmincontentadvertqueryscheduleAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页，单页数量
func (r *YunostvpubadmincontentadvertqueryscheduleAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunostvpubadmincontentadvertqueryscheduleAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetLicense is License Setter
// 牌照方，1-华数，7-CIBN
func (r *YunostvpubadmincontentadvertqueryscheduleAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmincontentadvertqueryscheduleAPIRequest) GetLicense() int64 {
	return r._license
}

// SetSityType is SityType Setter
// 广告类型
func (r *YunostvpubadmincontentadvertqueryscheduleAPIRequest) SetSityType(_sityType int64) error {
	r._sityType = _sityType
	r.Set("sity_type", _sityType)
	return nil
}

// GetSityType SityType Getter
func (r YunostvpubadmincontentadvertqueryscheduleAPIRequest) GetSityType() int64 {
	return r._sityType
}
