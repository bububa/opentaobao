package tvupadmin

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminUserRightsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.user.rights"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminUserRightsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Uid Setter
// 用户ID
func (r *YunosTvpubadminUserRightsAPIRequest) SetUid(_uid int64) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// Get Uid Getter
func (r YunosTvpubadminUserRightsAPIRequest) GetUid() int64 {
	return r._uid
}

// Set is PageNo Setter
// 页码值
func (r *YunosTvpubadminUserRightsAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r YunosTvpubadminUserRightsAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 每页行数
func (r *YunosTvpubadminUserRightsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r YunosTvpubadminUserRightsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
