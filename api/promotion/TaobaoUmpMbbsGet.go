package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
获取营销积木块列表 
taobao.ump.mbbs.get

获取营销积木块列表，可以根据类型获取，也可以将该字段设为空，获取所有的
*/
func TaobaoUmpMbbsGet(clt *core.SDKClient, req *promotion.TaobaoUmpMbbsGetRequest, session string) (*promotion.TaobaoUmpMbbsGetAPIResponse, error) {
    var resp promotion.TaobaoUmpMbbsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
