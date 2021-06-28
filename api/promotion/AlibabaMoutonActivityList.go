package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
商家自运营活动列表 
alibaba.mouton.activity.list

商家查询自己配置的活动列表
*/
func AlibabaMoutonActivityList(clt *core.SDKClient, req *promotion.AlibabaMoutonActivityListRequest, session string) (*promotion.AlibabaMoutonActivityListAPIResponse, error) {
    var resp promotion.AlibabaMoutonActivityListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
