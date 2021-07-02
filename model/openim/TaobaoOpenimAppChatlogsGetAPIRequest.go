package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimAppChatlogsGetAPIRequest openim应用聊天记录查询 API请求
// taobao.openim.app.chatlogs.get
//
// 查询openim应用的聊天记录
type TaobaoOpenimAppChatlogsGetAPIRequest struct {
	model.Params
	// 查询结束时间。UTC时间。精度到秒
	_beg int64
	// 查询结束时间。UTC时间。精度到秒
	_end int64
	// 查询最大条数
	_count int64
	// 迭代key
	_next string
}

// NewTaobaoOpenimAppChatlogsGetRequest 初始化TaobaoOpenimAppChatlogsGetAPIRequest对象
func NewTaobaoOpenimAppChatlogsGetRequest() *TaobaoOpenimAppChatlogsGetAPIRequest {
	return &TaobaoOpenimAppChatlogsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimAppChatlogsGetAPIRequest) GetApiMethodName() string {
	return "taobao.openim.app.chatlogs.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimAppChatlogsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Beg Setter
// 查询结束时间。UTC时间。精度到秒
func (r *TaobaoOpenimAppChatlogsGetAPIRequest) SetBeg(_beg int64) error {
	r._beg = _beg
	r.Set("beg", _beg)
	return nil
}

// Get Beg Getter
func (r TaobaoOpenimAppChatlogsGetAPIRequest) GetBeg() int64 {
	return r._beg
}

// Set is End Setter
// 查询结束时间。UTC时间。精度到秒
func (r *TaobaoOpenimAppChatlogsGetAPIRequest) SetEnd(_end int64) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// Get End Getter
func (r TaobaoOpenimAppChatlogsGetAPIRequest) GetEnd() int64 {
	return r._end
}

// Set is Count Setter
// 查询最大条数
func (r *TaobaoOpenimAppChatlogsGetAPIRequest) SetCount(_count int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// Get Count Getter
func (r TaobaoOpenimAppChatlogsGetAPIRequest) GetCount() int64 {
	return r._count
}

// Set is Next Setter
// 迭代key
func (r *TaobaoOpenimAppChatlogsGetAPIRequest) SetNext(_next string) error {
	r._next = _next
	r.Set("next", _next)
	return nil
}

// Get Next Getter
func (r TaobaoOpenimAppChatlogsGetAPIRequest) GetNext() string {
	return r._next
}
