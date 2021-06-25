package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
国际酒店集团价库变更通知 
taobao.xhotel.intl.ari.notify

国际酒店集团价库变更时通知变更内容，平台及时更新价库信息，保证价库新鲜度
*/
func TaobaoXhotelIntlAriNotify(clt *core.SDKClient, req *product.TaobaoXhotelIntlAriNotifyRequest, session string) (*product.TaobaoXhotelIntlAriNotifyResponse, error) {
    var resp product.TaobaoXhotelIntlAriNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
