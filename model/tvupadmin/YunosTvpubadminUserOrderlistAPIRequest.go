package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminuserorderlistAPIRequest 获取用户订单列表 API请求
// yunos.tvpubadmin.user.orderlist
//
// 获取用户订单列表
type YunostvpubadminuserorderlistAPIRequest struct {
	model.Params
	// 开始时间
	_createTimeStartStr string
	// 结束时间
	_createTimeEndStr string
	// 用户ID
	_uid int64
	// 牌照方
	_license int64
	// 页码值
	_pageNo int64
	// 每页行数
	_pageSize int64
}

// NewYunostvpubadminuserorderlistRequest 初始化YunostvpubadminuserorderlistAPIRequest对象
func NewYunostvpubadminuserorderlistRequest() *YunostvpubadminuserorderlistAPIRequest {
	return &YunostvpubadminuserorderlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminuserorderlistAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.user.orderlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminuserorderlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminuserorderlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateTimeStartStr is CreateTimeStartStr Setter
// 开始时间
func (r *YunostvpubadminuserorderlistAPIRequest) SetCreateTimeStartStr(_createTimeStartStr string) error {
	r._createTimeStartStr = _createTimeStartStr
	r.Set("create_time_start_str", _createTimeStartStr)
	return nil
}

// GetCreateTimeStartStr CreateTimeStartStr Getter
func (r YunostvpubadminuserorderlistAPIRequest) GetCreateTimeStartStr() string {
	return r._createTimeStartStr
}

// SetCreateTimeEndStr is CreateTimeEndStr Setter
// 结束时间
func (r *YunostvpubadminuserorderlistAPIRequest) SetCreateTimeEndStr(_createTimeEndStr string) error {
	r._createTimeEndStr = _createTimeEndStr
	r.Set("create_time_end_str", _createTimeEndStr)
	return nil
}

// GetCreateTimeEndStr CreateTimeEndStr Getter
func (r YunostvpubadminuserorderlistAPIRequest) GetCreateTimeEndStr() string {
	return r._createTimeEndStr
}

// SetUid is Uid Setter
// 用户ID
func (r *YunostvpubadminuserorderlistAPIRequest) SetUid(_uid int64) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// GetUid Uid Getter
func (r YunostvpubadminuserorderlistAPIRequest) GetUid() int64 {
	return r._uid
}

// SetLicense is License Setter
// 牌照方
func (r *YunostvpubadminuserorderlistAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadminuserorderlistAPIRequest) GetLicense() int64 {
	return r._license
}

// SetPageNo is PageNo Setter
// 页码值
func (r *YunostvpubadminuserorderlistAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunostvpubadminuserorderlistAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页行数
func (r *YunostvpubadminuserorderlistAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunostvpubadminuserorderlistAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
