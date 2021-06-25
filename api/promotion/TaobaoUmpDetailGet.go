package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
查询活动详情 
taobao.ump.detail.get

查询活动详情
*/
func TaobaoUmpDetailGet(clt *core.SDKClient, req *promotion.TaobaoUmpDetailGetRequest, session string) (*promotion.TaobaoUmpDetailGetResponse, error) {
    var resp promotion.TaobaoUmpDetailGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
