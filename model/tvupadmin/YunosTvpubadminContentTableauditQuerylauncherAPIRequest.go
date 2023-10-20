package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentTableauditQuerylauncherAPIRequest 运营位管控-查询联盟一体机运营位元数据列表 API请求
// yunos.tvpubadmin.content.tableaudit.querylauncher
//
// 运营位管控-查询联盟一体机运营位元数据列表
type YunosTvpubadminContentTableauditQuerylauncherAPIRequest struct {
	model.Params
	// 桌面坑位审核查询条件,json格式
	_tableAuditQuery string
}

// NewYunosTvpubadminContentTableauditQuerylauncherRequest 初始化YunosTvpubadminContentTableauditQuerylauncherAPIRequest对象
func NewYunosTvpubadminContentTableauditQuerylauncherRequest() *YunosTvpubadminContentTableauditQuerylauncherAPIRequest {
	return &YunosTvpubadminContentTableauditQuerylauncherAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentTableauditQuerylauncherAPIRequest) Reset() {
	r._tableAuditQuery = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTableauditQuerylauncherAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.tableaudit.querylauncher"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTableauditQuerylauncherAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentTableauditQuerylauncherAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTableAuditQuery is TableAuditQuery Setter
// 桌面坑位审核查询条件,json格式
func (r *YunosTvpubadminContentTableauditQuerylauncherAPIRequest) SetTableAuditQuery(_tableAuditQuery string) error {
	r._tableAuditQuery = _tableAuditQuery
	r.Set("table_audit_query", _tableAuditQuery)
	return nil
}

// GetTableAuditQuery TableAuditQuery Getter
func (r YunosTvpubadminContentTableauditQuerylauncherAPIRequest) GetTableAuditQuery() string {
	return r._tableAuditQuery
}

var poolYunosTvpubadminContentTableauditQuerylauncherAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentTableauditQuerylauncherRequest()
	},
}

// GetYunosTvpubadminContentTableauditQuerylauncherRequest 从 sync.Pool 获取 YunosTvpubadminContentTableauditQuerylauncherAPIRequest
func GetYunosTvpubadminContentTableauditQuerylauncherAPIRequest() *YunosTvpubadminContentTableauditQuerylauncherAPIRequest {
	return poolYunosTvpubadminContentTableauditQuerylauncherAPIRequest.Get().(*YunosTvpubadminContentTableauditQuerylauncherAPIRequest)
}

// ReleaseYunosTvpubadminContentTableauditQuerylauncherAPIRequest 将 YunosTvpubadminContentTableauditQuerylauncherAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentTableauditQuerylauncherAPIRequest(v *YunosTvpubadminContentTableauditQuerylauncherAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentTableauditQuerylauncherAPIRequest.Put(v)
}
