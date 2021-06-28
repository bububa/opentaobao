package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
聚石塔订购公众号 
taobao.jst.sms.officialaccount.order

聚石塔订购公众号接口
*/
func TaobaoJstSmsOfficialaccountOrder(clt *core.SDKClient, req *jst.TaobaoJstSmsOfficialaccountOrderRequest, session string) (*jst.TaobaoJstSmsOfficialaccountOrderAPIResponse, error) {
    var resp jst.TaobaoJstSmsOfficialaccountOrderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
