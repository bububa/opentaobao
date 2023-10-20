package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentTableauditQuerymetaitemAPIRequest 运营位管控-查询魔盒运营位元数据列表 API请求
// yunos.tvpubadmin.content.tableaudit.querymetaitem
//
// 运营位管控-查询魔盒运营位元数据列表
type YunosTvpubadminContentTableauditQuerymetaitemAPIRequest struct {
	model.Params
	// 查询条件，json格式
	_tableAuditQueryBo string
}

// NewYunosTvpubadminContentTableauditQuerymetaitemRequest 初始化YunosTvpubadminContentTableauditQuerymetaitemAPIRequest对象
func NewYunosTvpubadminContentTableauditQuerymetaitemRequest() *YunosTvpubadminContentTableauditQuerymetaitemAPIRequest {
	return &YunosTvpubadminContentTableauditQuerymetaitemAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentTableauditQuerymetaitemAPIRequest) Reset() {
	r._tableAuditQueryBo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTableauditQuerymetaitemAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.tableaudit.querymetaitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTableauditQuerymetaitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentTableauditQuerymetaitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTableAuditQueryBo is TableAuditQueryBo Setter
// 查询条件，json格式
func (r *YunosTvpubadminContentTableauditQuerymetaitemAPIRequest) SetTableAuditQueryBo(_tableAuditQueryBo string) error {
	r._tableAuditQueryBo = _tableAuditQueryBo
	r.Set("table_audit_query_bo", _tableAuditQueryBo)
	return nil
}

// GetTableAuditQueryBo TableAuditQueryBo Getter
func (r YunosTvpubadminContentTableauditQuerymetaitemAPIRequest) GetTableAuditQueryBo() string {
	return r._tableAuditQueryBo
}

var poolYunosTvpubadminContentTableauditQuerymetaitemAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentTableauditQuerymetaitemRequest()
	},
}

// GetYunosTvpubadminContentTableauditQuerymetaitemRequest 从 sync.Pool 获取 YunosTvpubadminContentTableauditQuerymetaitemAPIRequest
func GetYunosTvpubadminContentTableauditQuerymetaitemAPIRequest() *YunosTvpubadminContentTableauditQuerymetaitemAPIRequest {
	return poolYunosTvpubadminContentTableauditQuerymetaitemAPIRequest.Get().(*YunosTvpubadminContentTableauditQuerymetaitemAPIRequest)
}

// ReleaseYunosTvpubadminContentTableauditQuerymetaitemAPIRequest 将 YunosTvpubadminContentTableauditQuerymetaitemAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentTableauditQuerymetaitemAPIRequest(v *YunosTvpubadminContentTableauditQuerymetaitemAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentTableauditQuerymetaitemAPIRequest.Put(v)
}
