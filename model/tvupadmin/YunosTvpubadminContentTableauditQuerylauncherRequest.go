package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
运营位管控-查询联盟一体机运营位元数据列表 API请求
yunos.tvpubadmin.content.tableaudit.querylauncher

运营位管控-查询联盟一体机运营位元数据列表
*/
type YunosTvpubadminContentTableauditQuerylauncherRequest struct {
    model.Params
    // 桌面坑位审核查询条件,json格式
    tableAuditQuery   string
}

// 初始化YunosTvpubadminContentTableauditQuerylauncherRequest对象
func NewYunosTvpubadminContentTableauditQuerylauncherRequest() *YunosTvpubadminContentTableauditQuerylauncherRequest{
    return &YunosTvpubadminContentTableauditQuerylauncherRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentTableauditQuerylauncherRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.tableaudit.querylauncher"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentTableauditQuerylauncherRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TableAuditQuery Setter
// 桌面坑位审核查询条件,json格式
func (r *YunosTvpubadminContentTableauditQuerylauncherRequest) SetTableAuditQuery(tableAuditQuery string) error {
    r.tableAuditQuery = tableAuditQuery
    r.Set("table_audit_query", tableAuditQuery)
    return nil
}

// TableAuditQuery Getter
func (r YunosTvpubadminContentTableauditQuerylauncherRequest) GetTableAuditQuery() string {
    return r.tableAuditQuery
}
