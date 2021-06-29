package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
更新应用升级状态 
yunos.osupdate.versionstatus.update

更新应用升级状态
*/
func YunosOsupdateVersionstatusUpdate(clt *core.SDKClient, req *tvupadmin.YunosOsupdateVersionstatusUpdateRequest, session string) (*tvupadmin.YunosOsupdateVersionstatusUpdateAPIResponse, error) {
    var resp tvupadmin.YunosOsupdateVersionstatusUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
