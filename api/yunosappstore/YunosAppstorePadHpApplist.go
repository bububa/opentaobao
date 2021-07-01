package yunosappstore

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/yunosappstore"
)

/* 
查询HpPad appList 
yunos.appstore.pad.hp.applist

提供hp pad应用群数据
*/
func YunosAppstorePadHpApplist(clt *core.SDKClient, req *yunosappstore.YunosAppstorePadHpApplistAPIRequest, session string) (*yunosappstore.YunosAppstorePadHpApplistAPIResponse, error) {
    var resp yunosappstore.YunosAppstorePadHpApplistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
