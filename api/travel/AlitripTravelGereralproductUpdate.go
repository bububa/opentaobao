package travel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/travel"
)

/* 
通用类目产品发布编辑 
alitrip.travel.gereralproduct.update

提供给飞猪供销平台供应商发布编辑通用类目产品的API
*/
func AlitripTravelGereralproductUpdate(clt *core.SDKClient, req *travel.AlitripTravelGereralproductUpdateAPIRequest, session string) (*travel.AlitripTravelGereralproductUpdateAPIResponse, error) {
    var resp travel.AlitripTravelGereralproductUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
