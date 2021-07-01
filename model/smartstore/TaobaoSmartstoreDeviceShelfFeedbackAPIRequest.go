package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件云货架数据回流 API请求
taobao.smartstore.device.shelf.feedback

智慧门店云货架设备回流
规则：
1. 回流的设备属于当前授权的用户
2. 回流的设备属于当前应用添加
*/
type TaobaoSmartstoreDeviceShelfFeedbackAPIRequest struct {
    model.Params
    // 硬件CODE
    _deviceCode   string
    // 本次操作结束时间
    _endTime   string
    // 本次操作开始时间
    _startTime   string
    // 操作，枚举值：ACTION枚举值： ITEM_CLICK（商品点击时必须设置ITEM_ID） RECEIVE_COUPONS（领取优惠券时必须设置COUPON_ID） BUY_CLICK（点击购买） SHARE_CLICK（点击分享）
    _action   string
    // 商品ID，action为item_click必填
    _itemId   string
    // "例如官方领取优惠券链接里的activityId： https://taoquan.taobao.com/coupon/unify_apply.htm?sellerId=2649119619&activityId=9d390579777e41a981b54aa4d6154f5e"
    _couponId   string
    // 硬件识别的用户标识
    _outerUser   string
}

// 初始化TaobaoSmartstoreDeviceShelfFeedbackAPIRequest对象
func NewTaobaoSmartstoreDeviceShelfFeedbackRequest() *TaobaoSmartstoreDeviceShelfFeedbackAPIRequest{
    return &TaobaoSmartstoreDeviceShelfFeedbackAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) GetApiMethodName() string {
    return "taobao.smartstore.device.shelf.feedback"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceCode Setter
// 硬件CODE
func (r *TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) SetDeviceCode(_deviceCode string) error {
    r._deviceCode = _deviceCode
    r.Set("device_code", _deviceCode)
    return nil
}

// DeviceCode Getter
func (r TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) GetDeviceCode() string {
    return r._deviceCode
}
// EndTime Setter
// 本次操作结束时间
func (r *TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) GetEndTime() string {
    return r._endTime
}
// StartTime Setter
// 本次操作开始时间
func (r *TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) GetStartTime() string {
    return r._startTime
}
// Action Setter
// 操作，枚举值：ACTION枚举值： ITEM_CLICK（商品点击时必须设置ITEM_ID） RECEIVE_COUPONS（领取优惠券时必须设置COUPON_ID） BUY_CLICK（点击购买） SHARE_CLICK（点击分享）
func (r *TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) GetAction() string {
    return r._action
}
// ItemId Setter
// 商品ID，action为item_click必填
func (r *TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) SetItemId(_itemId string) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) GetItemId() string {
    return r._itemId
}
// CouponId Setter
// "例如官方领取优惠券链接里的activityId： https://taoquan.taobao.com/coupon/unify_apply.htm?sellerId=2649119619&activityId=9d390579777e41a981b54aa4d6154f5e"
func (r *TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) SetCouponId(_couponId string) error {
    r._couponId = _couponId
    r.Set("coupon_id", _couponId)
    return nil
}

// CouponId Getter
func (r TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) GetCouponId() string {
    return r._couponId
}
// OuterUser Setter
// 硬件识别的用户标识
func (r *TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) SetOuterUser(_outerUser string) error {
    r._outerUser = _outerUser
    r.Set("outer_user", _outerUser)
    return nil
}

// OuterUser Getter
func (r TaobaoSmartstoreDeviceShelfFeedbackAPIRequest) GetOuterUser() string {
    return r._outerUser
}
