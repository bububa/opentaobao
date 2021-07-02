package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimChatlogsGetAPIRequest openim聊天记录查询接口 API请求
// taobao.openim.chatlogs.get
//
// 查询openim账号聊天记录
type TaobaoOpenimChatlogsGetAPIRequest struct {
	model.Params
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
	// 迭代key
	_nextKey string
}

// NewTaobaoOpenimChatlogsGetRequest 初始化TaobaoOpenimChatlogsGetAPIRequest对象
func NewTaobaoOpenimChatlogsGetRequest() *TaobaoOpenimChatlogsGetAPIRequest {
	return &TaobaoOpenimChatlogsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimChatlogsGetAPIRequest) GetApiMethodName() string {
	return "taobao.openim.chatlogs.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimChatlogsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is User1 Setter
// 用户1信息
func (r *TaobaoOpenimChatlogsGetAPIRequest) SetUser1(_user1 *OpenImUser) error {
	r._user1 = _user1
	r.Set("user1", _user1)
	return nil
}

// Get User1 Getter
func (r TaobaoOpenimChatlogsGetAPIRequest) GetUser1() *OpenImUser {
	return r._user1
}

// Set is User2 Setter
// 用户2信息
func (r *TaobaoOpenimChatlogsGetAPIRequest) SetUser2(_user2 *OpenImUser) error {
	r._user2 = _user2
	r.Set("user2", _user2)
	return nil
}

// Get User2 Getter
func (r TaobaoOpenimChatlogsGetAPIRequest) GetUser2() *OpenImUser {
	return r._user2
}

// Set is Begin Setter
// 查询开始时间（UTC时间）
func (r *TaobaoOpenimChatlogsGetAPIRequest) SetBegin(_begin int64) error {
	r._begin = _begin
	r.Set("begin", _begin)
	return nil
}

// Get Begin Getter
func (r TaobaoOpenimChatlogsGetAPIRequest) GetBegin() int64 {
	return r._begin
}

// Set is End Setter
// 查询结束时间（UTC时间）
func (r *TaobaoOpenimChatlogsGetAPIRequest) SetEnd(_end int64) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// Get End Getter
func (r TaobaoOpenimChatlogsGetAPIRequest) GetEnd() int64 {
	return r._end
}

// Set is Count Setter
// 查询条数
func (r *TaobaoOpenimChatlogsGetAPIRequest) SetCount(_count int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// Get Count Getter
func (r TaobaoOpenimChatlogsGetAPIRequest) GetCount() int64 {
	return r._count
}

// Set is NextKey Setter
// 迭代key
func (r *TaobaoOpenimChatlogsGetAPIRequest) SetNextKey(_nextKey string) error {
	r._nextKey = _nextKey
	r.Set("next_key", _nextKey)
	return nil
}

// Get NextKey Getter
func (r TaobaoOpenimChatlogsGetAPIRequest) GetNextKey() string {
	return r._nextKey
}
