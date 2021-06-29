package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
应用升级任务更新 
yunos.osupdate.appversion.update

应用升级任务更新
*/
func YunosOsupdateAppversionUpdate(clt *core.SDKClient, req *tvupadmin.YunosOsupdateAppversionUpdateRequest, session string) (*tvupadmin.YunosOsupdateAppversionUpdateAPIResponse, error) {
    var resp tvupadmin.YunosOsupdateAppversionUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
