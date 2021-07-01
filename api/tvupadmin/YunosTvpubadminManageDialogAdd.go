package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
新增全局弹窗 
yunos.tvpubadmin.manage.dialog.add

新增全局弹窗
*/
func YunosTvpubadminManageDialogAdd(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageDialogAddAPIRequest, session string) (*tvupadmin.YunosTvpubadminManageDialogAddAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminManageDialogAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
