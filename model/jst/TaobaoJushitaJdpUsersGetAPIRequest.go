package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJdpUsersGetAPIRequest 获取开通的订单同步服务的用户 API请求
// taobao.jushita.jdp.users.get
//
// 获取开通的订单同步服务的用户，含有rds的路由关系
type TaobaoJushitaJdpUsersGetAPIRequest struct {
	model.Params
	// 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
	_startModified string
	// 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
	_endModified string
	// 每页记录数，默认200
	_pageSize int64
	// 当前页数
	_pageNo int64
	// 如果传了user_id表示单条查询
	_userId int64
}

// NewTaobaoJushitaJdpUsersGetRequest 初始化TaobaoJushitaJdpUsersGetAPIRequest对象
func NewTaobaoJushitaJdpUsersGetRequest() *TaobaoJushitaJdpUsersGetAPIRequest {
	return &TaobaoJushitaJdpUsersGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJdpUsersGetAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jdp.users.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJdpUsersGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJushitaJdpUsersGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartModified is StartModified Setter
// 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
func (r *TaobaoJushitaJdpUsersGetAPIRequest) SetStartModified(_startModified string) error {
	r._startModified = _startModified
	r.Set("start_modified", _startModified)
	return nil
}

// GetStartModified StartModified Getter
func (r TaobaoJushitaJdpUsersGetAPIRequest) GetStartModified() string {
	return r._startModified
}

// SetEndModified is EndModified Setter
// 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
func (r *TaobaoJushitaJdpUsersGetAPIRequest) SetEndModified(_endModified string) error {
	r._endModified = _endModified
	r.Set("end_modified", _endModified)
	return nil
}

// GetEndModified EndModified Getter
func (r TaobaoJushitaJdpUsersGetAPIRequest) GetEndModified() string {
	return r._endModified
}

// SetPageSize is PageSize Setter
// 每页记录数，默认200
func (r *TaobaoJushitaJdpUsersGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoJushitaJdpUsersGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 当前页数
func (r *TaobaoJushitaJdpUsersGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoJushitaJdpUsersGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetUserId is UserId Setter
// 如果传了user_id表示单条查询
func (r *TaobaoJushitaJdpUsersGetAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoJushitaJdpUsersGetAPIRequest) GetUserId() int64 {
	return r._userId
}
