package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimchatlogsgetAPIRequest openim聊天记录查询接口 API请求
// taobao.openim.chatlogs.get
//
// 查询openim账号聊天记录
type TaobaoopenimchatlogsgetAPIRequest struct {
	model.Params
	// 迭代key
	_nextKey string
	// 用户1信息
	_user1 *OpenImUser
	// 用户2信息
	_user2 *OpenImUser
	// 查询开始时间（UTC时间）
	_begin int64
	// 查询结束时间（UTC时间）
	_end int64
	// 查询条数
	_count int64
}

// NewTaobaoopenimchatlogsgetRequest 初始化TaobaoopenimchatlogsgetAPIRequest对象
func NewTaobaoopenimchatlogsgetRequest() *TaobaoopenimchatlogsgetAPIRequest {
	return &TaobaoopenimchatlogsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimchatlogsgetAPIRequest) GetApiMethodName() string {
	return "taobao.openim.chatlogs.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimchatlogsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimchatlogsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNextKey is NextKey Setter
// 迭代key
func (r *TaobaoopenimchatlogsgetAPIRequest) SetNextKey(_nextKey string) error {
	r._nextKey = _nextKey
	r.Set("next_key", _nextKey)
	return nil
}

// GetNextKey NextKey Getter
func (r TaobaoopenimchatlogsgetAPIRequest) GetNextKey() string {
	return r._nextKey
}

// SetUser1 is User1 Setter
// 用户1信息
func (r *TaobaoopenimchatlogsgetAPIRequest) SetUser1(_user1 *OpenImUser) error {
	r._user1 = _user1
	r.Set("user1", _user1)
	return nil
}

// GetUser1 User1 Getter
func (r TaobaoopenimchatlogsgetAPIRequest) GetUser1() *OpenImUser {
	return r._user1
}

// SetUser2 is User2 Setter
// 用户2信息
func (r *TaobaoopenimchatlogsgetAPIRequest) SetUser2(_user2 *OpenImUser) error {
	r._user2 = _user2
	r.Set("user2", _user2)
	return nil
}

// GetUser2 User2 Getter
func (r TaobaoopenimchatlogsgetAPIRequest) GetUser2() *OpenImUser {
	return r._user2
}

// SetBegin is Begin Setter
// 查询开始时间（UTC时间）
func (r *TaobaoopenimchatlogsgetAPIRequest) SetBegin(_begin int64) error {
	r._begin = _begin
	r.Set("begin", _begin)
	return nil
}

// GetBegin Begin Getter
func (r TaobaoopenimchatlogsgetAPIRequest) GetBegin() int64 {
	return r._begin
}

// SetEnd is End Setter
// 查询结束时间（UTC时间）
func (r *TaobaoopenimchatlogsgetAPIRequest) SetEnd(_end int64) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// GetEnd End Getter
func (r TaobaoopenimchatlogsgetAPIRequest) GetEnd() int64 {
	return r._end
}

// SetCount is Count Setter
// 查询条数
func (r *TaobaoopenimchatlogsgetAPIRequest) SetCount(_count int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// GetCount Count Getter
func (r TaobaoopenimchatlogsgetAPIRequest) GetCount() int64 {
	return r._count
}
