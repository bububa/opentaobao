package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSmartstoreDeviceShelfFeedbackAPIRequest
智能硬件云货架数据回流 API请求
taobao.smartstore.device.shelf.feedback

智慧门店云货架设备回流
规则：
1. 回流的设备属于当前授权的用户
2. 回流的设备属于当前应用添加 */
type TaobaoSmartstoreDeviceShelfFeedbackAPIRequest struct {
	model.Params
	// 硬件CODE
	_deviceCode string
	// 本次操作结束时间
	_endTime string
	// 本次操作开始时间
	_startTime string
	// 操作，枚举值：ACTION枚举值： ITEM_CLICK（商品点击时必须设置ITEM_ID） RECEIVE_COUPONS（领取优惠券时必须设置COUPON_ID） BUY_CLICK（点击购买） SHARE_CLICK（点击分享）
	_action string
	// 商品ID，action为item_click必填
	_itemId string
	// "例如官方领取优惠券链接里的activityId： https://taoquan.taobao.com/coupon/unify_apply.htm?sellerId=2649119619&activityId=9d390579777e41a981b54aa4d6154f5e"
	_couponId string
	// 硬件识别的用户标识
	_outerUser string
}

// New
