package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminuserrightsAPIRequest 获取用户权益 API请求
// yunos.tvpubadmin.user.rights
//
// 获取用户权益
type YunostvpubadminuserrightsAPIRequest struct {
	model.Params
	// 用户ID
	_uid int64
	// 页码值
	_pageNo int64
	// 每页行数
	_pageSize int64
}

// NewYunostvpubadminuserrightsRequest 初始化YunostvpubadminuserrightsAPIRequest对象
func NewYunostvpubadminuserrightsRequest() *YunostvpubadminuserrightsAPIRequest {
	return &YunostvpubadminuserrightsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminuserrightsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.user.rights"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminuserrightsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminuserrightsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUid is Uid Setter
// 用户ID
func (r *YunostvpubadminuserrightsAPIRequest) SetUid(_uid int64) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// GetUid Uid Getter
func (r YunostvpubadminuserrightsAPIRequest) GetUid() int64 {
	return r._uid
}

// SetPageNo is PageNo Setter
// 页码值
func (r *YunostvpubadminuserrightsAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunostvpubadminuserrightsAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页行数
func (r *YunostvpubadminuserrightsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunostvpubadminuserrightsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
