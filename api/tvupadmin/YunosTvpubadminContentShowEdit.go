package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
媒资节目信息修改 
yunos.tvpubadmin.content.show.edit

供迎客松修改媒资节目信息
*/
func YunosTvpubadminContentShowEdit(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentShowEditRequest, session string) (*tvupadmin.YunosTvpubadminContentShowEditAPIResponse, error) {
    var resp tvupadmin.YunosTvpubadminContentShowEditAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
