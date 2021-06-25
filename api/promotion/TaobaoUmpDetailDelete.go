package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
删除活动详情 
taobao.ump.detail.delete

删除活动详情
*/
func TaobaoUmpDetailDelete(clt *core.SDKClient, req *promotion.TaobaoUmpDetailDeleteRequest, session string) (*promotion.TaobaoUmpDetailDeleteResponse, error) {
    var resp promotion.TaobaoUmpDetailDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
