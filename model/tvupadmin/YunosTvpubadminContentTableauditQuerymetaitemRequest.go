package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
运营位管控-查询魔盒运营位元数据列表 APIRequest
yunos.tvpubadmin.content.tableaudit.querymetaitem

运营位管控-查询魔盒运营位元数据列表
*/
type YunosTvpubadminContentTableauditQuerymetaitemRequest struct {
    model.Params

    // 查询条件，json格式
    tableAuditQueryBo   string 

}

func NewYunosTvpubadminContentTableauditQuerymetaitemRequest() *YunosTvpubadminContentTableauditQuerymetaitemRequest{
    return &YunosTvpubadminContentTableauditQuerymetaitemRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentTableauditQuerymetaitemRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.tableaudit.querymetaitem"
}

func (r YunosTvpubadminContentTableauditQuerymetaitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentTableauditQuerymetaitemRequest) SetTableAuditQueryBo(tableAuditQueryBo string) error {
    r.tableAuditQueryBo = tableAuditQueryBo
    r.Set("table_audit_query_bo", tableAuditQueryBo)
    return nil
}

func (r YunosTvpubadminContentTableauditQuerymetaitemRequest) GetTableAuditQueryBo() string {
    return r.tableAuditQueryBo
}

