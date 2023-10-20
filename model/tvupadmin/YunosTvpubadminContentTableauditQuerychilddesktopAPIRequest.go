package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontenttableauditquerychilddesktopAPIRequest 迎客松查看小酷宝桌面坑位元数据列表 API请求
// yunos.tvpubadmin.content.tableaudit.querychilddesktop
//
// 迎客松查看小酷宝桌面坑位元数据列表
type YunostvpubadmincontenttableauditquerychilddesktopAPIRequest struct {
	model.Params
	// 小酷宝桌面坑位查询参数
	_query string
}

// NewYunostvpubadmincontenttableauditquerychilddesktopRequest 初始化YunostvpubadmincontenttableauditquerychilddesktopAPIRequest对象
func NewYunostvpubadmincontenttableauditquerychilddesktopRequest() *YunostvpubadmincontenttableauditquerychilddesktopAPIRequest {
	return &YunostvpubadmincontenttableauditquerychilddesktopAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontenttableauditquerychilddesktopAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.tableaudit.querychilddesktop"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontenttableauditquerychilddesktopAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontenttableauditquerychilddesktopAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 小酷宝桌面坑位查询参数
func (r *YunostvpubadmincontenttableauditquerychilddesktopAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r YunostvpubadmincontenttableauditquerychilddesktopAPIRequest) GetQuery() string {
	return r._query
}
