package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
根据id查询全局弹窗 
yunos.tvpubadmin.manage.dialog.findbyid

根据id查询全局弹窗
*/
func YunosTvpubadminManageDialogFindbyid(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageDialogFindbyidRequest, session string) (*tvupadmin.YunosTvpubadminManageDialogFindbyidAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminManageDialogFindbyidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
