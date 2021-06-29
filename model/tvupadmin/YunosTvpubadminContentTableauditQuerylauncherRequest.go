package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
运营位管控-查询联盟一体机运营位元数据列表 APIRequest
yunos.tvpubadmin.content.tableaudit.querylauncher

运营位管控-查询联盟一体机运营位元数据列表
*/
type YunosTvpubadminContentTableauditQuerylauncherRequest struct {
    model.Params

    // 桌面坑位审核查询条件,json格式
    tableAuditQuery   string 

}

func NewYunosTvpubadminContentTableauditQuerylauncherRequest() *YunosTvpubadminContentTableauditQuerylauncherRequest{
    return &YunosTvpubadminContentTableauditQuerylauncherRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentTableauditQuerylauncherRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.tableaudit.querylauncher"
}

func (r YunosTvpubadminContentTableauditQuerylauncherRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentTableauditQuerylauncherRequest) SetTableAuditQuery(tableAuditQuery string) error {
    r.tableAuditQuery = tableAuditQuery
    r.Set("table_audit_query", tableAuditQuery)
    return nil
}

func (r YunosTvpubadminContentTableauditQuerylauncherRequest) GetTableAuditQuery() string {
    return r.tableAuditQuery
}

