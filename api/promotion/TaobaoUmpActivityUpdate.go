package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
修改活动信息 
taobao.ump.activity.update

修改营销活动
*/
func TaobaoUmpActivityUpdate(clt *core.SDKClient, req *promotion.TaobaoUmpActivityUpdateRequest, session string) (*promotion.TaobaoUmpActivityUpdateAPIResponse, error) {
    var resp promotion.TaobaoUmpActivityUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
