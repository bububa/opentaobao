package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
编辑全局弹窗 
yunos.tvpubadmin.manage.dialog.edit

编辑全局弹窗
*/
func YunosTvpubadminManageDialogEdit(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageDialogEditAPIRequest, session string) (*tvupadmin.YunosTvpubadminManageDialogEditAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminManageDialogEditAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
