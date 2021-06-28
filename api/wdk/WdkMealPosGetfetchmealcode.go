package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
五道口餐饮取餐号获取接口 
wdk.meal.pos.getfetchmealcode

pos机创建订单前获取餐饮取餐号
*/
func WdkMealPosGetfetchmealcode(clt *core.SDKClient, req *wdk.WdkMealPosGetfetchmealcodeRequest, session string) (*wdk.WdkMealPosGetfetchmealcodeAPIResponse, error) {
    var resp wdk.WdkMealPosGetfetchmealcodeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
