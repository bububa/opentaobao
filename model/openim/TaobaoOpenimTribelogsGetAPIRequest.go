package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribelogsgetAPIRequest openim 群聊天记录导出接口 API请求
// taobao.openim.tribelogs.get
//
// 获取openim账号的群聊天记录
type TaobaoopenimtribelogsgetAPIRequest struct {
	model.Params
	// 群号
	_tribeId string
	// 迭代key
	_next string
	// 查询起始时间，UTC秒数。必须在一个月内。
	_begin int64
	// 查询结束时间，UTC秒数。必须大于起始时间并小于当前时间
	_end int64
	// 查询条数
	_count int64
}

// NewTaobaoopenimtribelogsgetRequest 初始化TaobaoopenimtribelogsgetAPIRequest对象
func NewTaobaoopenimtribelogsgetRequest() *TaobaoopenimtribelogsgetAPIRequest {
	return &TaobaoopenimtribelogsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimtribelogsgetAPIRequest) GetApiMethodName() string {
	return "taobao.openim.tribelogs.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimtribelogsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimtribelogsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTribeId is TribeId Setter
// 群号
func (r *TaobaoopenimtribelogsgetAPIRequest) SetTribeId(_tribeId string) error {
	r._tribeId = _tribeId
	r.Set("tribe_id", _tribeId)
	return nil
}

// GetTribeId TribeId Getter
func (r TaobaoopenimtribelogsgetAPIRequest) GetTribeId() string {
	return r._tribeId
}

// SetNext is Next Setter
// 迭代key
func (r *TaobaoopenimtribelogsgetAPIRequest) SetNext(_next string) error {
	r._next = _next
	r.Set("next", _next)
	return nil
}

// GetNext Next Getter
func (r TaobaoopenimtribelogsgetAPIRequest) GetNext() string {
	return r._next
}

// SetBegin is Begin Setter
// 查询起始时间，UTC秒数。必须在一个月内。
func (r *TaobaoopenimtribelogsgetAPIRequest) SetBegin(_begin int64) error {
	r._begin = _begin
	r.Set("begin", _begin)
	return nil
}

// GetBegin Begin Getter
func (r TaobaoopenimtribelogsgetAPIRequest) GetBegin() int64 {
	return r._begin
}

// SetEnd is End Setter
// 查询结束时间，UTC秒数。必须大于起始时间并小于当前时间
func (r *TaobaoopenimtribelogsgetAPIRequest) SetEnd(_end int64) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// GetEnd End Getter
func (r TaobaoopenimtribelogsgetAPIRequest) GetEnd() int64 {
	return r._end
}

// SetCount is Count Setter
// 查询条数
func (r *TaobaoopenimtribelogsgetAPIRequest) SetCount(_count int64) error {
	r._count = _count
	r.Set("count", _count)
	return nil
}

// GetCount Count Getter
func (r TaobaoopenimtribelogsgetAPIRequest) GetCount() int64 {
	return r._count
}
