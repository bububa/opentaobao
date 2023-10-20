package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkMealPosGetfetchmealcode 五道口餐饮取餐号获取接口
// wdk.meal.pos.getfetchmealcode
//
// pos机创建订单前获取餐饮取餐号
func WdkMealPosGetfetchmealcode(clt *core.SDKClient, req *wdk.WdkMealPosGetfetchmealcodeAPIRequest, resp *wdk.WdkMealPosGetfetchmealcodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
