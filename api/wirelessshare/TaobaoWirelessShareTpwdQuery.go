package wirelessshare

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wirelessshare"
)

/* 
查询解析淘口令 
taobao.wireless.share.tpwd.query

查询解析淘口令
*/
func TaobaoWirelessShareTpwdQuery(clt *core.SDKClient, req *wirelessshare.TaobaoWirelessShareTpwdQueryRequest, session string) (*wirelessshare.TaobaoWirelessShareTpwdQueryAPIResponse, error) {
    var resp wirelessshare.TaobaoWirelessShareTpwdQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
