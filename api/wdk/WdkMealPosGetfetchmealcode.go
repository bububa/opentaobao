package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Wdkmealposgetfetchmealcode 五道口餐饮取餐号获取接口
// wdk.meal.pos.getfetchmealcode
//
// pos机创建订单前获取餐饮取餐号
func Wdkmealposgetfetchmealcode(clt *core.SDKClient, req *wdk.WdkmealposgetfetchmealcodeAPIRequest, session string) (*wdk.WdkmealposgetfetchmealcodeAPIResponse, error) {
	var resp wdk.WdkmealposgetfetchmealcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
