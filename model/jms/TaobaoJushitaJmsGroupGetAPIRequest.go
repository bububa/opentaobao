package jms

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJmsGroupGetAPIRequest 查询ONS分组 API请求
// taobao.jushita.jms.group.get
//
// 查询当前appkey在ONS中已有的分组
type TaobaoJushitaJmsGroupGetAPIRequest struct {
	model.Params
	// 要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。
	_groupNames []string
	// 页码
	_pageNo int64
	// 每页返回多少个分组
	_pageSize int64
}

// NewTaobaoJushitaJmsGroupGetRequest 初始化TaobaoJushitaJmsGroupGetAPIRequest对象
func NewTaobaoJushitaJmsGroupGetRequest() *TaobaoJushitaJmsGroupGetAPIRequest {
	return &TaobaoJushitaJmsGroupGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJushitaJmsGroupGetAPIRequest) Reset() {
	r._groupNames = r._groupNames[:0]
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJmsGroupGetAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jms.group.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJmsGroupGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJushitaJmsGroupGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupNames is GroupNames Setter
// 要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。
func (r *TaobaoJushitaJmsGroupGetAPIRequest) SetGroupNames(_groupNames []string) error {
	r._groupNames = _groupNames
	r.Set("group_names", _groupNames)
	return nil
}

// GetGroupNames GroupNames Getter
func (r TaobaoJushitaJmsGroupGetAPIRequest) GetGroupNames() []string {
	return r._groupNames
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoJushitaJmsGroupGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoJushitaJmsGroupGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页返回多少个分组
func (r *TaobaoJushitaJmsGroupGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoJushitaJmsGroupGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoJushitaJmsGroupGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJushitaJmsGroupGetRequest()
	},
}

// GetTaobaoJushitaJmsGroupGetRequest 从 sync.Pool 获取 TaobaoJushitaJmsGroupGetAPIRequest
func GetTaobaoJushitaJmsGroupGetAPIRequest() *TaobaoJushitaJmsGroupGetAPIRequest {
	return poolTaobaoJushitaJmsGroupGetAPIRequest.Get().(*TaobaoJushitaJmsGroupGetAPIRequest)
}

// ReleaseTaobaoJushitaJmsGroupGetAPIRequest 将 TaobaoJushitaJmsGroupGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoJushitaJmsGroupGetAPIRequest(v *TaobaoJushitaJmsGroupGetAPIRequest) {
	v.Reset()
	poolTaobaoJushitaJmsGroupGetAPIRequest.Put(v)
}
