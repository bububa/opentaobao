package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentTableauditQuerymetaitemAPIRequest
运营位管控-查询魔盒运营位元数据列表 API请求
yunos.tvpubadmin.content.tableaudit.querymetaitem

运营位管控-查询魔盒运营位元数据列表 */
type YunosTvpubadminContentTableauditQuerymetaitemAPIRequest struct {
	model.Params
	// 查询条件，json格式
	_tableAuditQueryBo string
}

// NewYunosTvpubadminContentTableauditQuerymetaitemRequest 初始化YunosTvpubadminContentTableauditQuerymetaitemAPIRequest对象
func NewYunosTvpubadminContentTableauditQuerymetaitemRequest() *YunosTvpubadminContentTableauditQuerymetaitemAPIRequest {
	return &YunosTvpubadminContentTableauditQuerymetaitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTableauditQuerymetaitemAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.tableaudit.querymetaitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTableauditQuerymetaitemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TableAuditQueryBo Setter
// 查询条件，json格式
func (r *YunosTvpubadminContentTableauditQuerymetaitemAPIRequest) SetTableAuditQueryBo(_tableAuditQueryBo string) error {
	r._tableAuditQueryBo = _tableAuditQueryBo
	r.Set("table_audit_query_bo", _tableAuditQueryBo)
	return nil
}

// Get TableAuditQueryBo Getter
func (r YunosTvpubadminContentTableauditQuerymetaitemAPIRequest) GetTableAuditQueryBo() string {
	return r._tableAuditQueryBo
}
