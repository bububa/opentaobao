package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagetopiclistAPIRequest 专题内容操作列表 API请求
// yunos.tvpubadmin.manage.topic.list
//
// 获取外部可操作编辑的专题列表
type YunostvpubadminmanagetopiclistAPIRequest struct {
	model.Params
	// 查询条件
	_query string
}

// NewYunostvpubadminmanagetopiclistRequest 初始化YunostvpubadminmanagetopiclistAPIRequest对象
func NewYunostvpubadminmanagetopiclistRequest() *YunostvpubadminmanagetopiclistAPIRequest {
	return &YunostvpubadminmanagetopiclistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminmanagetopiclistAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminmanagetopiclistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminmanagetopiclistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询条件
func (r *YunostvpubadminmanagetopiclistAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r YunostvpubadminmanagetopiclistAPIRequest) GetQuery() string {
	return r._query
}
