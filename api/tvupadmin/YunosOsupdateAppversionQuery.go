package tvupadmin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvupadmin"
)

/* 
分页获取桌面升级任务 
yunos.osupdate.appversion.query

分页获取桌面升级任务
*/
func YunosOsupdateAppversionQuery(clt *core.SDKClient, req *tvupadmin.YunosOsupdateAppversionQueryAPIRequest, session string) (*tvupadmin.YunosOsupdateAppversionQueryAPIResponse, error) {
    var resp tvupadmin.YunosOsupdateAppversionQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
