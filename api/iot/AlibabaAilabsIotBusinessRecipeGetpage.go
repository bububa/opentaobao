package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
分页查询食谱 
alibaba.ailabs.iot.business.recipe.getpage

分页查询食谱数据
*/
func AlibabaAilabsIotBusinessRecipeGetpage(clt *core.SDKClient, req *iot.AlibabaAilabsIotBusinessRecipeGetpageRequest, session string) (*iot.AlibabaAilabsIotBusinessRecipeGetpageResponse, error) {
    var resp iot.AlibabaAilabsIotBusinessRecipeGetpageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
