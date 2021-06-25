package iot

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iot"
)

/* 
获取食谱详情 
alibaba.ailabs.iot.business.recipe.getdetail

获取食谱详情接口，获取ISV自己的食谱详情数据
*/
func AlibabaAilabsIotBusinessRecipeGetdetail(clt *core.SDKClient, req *iot.AlibabaAilabsIotBusinessRecipeGetdetailRequest, session string) (*iot.AlibabaAilabsIotBusinessRecipeGetdetailResponse, error) {
    var resp iot.AlibabaAilabsIotBusinessRecipeGetdetailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
