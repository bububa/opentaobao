package tvupadmin

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest) Reset() {
	r._query = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.tableaudit.querychilddesktop"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 小酷宝桌面坑位查询参数
func (r *YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest) GetQuery() string {
	return r._query
}

var poolYunosTvpubadminContentTableauditQuerychilddesktopAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentTableauditQuerychilddesktopRequest()
	},
}

// GetYunosTvpubadminContentTableauditQuerychilddesktopRequest 从 sync.Pool 获取 YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest
func GetYunosTvpubadminContentTableauditQuerychilddesktopAPIRequest() *YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest {
	return poolYunosTvpubadminContentTableauditQuerychilddesktopAPIRequest.Get().(*YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest)
}

// ReleaseYunosTvpubadminContentTableauditQuerychilddesktopAPIRequest 将 YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentTableauditQuerychilddesktopAPIRequest(v *YunosTvpubadminContentTableauditQuerychilddesktopAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentTableauditQuerychilddesktopAPIRequest.Put(v)
}
