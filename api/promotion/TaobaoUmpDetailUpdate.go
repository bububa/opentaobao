package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
修改活动详情 
taobao.ump.detail.update

更新活动详情
*/
func TaobaoUmpDetailUpdate(clt *core.SDKClient, req *promotion.TaobaoUmpDetailUpdateRequest, session string) (*promotion.TaobaoUmpDetailUpdateResponse, error) {
    var resp promotion.TaobaoUmpDetailUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
