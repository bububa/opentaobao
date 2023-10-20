package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontenttableauditquerymetaitemAPIRequest 运营位管控-查询魔盒运营位元数据列表 API请求
// yunos.tvpubadmin.content.tableaudit.querymetaitem
//
// 运营位管控-查询魔盒运营位元数据列表
type YunostvpubadmincontenttableauditquerymetaitemAPIRequest struct {
	model.Params
	// 查询条件，json格式
	_tableAuditQueryBo string
}

// NewYunostvpubadmincontenttableauditquerymetaitemRequest 初始化YunostvpubadmincontenttableauditquerymetaitemAPIRequest对象
func NewYunostvpubadmincontenttableauditquerymetaitemRequest() *YunostvpubadmincontenttableauditquerymetaitemAPIRequest {
	return &YunostvpubadmincontenttableauditquerymetaitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontenttableauditquerymetaitemAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.tableaudit.querymetaitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontenttableauditquerymetaitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontenttableauditquerymetaitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTableAuditQueryBo is TableAuditQueryBo Setter
// 查询条件，json格式
func (r *YunostvpubadmincontenttableauditquerymetaitemAPIRequest) SetTableAuditQueryBo(_tableAuditQueryBo string) error {
	r._tableAuditQueryBo = _tableAuditQueryBo
	r.Set("table_audit_query_bo", _tableAuditQueryBo)
	return nil
}

// GetTableAuditQueryBo TableAuditQueryBo Getter
func (r YunostvpubadmincontenttableauditquerymetaitemAPIRequest) GetTableAuditQueryBo() string {
	return r._tableAuditQueryBo
}
