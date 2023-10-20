package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageDialogListAPIRequest 分页获取弹窗列表 API请求
// yunos.tvpubadmin.manage.dialog.list
//
// 分页获取弹窗配置列表
type YunosTvpubadminManageDialogListAPIRequest struct {
	model.Params
	// 查询的query参数
	_query string
}

// NewYunosTvpubadminManageDialogListRequest 初始化YunosTvpubadminManageDialogListAPIRequest对象
func NewYunosTvpubadminManageDialogListRequest() *YunosTvpubadminManageDialogListAPIRequest {
	return &YunosTvpubadminManageDialogListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageDialogListAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.dialog.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageDialogListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminManageDialogListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询的query参数
func (r *YunosTvpubadminManageDialogListAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r YunosTvpubadminManageDialogListAPIRequest) GetQuery() string {
	return r._query
}
