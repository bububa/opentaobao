package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentAppQuerybylicenceAPIRequest 按牌照查询应用 API请求
// yunos.tvpubadmin.content.app.querybylicence
//
// 按牌照查询应用
type YunosTvpubadminContentAppQuerybylicenceAPIRequest struct {
	model.Params
	// 查询条件
	_query string
}

// NewYunosTvpubadminContentAppQuerybylicenceRequest 初始化YunosTvpubadminContentAppQuerybylicenceAPIRequest对象
func NewYunosTvpubadminContentAppQuerybylicenceRequest() *YunosTvpubadminContentAppQuerybylicenceAPIRequest {
	return &YunosTvpubadminContentAppQuerybylicenceAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentAppQuerybylicenceAPIRequest) Reset() {
	r._query = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentAppQuerybylicenceAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.app.querybylicence"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentAppQuerybylicenceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentAppQuerybylicenceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询条件
func (r *YunosTvpubadminContentAppQuerybylicenceAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r YunosTvpubadminContentAppQuerybylicenceAPIRequest) GetQuery() string {
	return r._query
}

var poolYunosTvpubadminContentAppQuerybylicenceAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentAppQuerybylicenceRequest()
	},
}

// GetYunosTvpubadminContentAppQuerybylicenceRequest 从 sync.Pool 获取 YunosTvpubadminContentAppQuerybylicenceAPIRequest
func GetYunosTvpubadminContentAppQuerybylicenceAPIRequest() *YunosTvpubadminContentAppQuerybylicenceAPIRequest {
	return poolYunosTvpubadminContentAppQuerybylicenceAPIRequest.Get().(*YunosTvpubadminContentAppQuerybylicenceAPIRequest)
}

// ReleaseYunosTvpubadminContentAppQuerybylicenceAPIRequest 将 YunosTvpubadminContentAppQuerybylicenceAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentAppQuerybylicenceAPIRequest(v *YunosTvpubadminContentAppQuerybylicenceAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentAppQuerybylicenceAPIRequest.Put(v)
}
