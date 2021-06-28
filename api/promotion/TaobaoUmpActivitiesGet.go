package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
查询活动列表 
taobao.ump.activities.get

查询活动列表
*/
func TaobaoUmpActivitiesGet(clt *core.SDKClient, req *promotion.TaobaoUmpActivitiesGetRequest, session string) (*promotion.TaobaoUmpActivitiesGetAPIResponse, error) {
    var resp promotion.TaobaoUmpActivitiesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
