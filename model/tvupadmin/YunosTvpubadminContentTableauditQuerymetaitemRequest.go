package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
运营位管控-查询魔盒运营位元数据列表 API请求
yunos.tvpubadmin.content.tableaudit.querymetaitem

运营位管控-查询魔盒运营位元数据列表
*/
type YunosTvpubadminContentTableauditQuerymetaitemRequest struct {
    model.Params
    // 查询条件，json格式
    _tableAuditQueryBo   string
}

// 初始化YunosTvpubadminContentTableauditQuerymetaitemRequest对象
func NewYunosTvpubadminContentTableauditQuerymetaitemRequest() *YunosTvpubadminContentTableauditQuerymetaitemRequest{
    return &YunosTvpubadminContentTableauditQuerymetaitemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTableauditQuerymetaitemRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.tableaudit.querymetaitem"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTableauditQuerymetaitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TableAuditQueryBo Setter
// 查询条件，json格式
func (r *YunosTvpubadminContentTableauditQuerymetaitemRequest) SetTableAuditQueryBo(_tableAuditQueryBo string) error {
    r._tableAuditQueryBo = _tableAuditQueryBo
    r.Set("table_audit_query_bo", _tableAuditQueryBo)
    return nil
}

// TableAuditQueryBo Getter
func (r YunosTvpubadminContentTableauditQuerymetaitemRequest) GetTableAuditQueryBo() string {
    return r._tableAuditQueryBo
}
