package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentappqueryappAPIRequest 查询应用信息 API请求
// yunos.tvpubadmin.content.app.queryapp
//
// 查询应用信息
type YunostvpubadmincontentappqueryappAPIRequest struct {
	model.Params
	// 查询条件
	_query string
}

// NewYunostvpubadmincontentappqueryappRequest 初始化YunostvpubadmincontentappqueryappAPIRequest对象
func NewYunostvpubadmincontentappqueryappRequest() *YunostvpubadmincontentappqueryappAPIRequest {
	return &YunostvpubadmincontentappqueryappAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentappqueryappAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.app.queryapp"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentappqueryappAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentappqueryappAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询条件
func (r *YunostvpubadmincontentappqueryappAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r YunostvpubadmincontentappqueryappAPIRequest) GetQuery() string {
	return r._query
}
