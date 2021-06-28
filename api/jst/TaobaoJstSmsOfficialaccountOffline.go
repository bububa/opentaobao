package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
聚石塔公众号下线 
taobao.jst.sms.officialaccount.offline

聚石塔公众号下线
*/
func TaobaoJstSmsOfficialaccountOffline(clt *core.SDKClient, req *jst.TaobaoJstSmsOfficialaccountOfflineRequest, session string) (*jst.TaobaoJstSmsOfficialaccountOfflineAPIResponse, error) {
    var resp jst.TaobaoJstSmsOfficialaccountOfflineAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
