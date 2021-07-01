package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
删除全局弹窗 
yunos.tvpubadmin.manage.dialog.delete

删除全局弹窗
*/
func YunosTvpubadminManageDialogDelete(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageDialogDeleteAPIRequest, session string) (*tvupadmin.YunosTvpubadminManageDialogDeleteAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminManageDialogDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
