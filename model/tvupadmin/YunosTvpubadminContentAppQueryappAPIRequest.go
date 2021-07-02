package tvupadmin

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentAppQueryappAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.app.queryapp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentAppQueryappAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
