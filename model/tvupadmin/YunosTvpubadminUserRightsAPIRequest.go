package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminUserRightsAPIRequest 获取用户权益 API请求
// yunos.tvpubadmin.user.rights
//
// 获取用户权益
type YunosTvpubadminUserRightsAPIRequest struct {
	model.Params
	// 用户ID
	_uid int64
	// 页码值
	_pageNo int64
	// 每页行数
	_pageSize int64
}

// NewYunosTvpubadminUserRightsRequest 初始化YunosTvpubadminUserRightsAPIRequest对象
func NewYunosTvpubadminUserRightsRequest() *YunosTvpubadminUserRightsAPIRequest {
	return &YunosTvpubadminUserRightsAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminUserRightsAPIRequest) Reset() {
	r._uid = 0
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminUserRightsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.user.rights"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminUserRightsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminUserRightsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUid is Uid Setter
// 用户ID
func (r *YunosTvpubadminUserRightsAPIRequest) SetUid(_uid int64) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// GetUid Uid Getter
func (r YunosTvpubadminUserRightsAPIRequest) GetUid() int64 {
	return r._uid
}

// SetPageNo is PageNo Setter
// 页码值
func (r *YunosTvpubadminUserRightsAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunosTvpubadminUserRightsAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页行数
func (r *YunosTvpubadminUserRightsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunosTvpubadminUserRightsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolYunosTvpubadminUserRightsAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminUserRightsRequest()
	},
}

// GetYunosTvpubadminUserRightsRequest 从 sync.Pool 获取 YunosTvpubadminUserRightsAPIRequest
func GetYunosTvpubadminUserRightsAPIRequest() *YunosTvpubadminUserRightsAPIRequest {
	return poolYunosTvpubadminUserRightsAPIRequest.Get().(*YunosTvpubadminUserRightsAPIRequest)
}

// ReleaseYunosTvpubadminUserRightsAPIRequest 将 YunosTvpubadminUserRightsAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminUserRightsAPIRequest(v *YunosTvpubadminUserRightsAPIRequest) {
	v.Reset()
	poolYunosTvpubadminUserRightsAPIRequest.Put(v)
}
