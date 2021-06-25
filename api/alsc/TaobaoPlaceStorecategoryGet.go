package alsc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alsc"
)

/* 
获取门店类目信息 
taobao.place.storecategory.get

获取门店类目信息
*/
func TaobaoPlaceStorecategoryGet(clt *core.SDKClient, req *alsc.TaobaoPlaceStorecategoryGetRequest, session string) (*alsc.TaobaoPlaceStorecategoryGetResponse, error) {
    var resp alsc.TaobaoPlaceStorecategoryGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
