package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
系统升级分页查询 
yunos.osupdate.osfota.query

分页查询osoupdate系统升级列表
*/
func YunosOsupdateOsfotaQuery(clt *core.SDKClient, req *tvupadmin.YunosOsupdateOsfotaQueryRequest, session string) (*tvupadmin.YunosOsupdateOsfotaQueryAPIResponse, error) {
    var resp tvupadmin.YunosOsupdateOsfotaQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
