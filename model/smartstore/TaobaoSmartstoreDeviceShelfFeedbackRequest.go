package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件云货架数据回流 APIRequest
taobao.smartstore.device.shelf.feedback

智慧门店云货架设备回流
规则：
1. 回流的设备属于当前授权的用户
2. 回流的设备属于当前应用添加
*/
type TaobaoSmartstoreDeviceShelfFeedbackRequest struct {
    model.Params

    // 硬件CODE
    deviceCode   string 

    // 本次操作结束时间
    endTime   string 

    // 本次操作开始时间
    startTime   string 

    // 操作，枚举值：ACTION枚举值： ITEM_CLICK（商品点击时必须设置ITEM_ID） RECEIVE_COUPONS（领取优惠券时必须设置COUPON_ID） BUY_CLICK（点击购买） SHARE_CLICK（点击分享）
    action   string 

    // 商品ID，action为item_click必填
    itemId   string 

    // "例如官方领取优惠券链接里的activityId： https://taoquan.taobao.com/coupon/unify_apply.htm?sellerId=2649119619&activityId=9d390579777e41a981b54aa4d6154f5e"
    couponId   string 

    // 硬件识别的用户标识
    outerUser   string 

}

func NewTaobaoSmartstoreDeviceShelfFeedbackRequest() *TaobaoSmartstoreDeviceShelfFeedbackRequest{
    return &TaobaoSmartstoreDeviceShelfFeedbackRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSmartstoreDeviceShelfFeedbackRequest) GetApiMethodName() string {
    return "taobao.smartstore.device.shelf.feedback"
}

func (r TaobaoSmartstoreDeviceShelfFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSmartstoreDeviceShelfFeedbackRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

func (r TaobaoSmartstoreDeviceShelfFeedbackRequest) GetDeviceCode() string {
    return r.deviceCode
}

func (r *TaobaoSmartstoreDeviceShelfFeedbackRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoSmartstoreDeviceShelfFeedbackRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoSmartstoreDeviceShelfFeedbackRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSmartstoreDeviceShelfFeedbackRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSmartstoreDeviceShelfFeedbackRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

func (r TaobaoSmartstoreDeviceShelfFeedbackRequest) GetAction() string {
    return r.action
}

func (r *TaobaoSmartstoreDeviceShelfFeedbackRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoSmartstoreDeviceShelfFeedbackRequest) GetItemId() string {
    return r.itemId
}

func (r *TaobaoSmartstoreDeviceShelfFeedbackRequest) SetCouponId(couponId string) error {
    r.couponId = couponId
    r.Set("coupon_id", couponId)
    return nil
}

func (r TaobaoSmartstoreDeviceShelfFeedbackRequest) GetCouponId() string {
    return r.couponId
}

func (r *TaobaoSmartstoreDeviceShelfFeedbackRequest) SetOuterUser(outerUser string) error {
    r.outerUser = outerUser
    r.Set("outer_user", outerUser)
    return nil
}

func (r TaobaoSmartstoreDeviceShelfFeedbackRequest) GetOuterUser() string {
    return r.outerUser
}

