package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
运营位管控-查询联盟一体机运营位元数据列表 
yunos.tvpubadmin.content.tableaudit.querylauncher

运营位管控-查询联盟一体机运营位元数据列表
*/
func YunosTvpubadminContentTableauditQuerylauncher(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentTableauditQuerylauncherRequest, session string) (*tvupadmin.YunosTvpubadminContentTableauditQuerylauncherAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminContentTableauditQuerylauncherAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
