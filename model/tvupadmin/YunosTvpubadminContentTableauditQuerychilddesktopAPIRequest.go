package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest 迎客松查看小酷宝桌面坑位元数据列表 API请求
// yunos.tvpubadmin.content.tableaudit.querychilddesktop
//
// 迎客松查看小酷宝桌面坑位元数据列表
type YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest struct {
	model.Params
	// 小酷宝桌面坑位查询参数
	_query string
}

// NewYunosTvpubadminContentTableauditQuerychilddesktopRequest 初始化YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest对象
func NewYunosTvpubadminContentTableauditQuerychilddesktopRequest() *YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest {
	return &YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.tableaudit.querychilddesktop"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Query Setter
// 小酷宝桌面坑位查询参数
func (r *YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// Get Query Getter
func (r YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest) GetQuery() string {
	return r._query
}
