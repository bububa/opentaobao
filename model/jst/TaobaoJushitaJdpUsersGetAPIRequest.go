package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojushitajdpusersgetAPIRequest 获取开通的订单同步服务的用户 API请求
// taobao.jushita.jdp.users.get
//
// 获取开通的订单同步服务的用户，含有rds的路由关系
type TaobaojushitajdpusersgetAPIRequest struct {
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

// NewTaobaojushitajdpusersgetRequest 初始化TaobaojushitajdpusersgetAPIRequest对象
func NewTaobaojushitajdpusersgetRequest() *TaobaojushitajdpusersgetAPIRequest {
	return &TaobaojushitajdpusersgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojushitajdpusersgetAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jdp.users.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojushitajdpusersgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojushitajdpusersgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartModified is StartModified Setter
// 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
func (r *TaobaojushitajdpusersgetAPIRequest) SetStartModified(_startModified string) error {
	r._startModified = _startModified
	r.Set("start_modified", _startModified)
	return nil
}

// GetStartModified StartModified Getter
func (r TaobaojushitajdpusersgetAPIRequest) GetStartModified() string {
	return r._startModified
}

// SetEndModified is EndModified Setter
// 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户
func (r *TaobaojushitajdpusersgetAPIRequest) SetEndModified(_endModified string) error {
	r._endModified = _endModified
	r.Set("end_modified", _endModified)
	return nil
}

// GetEndModified EndModified Getter
func (r TaobaojushitajdpusersgetAPIRequest) GetEndModified() string {
	return r._endModified
}

// SetPageSize is PageSize Setter
// 每页记录数，默认200
func (r *TaobaojushitajdpusersgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaojushitajdpusersgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 当前页数
func (r *TaobaojushitajdpusersgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaojushitajdpusersgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetUserId is UserId Setter
// 如果传了user_id表示单条查询
func (r *TaobaojushitajdpusersgetAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaojushitajdpusersgetAPIRequest) GetUserId() int64 {
	return r._userId
}
