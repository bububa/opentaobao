package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
添加系统升级任务 
yunos.osupdate.osfota.add

添加osupdate系统升级任务
*/
func YunosOsupdateOsfotaAdd(clt *core.SDKClient, req *tvupadmin.YunosOsupdateOsfotaAddRequest, session string) (*tvupadmin.YunosOsupdateOsfotaAddAPIResponse, error) {
    var resp tvupadmin.YunosOsupdateOsfotaAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
