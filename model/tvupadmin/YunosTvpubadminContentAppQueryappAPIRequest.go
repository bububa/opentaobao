package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentAppQueryappAPIRequest 查询应用信息 API请求
// yunos.tvpubadmin.content.app.queryapp
//
// 查询应用信息
type YunosTvpubadminContentAppQueryappAPIRequest struct {
	model.Params
	// 查询条件
	_query string
}

// NewYunosTvpubadminContentAppQueryappRequest 初始化YunosTvpubadminContentAppQueryappAPIRequest对象
func NewYunosTvpubadminContentAppQueryappRequest() *YunosTvpubadminContentAppQueryappAPIRequest {
	return &YunosTvpubadminContentAppQueryappAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentAppQueryappAPIRequest) Reset() {
	r._query = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentAppQueryappAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.app.queryapp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentAppQueryappAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentAppQueryappAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询条件
func (r *YunosTvpubadminContentAppQueryappAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r YunosTvpubadminContentAppQueryappAPIRequest) GetQuery() string {
	return r._query
}

var poolYunosTvpubadminContentAppQueryappAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentAppQueryappRequest()
	},
}

// GetYunosTvpubadminContentAppQueryappRequest 从 sync.Pool 获取 YunosTvpubadminContentAppQueryappAPIRequest
func GetYunosTvpubadminContentAppQueryappAPIRequest() *YunosTvpubadminContentAppQueryappAPIRequest {
	return poolYunosTvpubadminContentAppQueryappAPIRequest.Get().(*YunosTvpubadminContentAppQueryappAPIRequest)
}

// ReleaseYunosTvpubadminContentAppQueryappAPIRequest 将 YunosTvpubadminContentAppQueryappAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentAppQueryappAPIRequest(v *YunosTvpubadminContentAppQueryappAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentAppQueryappAPIRequest.Put(v)
}
