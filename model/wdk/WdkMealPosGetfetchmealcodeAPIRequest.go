package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* WdkMealPosGetfetchmealcodeAPIRequest
五道口餐饮取餐号获取接口 API请求
wdk.meal.pos.getfetchmealcode

pos机创建订单前获取餐饮取餐号 */
type WdkMealPosGetfetchmealcodeAPIRequest struct {
	model.Params
	// 渠道店id
	_channelShopId string
}

// New
