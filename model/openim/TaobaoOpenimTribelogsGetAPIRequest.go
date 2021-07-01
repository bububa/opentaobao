package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribelogsGetAPIRequest
openim 群聊天记录导出接口 API请求
taobao.openim.tribelogs.get

获取openim账号的群聊天记录 */
type TaobaoOpenimTribelogsGetAPIRequest struct {
	model.Params
	// 群号
	_tribeId string
	// 查询起始时间，UTC秒数。必须在一个月内。
	_begin int64
	// 查询结束时间，UTC秒数。必须大于起始时间并小于当前时间
	_end int64
	// 查询条数
	_count int64
	// 迭代key
	_next string
}

// NewTaobaoOpenimTribelogsGetRequest 初始化TaobaoOpenimTribelogsGetAPIRequest对象
func NewTaobaoOpenimTribelogsGetRequest() *TaobaoOpenimTribelogsGetAPIRequest {
	return &TaobaoOpenimTribelogsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribelogsGetAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribelogs.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribelogsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TribeId Setter
// 群号
func (r *TaobaoOpenimTribelogsGetAPIRequest) SetTribeId(_tribeId string) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// Get TribeId Getter
func (r TaobaoOpenimTribelogsGetAPIRequest) GetTribeId() string {
	return r._tribeId
}

// Set is Begin Setter
// 查询起始时间，UTC秒数。必须在一个月内。
func (r *TaobaoOpenimTribelogsGetAPIRequest) SetBegin(_begin int64) error {
	r._begin = _begin
	r.Set("begin", _begin)
	return nil
}

// Get Begin Getter
func (r TaobaoOpenimTribelogsGetAPIRequest) GetBegin() int64 {
	return r._begin
}

// Set is End Setter
// 查询结束时间，UTC秒数。必须大于起始时间并小于当前时间
func (r *TaobaoOpenimTribelogsGetAPIRequest) SetEnd(_end int64) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// Get End Getter
func (r TaobaoOpenimTribelogsGetAPIRequest) GetEnd() int64 {
	return r._end
}

// Set is Count Setter
// 查询条数
func (r *TaobaoOpenimTribelogsGetAPIRequest) SetCount(_count int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// Get Count Getter
func (r TaobaoOpenimTribelogsGetAPIRequest) GetCount() int64 {
	return r._count
}

// Set is Next Setter
// 迭代key
func (r *TaobaoOpenimTribelogsGetAPIRequest) SetNext(_next string) error {
	r._next = _next
	r.Set("next", _next)
	return nil
}

// Get Next Getter
func (r TaobaoOpenimTribelogsGetAPIRequest) GetNext() string {
	return r._next
}
